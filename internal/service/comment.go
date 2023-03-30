package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/object"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/static"
	"net/http"
	"net/url"
	"strings"
)

type CommentService struct {
	httpService *HttpService
}

func NewCommentService() *CommentService {
	return &CommentService{
		httpService: xmodule.MustGetByName(sn.SHttpService).(*HttpService),
	}
}

func (c *CommentService) GetComments(mid uint64, page int32) ([]*object.Comment, int32, error) {
	u := fmt.Sprintf(static.MANGA_COMMENT_URL, mid, page)
	bs, _, err := c.httpService.HttpGet(u, nil)
	if err != nil {
		return nil, 0, err
	}
	commentsObj := &object.Comments{}
	err = json.Unmarshal(bs, commentsObj)
	if err != nil {
		return nil, 0, err
	}

	objArr := commentsObj.CommentIds
	chains := make([][]string, len(objArr))
	for idx, idsStr := range objArr {
		chain := strings.Split(idsStr, ",")
		if len(chain) > 0 {
			chains[idx] = chain
		}
	}

	out := make([]*object.Comment, 0, len(chains))
	for _, chain := range chains {
		cmt, ok := commentsObj.Comments[chain[0]]
		if !ok {
			continue
		}
		if cmt.Username == "" {
			cmt.Username = "-"
		}
		if cmt.Avatar == "" {
			cmt.Avatar = static.DEFAULT_USER_AVATAR_URL
		}

		timeline := make([]*object.RepliedComment, 0, len(chain)-1)
		if len(chain) > 1 {
			for idx := len(chain) - 1; idx >= 1; idx-- {
				repliedId := chain[idx]
				if reply, ok := commentsObj.Comments[repliedId]; ok {
					cmt := object.NewRepliedComment(reply)
					if cmt.Username == "" {
						cmt.Username = "-"
					}
					if cmt.Avatar == "" {
						cmt.Avatar = static.DEFAULT_USER_AVATAR_URL
					}
					timeline = append(timeline, cmt)
				}
			}
		}
		cmt.ReplyTimeline = timeline
		out = append(out, cmt)
	}

	return out, commentsObj.Total, nil
}

func (c *CommentService) LikeComment(cid uint64) error {
	u := fmt.Sprintf(static.MANGA_LIKE_COMMENT_URL, cid)
	bs, _, err := c.httpService.HttpGet(u, nil)
	if err != nil {
		return err
	}

	type model struct {
		Status int32 `json:"status"`
	}
	m := &model{}
	err = json.Unmarshal(bs, m)
	if err != nil {
		return err
	}

	// {"status":1}
	if m.Status == 0 {
		return fmt.Errorf("can not like comment %d", cid)
	}
	return nil
}

func (c *CommentService) _httpPostWithToken(url, token string, form *url.Values) ([]byte, error) {
	bs, _, err := c.httpService.HttpPost(url, strings.NewReader(form.Encode()), func(req *http.Request) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", "my="+token)
	})
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (c *CommentService) AddComment(token string, mid uint64, text string) (comment *object.AddedComment, auth bool, err error) {
	return c.addComment(token, mid, 0, text)
}

func (c *CommentService) ReplyComment(token string, mid, cid uint64, text string) (comment *object.AddedComment, auth bool, err error) {
	return c.addComment(token, mid, cid, text)
}

func (c *CommentService) addComment(token string, mid, repliedCid uint64, text string) (comment *object.AddedComment, auth bool, err error) {
	escapedText := url.PathEscape(text) // escape first
	form := &url.Values{"book_id": {xnumber.U64toa(mid)}, "to_comment_id": {xnumber.U64toa(repliedCid)}, "txtContent": {escapedText}}
	bs, err := c._httpPostWithToken(static.MANGA_ADD_COMMENT_URL, token, form)
	if err != nil {
		return nil, false, err
	}

	type model struct {
		Status    int32  `json:"status"`
		Msg       string `json:"msg"`
		CommentId uint64 `json:"comment_id"`
	}
	m := &model{}
	err = json.Unmarshal(bs, m)
	if err != nil {
		return nil, false, err
	}

	// {"status": 1, "msg": "恭喜您，评论提交成功！", "comment_id":1247467}
	// {"status": 0, "msg": "对不起，登录后才能评论！"}
	if m.Status == 0 {
		auth = m.Msg != "对不起，登录后才能评论！"
		return nil, auth, nil
	}
	comment = object.NewAddedComment(m.CommentId, mid, repliedCid, text)
	return comment, true, nil
}
