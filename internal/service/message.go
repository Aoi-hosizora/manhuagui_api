package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/ahlib/xpointer"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/object"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/static"
	"gopkg.in/yaml.v2"
	"math"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

type MessageService struct {
	httpService *HttpService
}

func NewMessageService() *MessageService {
	return &MessageService{
		httpService: xmodule.MustGetByName(sn.SHttpService).(*HttpService),
	}
}

func (m *MessageService) GetAllMessages(token string) ([]*object.Message, error) {
	pageCount, err := m.getCommentPageCount(token)
	if err != nil {
		return nil, err
	}

	// get messages
	messageLists := make([][]*object.Message, pageCount)
	err = error(nil)
	once := sync.Once{}
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for page := range messageLists {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			messages, thisErr := m.getMessages(int32(page)+1, token)
			if thisErr != nil {
				once.Do(func() { err = thisErr })
			} else {
				mu.Lock()
				messageLists[page] = messages
				mu.Unlock()
			}
		}(page)
	}
	wg.Wait()
	if err != nil {
		return nil, err
	}

	// merge messages
	vis := make(map[uint64]bool)
	messages := make([]*object.Message, 0)
	for _, messageList := range messageLists {
		for _, msg := range messageList {
			if _, ok := vis[msg.Mid]; !ok {
				messages = append(messages, msg)
				vis[msg.Mid] = true
			}
		}
	}
	sort.Slice(messages, func(i, j int) bool {
		return messages[j].Mid < messages[i].Mid // desc
	})
	return messages, nil
}

func (m *MessageService) GetLatestMessage(token string) (*object.LatestMessage, error) {
	messages, err := m.GetAllMessages(token) // almost one page
	if err != nil {
		return nil, err
	}

	latestMessage := &object.LatestMessage{}
	for _, msg := range messages {
		if msg.Notification != nil && *msg.Notification.Dismissible {
			latestMessage.Notification = msg
			break
		}
	}
	for _, msg := range messages {
		if msg.Notification != nil && !*msg.Notification.Dismissible {
			latestMessage.NotDismissibleNotification = msg
			break
		}
	}
	for _, msg := range messages {
		if msg.NewVersion != nil && !*msg.NewVersion.MustUpgrade {
			latestMessage.NewVersion = msg
			break
		}
	}
	for _, msg := range messages {
		if msg.NewVersion != nil && *msg.NewVersion.MustUpgrade {
			latestMessage.MustUpgradeNewVersion = msg
			break
		}
	}
	return latestMessage, nil
}

// ===

func (m *MessageService) getCommentPageCount(token string) (int32, error) {
	bs, resp, err := m.httpService.HttpGet(static.MESSAGE_ISSUE_API, func(r *http.Request) {
		r.Header.Set("Accept", static.GITHUB_ACCEPT)
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	})
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("got %d response code", resp.StatusCode)
	}

	issue := &object.Issue{}
	err = json.Unmarshal(bs, issue)
	if err != nil {
		return 0, err
	}

	commentCount := float64(issue.Comments)
	pageCount := math.Ceil(commentCount / static.MESSAGE_COMMENTS_PERPAGE /* 100 */) // almost one page
	return int32(pageCount), nil
}

func (m *MessageService) getMessages(page int32, token string) ([]*object.Message, error) {
	apiUrl := fmt.Sprintf(static.MESSAGE_COMMENTS_API, page, static.MESSAGE_COMMENTS_PERPAGE /* 100 */)
	bs, resp, err := m.httpService.HttpGet(apiUrl, func(r *http.Request) {
		r.Header.Set("Accept", static.GITHUB_ACCEPT)
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("got non-200 response")
	}

	// unmarshal
	comments := make([]*object.IssueComment, 0)
	err = json.Unmarshal(bs, &comments)
	if err != nil {
		return nil, err
	}

	// parse body
	messages := make([]*object.Message, 0, len(comments))
	for _, comment := range comments {
		bodyString := strings.TrimSpace(strings.Trim(strings.Trim(strings.TrimSpace(comment.Body), "```yaml"), "```"))
		bodyObj := &object.IssueCommentBody{}
		err = yaml.Unmarshal([]byte(bodyString), bodyObj)
		if err != nil {
			continue
		}

		message := m.parseCommentBody(comment, bodyObj)
		if message != nil {
			messages = append(messages, message)
		}
	}
	return messages, nil
}

func (m *MessageService) parseCommentBody(comment *object.IssueComment, body *object.IssueCommentBody) *object.Message {
	out := &object.Message{
		Mid:          comment.Id,
		Title:        body.Title,
		Notification: body.Notification,
		NewVersion:   body.NewVersion,
		CreatedAt:    comment.CreatedAt.Local(),
		UpdatedAt:    comment.UpdatedAt.Local(),
	}

	if out.Notification == nil && out.NewVersion == nil {
		return nil // error
	}
	if out.Notification != nil && out.NewVersion != nil {
		return nil // error
	}

	if id, err := xnumber.ParseUint64(body.MidStr, 10); err == nil {
		out.Mid = id
	}
	if createdAt, err := time.Parse(time.RFC3339, body.CreatedAt); err == nil {
		out.CreatedAt = createdAt
	}
	if updatedAt, err := time.Parse(time.RFC3339, body.UpdatedAt); err == nil {
		out.UpdatedAt = updatedAt
	}

	out.Title = strings.TrimSpace(out.Title)
	if out.Notification != nil {
		out.Notification.Content = strings.TrimSpace(out.Notification.Content)
		if out.Notification.Dismissible == nil {
			out.Notification.Dismissible = xpointer.BoolPtr(true) // <<< defaults to true
		}
		out.Notification.Link = strings.TrimSpace(out.Notification.Link)
	}
	if out.NewVersion != nil {
		out.NewVersion.Version = strings.TrimSpace(out.NewVersion.Version)
		if out.NewVersion.MustUpgrade == nil {
			out.NewVersion.MustUpgrade = xpointer.BoolPtr(false) // <<< defaults to false
		}
		out.NewVersion.ChangeLogs = strings.TrimSpace(out.NewVersion.ChangeLogs)
		out.NewVersion.ReleasePage = strings.TrimSpace(out.NewVersion.ReleasePage)
	}
	return out
}
