package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
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
		httpService: xdi.GetByNameForce(sn.SHttpService).(*HttpService),
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

	username := doc.Find("div.head-box div.inner h3").Text()
	username = strings.TrimSuffix(strings.TrimPrefix(username, "尊敬的会员 "), "，欢迎您！")
	avatar := doc.Find("div.head-box div.img-box img").AttrOr("src", "")
	class := doc.Find("div.head-box div.inner p").First().Text()
	class = strings.TrimPrefix(class, "您的会员等级：")
	scoreStr := doc.Find("div.head-box div.inner p").First().Next().Text()
	score, _ := xnumber.Atoi32(strings.TrimSuffix(strings.TrimPrefix(scoreStr, "个人成长值："), "点"))

	recordDiv := doc.Find("div.head-inner:last-of-type")
	loginIP := recordDiv.Find("dl:nth-of-type(1) dd").Text()
	lastLoginIP := recordDiv.Find("dl:nth-of-type(2) dd").Text()
	registerTime := recordDiv.Find("dl:nth-of-type(3) dd").Text()
	lastLoginTime := recordDiv.Find("dl:nth-of-type(4) dd").Text()

	user := &vo.User{
		Username:      username,
		Avatar:        avatar,
		Class:         class,
		Score:         score,
		LoginIP:       loginIP,
		LastLoginIP:   lastLoginIP,
		RegisterTime:  registerTime,
		LastLoginTime: lastLoginTime,
	}
	return user, nil
}

func (u *UserService) RecordManga(token string, mid, cid uint64) error {
	ur := fmt.Sprintf(static.MANGA_COUNT_URL, mid, cid)
	_, _, err := u._httpGetWithToken(ur, token)
	return err
}
