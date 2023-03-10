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
	"math"
	"net/http"
	"net/url"
	"strings"
)

type ShelfService struct {
	httpService *HttpService
}

func NewShelfService() *ShelfService {
	return &ShelfService{
		httpService: xmodule.MustGetByName(sn.SHttpService).(*HttpService),
	}
}

func (s *ShelfService) _httpGetWithToken(url, token string) ([]byte, *goquery.Document, error) {
	bs, doc, err := s.httpService.HttpGetDocument(url, func(req *http.Request) {
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

func (s *ShelfService) _httpPostWithToken(url, token string, form *url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "my="+token)
	bs, _, err := s.httpService.DoRequest(req)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (s *ShelfService) GetShelfMangas(token string, page int32) ([]*vo.ShelfManga, int32, int32, error) {
	u := fmt.Sprintf(static.MANGA_SHELF_URL, page)
	_, doc, err := s._httpGetWithToken(u, token)
	if err != nil {
		return nil, 0, 0, err
	} else if doc == nil {
		return nil, 0, 0, nil
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

	totalSpan := doc.Find("div.flickr span:first-child")
	total, _ := xnumber.Atoi32(strings.TrimSuffix(strings.TrimPrefix(totalSpan.Text(), "共"), "记录"))
	limit := int32(20)
	if totalSpan.Length() > 0 {
		pages := int32(math.Ceil(float64(total) / float64(limit)))
		if page > pages {
			return []*vo.ShelfManga{}, limit, total, nil
		}
	} else {
		total = int32(len(mangas))
		if page > 1 {
			return []*vo.ShelfManga{}, limit, total, nil
		}
	}

	return mangas, limit, total, nil
}

func (s *ShelfService) CheckMangaInShelf(token string, mid uint64) (*vo.ShelfStatus, error) {
	u := fmt.Sprintf(static.MANGA_SHELF_CHECK_URL, mid)
	bs, doc, err := s._httpGetWithToken(u, token)
	if err != nil {
		return nil, err
	} else if doc == nil {
		return nil, nil
	}

	// {"status":0, "total":3274}
	status := &vo.ShelfStatus{}
	err = json.Unmarshal(bs, status)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (s *ShelfService) SaveMangaToShelf(token string, mid uint64) (auth bool, existed bool, err error) {
	bs, err := s._httpPostWithToken(static.MANGA_SHELF_ADD_URL, token, &url.Values{"book_id": {xnumber.U64toa(mid)}})
	if err != nil {
		return false, false, err
	}

	type model struct {
		Status int32  `json:"status"`
		Msg    string `json:"msg"`
	}
	m := &model{}
	err = json.Unmarshal(bs, m)
	if err != nil {
		return true, false, err
	}

	// {"status": 1, "msg": "恭喜您，收藏到书架成功！"}
	// {"status":0, "msg":"对不起，用户尚未登录或已超时！"}
	// {"status":0, "msg":"您书架里已经有这部漫画了！"}
	if m.Status == 0 {
		auth = m.Msg != "对不起，用户尚未登录或已超时！"
		existed = m.Msg == "您书架里已经有这部漫画了！"
		return auth, existed, nil
	}
	return true, false, nil
}

func (s *ShelfService) RemoveMangaFromShelf(token string, mid uint64) (auth bool, notFound bool, err error) {
	bs, err := s._httpPostWithToken(static.MANGA_SHELF_DELETE_URL, token, &url.Values{"book_id": {xnumber.U64toa(mid)}})
	if err != nil {
		return false, false, err
	}

	type model struct {
		Status int32  `json:"status"`
		Msg    string `json:"msg"`
	}
	m := &model{}
	err = json.Unmarshal(bs, m)
	if err != nil {
		return true, false, err
	}

	// {"status":0, "msg":"恭喜您，从书架移除成功！"}
	// {"status":0, "msg":"对不起，用户尚未登录或已超时！"}
	// {"status":1, "msg":"恭喜您，从书架移除成功！"}
	if m.Status == 0.0 {
		auth = m.Msg != "对不起，用户尚未登录或已超时！"
		notFound = m.Msg == "恭喜您，从书架移除成功！"
		return auth, notFound, nil
	}
	return true, false, nil
}
