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

func (u *UserService) CheckLogin(token string) (bool, error) {
	req, err := http.NewRequest("POST", static.MANGA_CHECK_LOGIN_URL, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Cookie", "my="+token)
	bs, _, err := u.httpService.DoRequest(req)
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

func (u *UserService) _login(url, token string) ([]byte, *goquery.Document, error) {
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
	_, doc, err := u._login(static.MANGA_USER_URL, token)
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
	ur := fmt.Sprintf(static.MANGA_SHELF_URL, page)
	_, doc, err := u._login(ur, token)
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
		ur := sel.Find("h3 a").AttrOr("href", "")
		newestChapter := sel.Find("p:nth-of-type(1) em:nth-child(1)").Text()
		newestDuration := sel.Find("p:nth-of-type(1) em:nth-child(2)").Text()
		lastChapter := sel.Find("p:nth-of-type(2) em:nth-child(1)").Text()
		lastDuration := sel.Find("p:nth-of-type(2) em:nth-child(2)").Text()

		manga := &vo.ShelfManga{
			Mid:            static.ParseMid(ur),
			Title:          title,
			Cover:          static.ParseCoverUrl(cover),
			Url:            static.HOMEPAGE_URL + ur,
			NewestChapter:  newestChapter,
			NewestDuration: newestDuration,
			LastChapter:    lastChapter,
			LastDuration:   lastDuration,
		}
		mangas = append(mangas, manga)
	})

	return mangas, limit, total, nil
}

func (u *UserService) CheckMangaInShelf(token string, mid uint64) (auth bool, in bool, err error) {
	ur := fmt.Sprintf(static.MANGA_SHELF_CHECK_URL, mid)
	bs, doc, err := u._login(ur, token)
	if err != nil {
		return false, false, err
	} else if doc == nil {
		return false, false, nil
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(bs, &m)
	if err != nil {
		return true, false, err
	}
	status, ok := m["status"]
	if !ok {
		return true, false, fmt.Errorf("failed to check shelf")
	}

	// {"status":0, "total":3274}
	in = status == 1.0
	return true, in, nil
}

func (u *UserService) SaveMangaToShelf(token string, mid uint64) (auth bool, existed bool, err error) {
	form := url.Values{}
	form.Set("book_id", xnumber.U64toa(mid))
	req, err := http.NewRequest("POST", static.MANGA_SHELF_ADD_URL, strings.NewReader(form.Encode()))
	if err != nil {
		return false, false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "my="+token)
	bs, _, err := u.httpService.DoRequest(req)
	if err != nil {
		return false, false, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(bs, &m)
	if err != nil {
		return true, false, err
	}
	msg, ok1 := m["msg"]
	status, ok2 := m["status"]
	if !ok1 || !ok2 {
		return true, false, fmt.Errorf("failed to add to shelf")
	}

	// {"status": 1, "msg": "恭喜您，收藏到书架成功！"}
	// {"status":0, "msg":"对不起，用户尚未登录或已超时！"}
	// {"status":0, "msg":"您书架里已经有这部漫画了！"}
	if status == 0.0 {
		auth = msg != "对不起，用户尚未登录或已超时！"
		existed = msg == "您书架里已经有这部漫画了！"
		return auth, existed, nil
	}
	return true, false, nil
}

func (u *UserService) RemoveMangaFromShelf(token string, mid uint64) (auth bool, notFound bool, err error) {
	form := url.Values{}
	form.Set("book_id", xnumber.U64toa(mid))
	req, err := http.NewRequest("POST", static.MANGA_SHELF_DELETE_URL, strings.NewReader(form.Encode()))
	if err != nil {
		return false, false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "my="+token)
	bs, _, err := u.httpService.DoRequest(req)
	if err != nil {
		return false, false, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(bs, &m)
	if err != nil {
		return true, false, err
	}
	msg, ok1 := m["msg"]
	status, ok2 := m["status"]
	if !ok1 || !ok2 {
		return true, false, fmt.Errorf("failed to delete from shelf")
	}

	// {"status":0, "msg":"恭喜您，从书架移除成功！"}
	// {"status":0, "msg":"对不起，用户尚未登录或已超时！"}
	// {"status":1, "msg":"恭喜您，从书架移除成功！"}
	if status == 0.0 {
		auth = msg != "对不起，用户尚未登录或已超时！"
		notFound = msg == "恭喜您，从书架移除成功！"
		return auth, notFound, nil
	}
	return true, false, nil
}
