package service

import (
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
		Url:  static.HOMEPAGE_URL + url,
	}
}
