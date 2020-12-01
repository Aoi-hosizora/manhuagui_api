package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
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
	form := fmt.Sprintf("txtUserName=%s&txtPassword=%s", username, password)
	req, err := http.NewRequest("POST", static.MANGA_LOGIN_URL, strings.NewReader(form))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
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

func (u *UserService) CheckLogin(token string) (bool, error) {
	req, err := http.NewRequest("POST", static.MANGA_CHECK_LOGIN_URL, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Cookie", "my="+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return false, err
	}

	status := &vo.UserStatus{}
	err = json.Unmarshal(bs, status)
	if err != nil {
		return false, err
	}
	ok := status.Status == 1 && status.Username != "匿名用户"

	return ok, nil
}

func (u *UserService) _login(url, token string) (*goquery.Document, error) {
	bs, doc, err := u.httpService.HttpGetDocument(url, func(req *http.Request) {
		req.Header.Add("Cookie", "my="+token)
	})
	if err != nil {
		return nil, err
	}
	if bytes.Contains(bs, []byte(static.UNAUTHORIZED_TOKEN)) {
		return nil, nil
	}
	return doc, nil
}

func (u *UserService) GetUser(token string) (*vo.User, error) {
	_, err := u._login(static.MANGA_USER_URL, token)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserService) GetShelfMangas(token string) (interface{}, error) {
	_, err := u._login(static.MANGA_SHELF_URL, token)
	if err != nil {
		return nil, err
	}
	return true, nil
}
