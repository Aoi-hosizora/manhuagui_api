package service

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type AuthorService struct {
	httpService *HttpService
}

func NewAuthorService() *AuthorService {
	return &AuthorService{
		httpService: xdi.GetByNameForce(sn.SHttpService).(*HttpService),
	}
}

func (a *AuthorService) GetAuthorFromA(sel *goquery.Selection) *vo.TinyAuthor {
	name := sel.AttrOr("title", "")
	url := strings.TrimSuffix(sel.AttrOr("href", ""), "/")
	sp := strings.Split(url, "/")
	aid, _ := xnumber.Atou64(sp[len(sp)-1])
	return &vo.TinyAuthor{
		Aid:  aid,
		Name: name,
		Url:  strings.TrimSuffix(static.HOMEPAGE_URL+url, "/"),
	}
}

func (a *AuthorService) GetAllAuthors(page int32, orderByPopular bool) ([]*vo.SmallAuthor, int32, int32, error) {
	url := static.MANGA_AUTHORS_URL
	if orderByPopular {
		url += fmt.Sprintf("/%s_p%d.html", "view", page)
	} else {
		url += fmt.Sprintf("/%s_p%d.html", "index", page)
	}

	doc, err := a.httpService.HttpGetDocument(url)
	if err != nil {
		return nil, 0, 0, err
	}
	if doc == nil {
		return nil, 0, 0, nil
	}

	limit := int32(42)
	pages, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(2)").Text())
	total, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(3)").Text())

	authors := make([]*vo.SmallAuthor, 0)
	if page <= pages {
		listLis := doc.Find("ul#contList li")
		listLis.Each(func(idx int, li *goquery.Selection) {
			authors = append(authors, a.getSmallUserFromLi(li))
		})
	}

	return authors, limit, total, nil
}

func (a *AuthorService) getSmallUserFromLi(li *goquery.Selection) *vo.SmallAuthor {
	name := li.Find("p a").AttrOr("title", "")
	url := li.Find("p a").AttrOr("href", "")
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	aid, _ := xnumber.Atou64(sp[len(sp)-1])
	zone := li.Find("font:nth-child(2)").Text()
	mangaCount, _ := xnumber.Atoi32(li.Find("font:nth-child(3)").Text())
	score := li.Find("span.updateon em").Text()
	newestDate := strings.TrimPrefix(strings.TrimSuffix(li.Find("span.updateon").Text(), score), "更新于：")
	return &vo.SmallAuthor{
		Aid:        aid,
		Name:       name,
		Zone:       zone,
		Url:        strings.TrimSuffix(static.HOMEPAGE_URL+url, "/"),
		MangaCount: mangaCount,
		NewestDate: newestDate,
	}
}
