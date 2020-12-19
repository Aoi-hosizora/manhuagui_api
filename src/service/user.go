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
	"io/ioutil"
	"math"
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
	doc, err := u._login(static.MANGA_USER_URL, token)
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

func (u *UserService) GetShelfMangas(token string, page int32) ([]*vo.ShelfManga, int32, int32, error) {
	url := fmt.Sprintf(static.MANGA_SHELF_URL, page)
	doc, err := u._login(url, token)
	if err != nil {
		return nil, 0, 0, err
	} else if doc == nil {
		return nil, 0, 0, nil
	}

	total, _ := xnumber.Atoi32(strings.TrimSuffix(strings.TrimPrefix(doc.Find("div.flickr span:first-child").Text(), "共"), "记录"))
	limit := int32(20)
	pages := int32(math.Ceil(float64(total) / float64(limit)))
	if page > pages {
		return []*vo.ShelfManga{}, limit, total, nil
	}

	mangas := make([]*vo.ShelfManga, 0)
	divs := doc.Find("div.dy_content_li")
	divs.Each(func(idx int, sel *goquery.Selection) {
		cover := sel.Find("img").AttrOr("src", "")
		title := sel.Find("h3").Text()
		url := sel.Find("h3 a").AttrOr("href", "")
		newestChapter := sel.Find("p:nth-of-type(1) em:nth-child(1)").Text()
		newestDuration := sel.Find("p:nth-of-type(1) em:nth-child(2)").Text()
		lastChapter := sel.Find("p:nth-of-type(2) em:nth-child(1)").Text()
		lastDuration := sel.Find("p:nth-of-type(2) em:nth-child(2)").Text()

		manga := &vo.ShelfManga{
			Mid:            static.ParseMid(url),
			Title:          title,
			Cover:          static.ParseCoverUrl(cover),
			Url:            static.HOMEPAGE_URL + url,
			NewestChapter:  newestChapter,
			NewestDuration: newestDuration,
			LastChapter:    lastChapter,
			LastDuration:   lastDuration,
		}
		mangas = append(mangas, manga)
	})

	return mangas, limit, total, nil
}
