package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/vo"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/static"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

type UserService struct {
	httpService *HttpService
}

func NewUserService() *UserService {
	return &UserService{
		httpService: xmodule.MustGetByName(sn.SHttpService).(*HttpService),
	}
}

func (u *UserService) Login(username, password string) (string, error) {
	form := fmt.Sprintf("txtUserName=%s&txtPassword=%s", url.QueryEscape(username), url.QueryEscape(password))
	req, err := http.NewRequest("POST", static.MANGA_LOGIN_URL, strings.NewReader(form))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	bs, resp, err := u.httpService.DoRequest(req)
	if err != nil {
		return "", err
	}

	if !bytes.Contains(bs, []byte(static.LOGIN_SUCCESS_TOKEN)) {
		return "", nil
	}
	my := ""
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "my" {
			my = cookie.Value
		}
	}
	return my, nil
}

func (u *UserService) CheckLogin(token string) (bool, string, error) {
	req, err := http.NewRequest("POST", static.MANGA_CHECK_LOGIN_URL, nil)
	if err != nil {
		return false, "", err
	}
	req.Header.Set("Cookie", "my="+token)
	bs, _, err := u.httpService.DoRequest(req)
	if err != nil {
		return false, "", err
	}

	status := &vo.UserStatus{}
	err = json.Unmarshal(bs, status)
	if err != nil {
		return false, "", err
	}

	ok := status.Status == 1 && status.Username != "匿名用户"
	if !ok {
		return false, "", nil
	}
	return ok, status.Username, nil
}

func (u *UserService) _httpGetWithToken(url, token string) ([]byte, *goquery.Document, error) {
	bs, doc, err := u.httpService.HttpGetDocument(url, func(req *http.Request) {
		req.Header.Set("Cookie", "my="+token)
	})
	if err != nil {
		return nil, nil, err
	}
	if bytes.Contains(bs, []byte(static.UNAUTHORIZED_TOKEN)) {
		return nil, nil, nil
	}
	return bs, doc, nil
}

func (u *UserService) GetUser(token string) (*vo.User, error) {
	_, doc, err := u._httpGetWithToken(static.MANGA_USER_URL, token)
	if err != nil {
		return nil, err
	} else if doc == nil {
		return nil, nil
	}

	// 会员中心
	username := doc.Find("div.head-box div.inner h3").Text()
	username = strings.TrimSuffix(strings.TrimPrefix(username, "尊敬的会员 "), "，欢迎您！")
	avatar := doc.Find("div.head-box div.img-box img").AttrOr("src", "")
	class := doc.Find("div.head-box div.inner p").First().Text()
	class = strings.TrimPrefix(class, "您的会员等级：")
	scoreStr := doc.Find("div.head-box div.inner p").First().Next().Text()
	score, _ := xnumber.Atoi32(strings.TrimSuffix(strings.TrimPrefix(scoreStr, "个人成长值："), "点"))

	// 账户统计
	accountDiv := doc.Find("div.head-inner").First()
	accountPoint, _ := xnumber.Atoi32(accountDiv.Find("dl:nth-of-type(1) dd b").Text())
	unreadMessageCount, _ := xnumber.Atoi32(accountDiv.Find("dl:nth-of-type(2) dd b").Text())

	// 登录统计
	loginDiv := doc.Find("div.head-inner:last-of-type")
	loginIP := loginDiv.Find("dl:nth-of-type(1) dd").Text()
	lastLoginIP := loginDiv.Find("dl:nth-of-type(2) dd").Text()
	registerTime := loginDiv.Find("dl:nth-of-type(3) dd").Text()
	lastLoginTime := loginDiv.Find("dl:nth-of-type(4) dd").Text()
	cumulativeDayCount, _ := xnumber.Atoi32(loginDiv.Find("dl:nth-of-type(5) dd").Text())
	totalCommentCount, _ := xnumber.Atoi32(loginDiv.Find("dl:nth-of-type(6) dd").Text())

	user := &vo.User{
		Username:           username,
		Avatar:             avatar,
		Class:              class,
		Score:              score,
		AccountPoint:       accountPoint,
		UnreadMessageCount: unreadMessageCount,
		LoginIP:            loginIP,
		LastLoginIP:        lastLoginIP,
		RegisterTime:       registerTime,
		LastLoginTime:      lastLoginTime,
		CumulativeDayCount: cumulativeDayCount,
		TotalCommentCount:  totalCommentCount,
	}
	return user, nil
}

func (u *UserService) RecordManga(token string, mid, cid uint64) error {
	ur := fmt.Sprintf(static.MANGA_COUNT_URL, mid, cid)
	_, _, err := u._httpGetWithToken(ur, token)
	return err
}
