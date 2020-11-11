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

type MangaListService struct {
	httpService *HttpService
}

func NewMangaListService() *MangaListService {
	return &MangaListService{
		httpService: xdi.GetByNameForce(sn.SHttpService).(*HttpService),
	}
}

func (m *MangaListService) getMangas(doc *goquery.Document, tagIndex int, tagName string) ([]*vo.MangaPageGroup, []*vo.MangaPageGroup) {
	groups := make([]*vo.MangaPageGroup, 0)

	// get top mangas
	topMangas := make([]*vo.MangaPageLink, 0)
	topMangaUl := doc.Find("div.cmt-cont ul:nth-child(" + xnumber.Itoa(tagIndex) + ")") // <<<
	topMangaUl.Find("li").Each(func(idx int, li *goquery.Selection) {
		topMangas = append(topMangas, m.getMangaPageLinkFromLi(li))
	})
	groups = append(groups, &vo.MangaPageGroup{
		Title:  "*",
		Mangas: topMangas,
	})

	// get group mangas
	serialDiv := doc.Find("div#" + tagName + "Cont") // <<<
	serialDiv.Find("ul").Each(func(idx int, sel *goquery.Selection) {
		title := doc.Find("div#" + tagName + "Bar li:nth-child(" + xnumber.Itoa(idx+1) + ")").Text() // <<<
		groupMangas := make([]*vo.MangaPageLink, 0)
		sel.Find("li").Each(func(idx int, li *goquery.Selection) {
			groupMangas = append(groupMangas, m.getMangaPageLinkFromLi(li))
		})
		groups = append(groups, &vo.MangaPageGroup{
			Title:  title,
			Mangas: groupMangas,
		})
	})

	// get other group mangas
	// TODO

	return groups, []*vo.MangaPageGroup{}
}

func (m *MangaListService) getMangaPageLinkFromLi(li *goquery.Selection) *vo.MangaPageLink {
	url := li.Find("a").AttrOr("href", "")
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	bid, _ := xnumber.Atou64(sp[len(sp)-1])
	title := li.Find("a").AttrOr("title", "")
	pic := li.Find("a img").AttrOr("src", "")
	if pic == "" {
		pic = li.Find("a img").AttrOr("data-src", "")
	}
	tt := li.Find("span.tt").Text()
	newestChapter := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(tt, "更新至"), "共"), "[全]")
	return &vo.MangaPageLink{
		Bid:           bid,
		Bname:         title,
		Bpic:          pic,
		Url:           static.HOMEPAGE_URL + url,
		Finished:      strings.HasPrefix(tt, "共"),
		NewestChapter: newestChapter,
	}
}

func (m *MangaListService) GetHotSerialMangas() (*vo.MangaGroupList, error) {
	doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL)
	if err != nil {
		return nil, err
	}

	groups, otherGroups := m.getMangas(doc, 1, "serial")
	return &vo.MangaGroupList{
		Title:       "热门连载",
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil
}

func (m *MangaListService) GetFinishedMangas() (*vo.MangaGroupList, error) {
	doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL)
	if err != nil {
		return nil, err
	}

	groups, otherGroups := m.getMangas(doc, 1, "serial")

	return &vo.MangaGroupList{
		Title:       "经典完结",
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil

}

func (m *MangaListService) GetLatestMangas() (*vo.MangaGroupList, error) {
	doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL)
	if err != nil {
		return nil, err
	}

	groups, otherGroups := m.getMangas(doc, 1, "serial")

	return &vo.MangaGroupList{
		Title:       "最新上架",
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil
}
