package service

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/param"
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

func (m *MangaListService) getMangas(doc *goquery.Document, tagIndex int, tagName string) (*vo.MangaPageGroup, []*vo.MangaPageGroup, []*vo.MangaPageGroup) {
	// get top mangas
	topMangas := make([]*vo.MangaPageLink, 0)
	topMangaUl := doc.Find("div.cmt-cont ul:nth-child(" + xnumber.Itoa(tagIndex) + ")") // <<<
	topMangaUl.Find("li").Each(func(idx int, li *goquery.Selection) {
		topMangas = append(topMangas, m.getMangaPageLinkFromLi(li, true))
	})
	topGroup := &vo.MangaPageGroup{
		Title:  "",
		Mangas: topMangas,
	}

	// get group mangas
	groups := make([]*vo.MangaPageGroup, 0)
	otherMangaUl := doc.Find("div#" + tagName + "Cont ul") // <<<
	otherMangaUl.Each(func(idx int, sel *goquery.Selection) {
		groupTitle := doc.Find("div#" + tagName + "Bar li:nth-child(" + xnumber.Itoa(idx+1) + ")").Text() // <<<
		groupMangas := make([]*vo.MangaPageLink, 0)
		sel.Find("li").Each(func(idx int, li *goquery.Selection) {
			groupMangas = append(groupMangas, m.getMangaPageLinkFromLi(li, true))
		})
		groups = append(groups, &vo.MangaPageGroup{
			Title:  groupTitle,
			Mangas: groupMangas,
		})
	})

	// get other group mangas
	otherGroups := make([]*vo.MangaPageGroup, 0)
	if tagIndex == 1 || tagIndex == 2 {
		scContDiv := doc.Find("div.idx-sc-cont")
		scContDiv.Each(func(idx int, sel *goquery.Selection) {
			groupMangas := make([]*vo.MangaPageLink, 0)
			groupTitle := sel.Find("h4").Text()
			otherMangaUl := sel.Find("div.idx-sc-list ul:nth-child(" + xnumber.Itoa(tagIndex) + ")") // <<<
			otherMangaUl.Children().Each(func(idx int, li *goquery.Selection) {
				manga := m.getMangaPageLinkFromLi(li, false)
				manga.Finished = tagIndex == 2
				groupMangas = append(groupMangas, manga)
			})
			otherGroups = append(otherGroups, &vo.MangaPageGroup{
				Title:  groupTitle,
				Mangas: groupMangas,
			})
		})
	}

	return topGroup, groups, otherGroups
}

func (m *MangaListService) getMangaPageLinkFromLi(li *goquery.Selection, hasCover bool) *vo.MangaPageLink {
	if hasCover {
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
		newestDate := li.Find("span.dt").Text()
		return &vo.MangaPageLink{
			Bid:           bid,
			Bname:         title,
			Bpic:          pic,
			Url:           static.HOMEPAGE_URL + url,
			Finished:      strings.HasPrefix(tt, "共"),
			NewestChapter: newestChapter,
			NewestDate:    newestDate,
		}
	} else {
		title := li.Find("h6 a").AttrOr("title", "")
		url := li.Find("h6 a").AttrOr("href", "")
		sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
		bid, _ := xnumber.Atou64(sp[len(sp)-1])
		newestChapter := li.Find("h6 span a").AttrOr("title", "")
		return &vo.MangaPageLink{
			Bid:           bid,
			Bname:         title,
			Bpic:          "",
			Url:           static.HOMEPAGE_URL + url,
			Finished:      true,
			NewestChapter: newestChapter,
		}
	}
}

func (m *MangaListService) GetHotSerialMangas() (*vo.MangaGroupList, error) {
	doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL)
	if err != nil {
		return nil, err
	}

	topGroup, groups, otherGroups := m.getMangas(doc, 1, "serial")
	return &vo.MangaGroupList{
		Title:       "热门连载",
		TopGroup:    topGroup,
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil
}

func (m *MangaListService) GetFinishedMangas() (*vo.MangaGroupList, error) {
	doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL)
	if err != nil {
		return nil, err
	}

	topGroup, groups, otherGroups := m.getMangas(doc, 2, "finish")
	return &vo.MangaGroupList{
		Title:       "经典完结",
		TopGroup:    topGroup,
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil

}

func (m *MangaListService) GetLatestMangas() (*vo.MangaGroupList, error) {
	doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL)
	if err != nil {
		return nil, err
	}

	topGroup, groups, otherGroups := m.getMangas(doc, 3, "latest")
	return &vo.MangaGroupList{
		Title:       "最新上架",
		TopGroup:    topGroup,
		Groups:      groups,
		OtherGroups: otherGroups,
	}, nil
}

func (m *MangaListService) GetUpdatedMangas(pa *param.PageParam) ([]*vo.MangaPageLink, int32, error) {
	doc, err := m.httpService.HttpGetDocument(static.MANGA_UPDATE_URL)
	if err != nil {
		return nil, 0, err
	}

	latestLis := doc.Find("div.latest-list li")
	allMangas := make([]*vo.MangaPageLink, latestLis.Length())
	latestLis.Each(func(idx int, li *goquery.Selection) {
		allMangas[idx] = m.getMangaPageLinkFromLi(li, true)
	})
	totalLength := int32(len(allMangas))

	out := make([]*vo.MangaPageLink, 0)
	start := pa.Limit * (pa.Page - 1)
	end := start + pa.Limit
	for i := start; i < end && i < totalLength; i++ {
		out = append(out, allMangas[i])
	}
	return out, totalLength, nil
}
