package service

import (
	"fmt"
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

func (m *MangaListService) getMangas(doc *goquery.Document, tagIndex int, tagName string) (*vo.MangaGroup, []*vo.MangaGroup, []*vo.MangaGroup) {
	// get top mangas
	topMangas := make([]*vo.TinyManga, 0)
	topMangaUl := doc.Find("div.cmt-cont ul:nth-child(" + xnumber.Itoa(tagIndex) + ")") // <<<
	topMangaUl.Find("li").Each(func(idx int, li *goquery.Selection) {
		topMangas = append(topMangas, m.getTinyMangaPageFromLi(li, true))
	})
	topGroup := &vo.MangaGroup{
		Title:  "",
		Mangas: m.tinyMangasToTinyBlockMangas(topMangas),
	}

	// get group mangas
	groups := make([]*vo.MangaGroup, 0)
	otherMangaUl := doc.Find("div#" + tagName + "Cont ul") // <<<
	otherMangaUl.Each(func(idx int, sel *goquery.Selection) {
		groupTitle := doc.Find("div#" + tagName + "Bar li:nth-child(" + xnumber.Itoa(idx+1) + ")").Text() // <<<
		groupMangas := make([]*vo.TinyManga, 0)
		sel.Find("li").Each(func(idx int, li *goquery.Selection) {
			groupMangas = append(groupMangas, m.getTinyMangaPageFromLi(li, true))
		})
		groups = append(groups, &vo.MangaGroup{
			Title:  groupTitle,
			Mangas: m.tinyMangasToTinyBlockMangas(groupMangas),
		})
	})

	// get other group mangas
	otherGroups := make([]*vo.MangaGroup, 0)
	if tagIndex == 1 || tagIndex == 2 {
		scContDiv := doc.Find("div.idx-sc-cont")
		scContDiv.Each(func(idx int, sel *goquery.Selection) {
			groupMangas := make([]*vo.TinyManga, 0)
			groupTitle := sel.Find("h4").Text()
			otherMangaUl := sel.Find("div.idx-sc-list ul:nth-child(" + xnumber.Itoa(tagIndex) + ")") // <<<
			otherMangaUl.Children().Each(func(idx int, li *goquery.Selection) {
				manga := m.getTinyMangaPageFromLi(li, false)
				manga.Finished = tagIndex == 2
				groupMangas = append(groupMangas, manga)
			})
			otherGroups = append(otherGroups, &vo.MangaGroup{
				Title:  groupTitle,
				Mangas: m.tinyMangasToTinyBlockMangas(groupMangas),
			})
		})
	}

	return topGroup, groups, otherGroups
}

func (m *MangaListService) getTinyMangaPageFromLi(li *goquery.Selection, hasCover bool) *vo.TinyManga {
	if hasCover {
		url := li.Find("a").AttrOr("href", "")
		title := li.Find("a").AttrOr("title", "")
		cover := li.Find("a img").AttrOr("src", "")
		if cover == "" {
			cover = li.Find("a img").AttrOr("data-src", "")
		}
		tt := li.Find("span.tt").Text()
		newestChapter := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(tt, "更新至"), "共"), "[全]")
		newestDate := li.Find("span.dt").Text()
		return &vo.TinyManga{
			Mid:           static.ParseMid(url),
			Title:         title,
			Cover:         static.ParseCoverUrl(cover),
			Url:           static.HOMEPAGE_URL + url,
			Finished:      strings.HasPrefix(tt, "共"),
			NewestChapter: newestChapter,
			NewestDate:    newestDate,
		}
	} else {
		title := li.Find("h6 a").AttrOr("title", "")
		url := li.Find("h6 a").AttrOr("href", "")
		newestChapter := li.Find("h6 span a").AttrOr("title", "")
		id := static.ParseMid(url)
		return &vo.TinyManga{
			Mid:           id,
			Title:         title,
			Cover:         fmt.Sprintf(static.MANGA_COVER_URL, id),
			Url:           static.HOMEPAGE_URL + url,
			Finished:      true, // true
			NewestChapter: newestChapter,
			NewestDate:    "",
		}
	}
}

func (m *MangaListService) tinyMangasToTinyBlockMangas(mangas []*vo.TinyManga) []*vo.TinyBlockManga {
	out := make([]*vo.TinyBlockManga, len(mangas))
	for idx, manga := range mangas {
		out[idx] = &vo.TinyBlockManga{
			Mid:           manga.Mid,
			Title:         manga.Title,
			Cover:         manga.Cover,
			Url:           manga.Url,
			Finished:      manga.Finished,
			NewestChapter: manga.NewestChapter,
		}
	}
	return out
}

func (m *MangaListService) GetHotSerialMangas() (*vo.MangaGroupList, error) {
	_, doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
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
	_, doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
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
	_, doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
	if err != nil {
		return nil, err
	}

	topGroup, groups, otherGroups := m.getMangas(doc, 3, "latest")
	return &vo.MangaGroupList{
		Title:       "最新上架",
		TopGroup:    topGroup,
		Groups:      groups,
		OtherGroups: otherGroups, // X
	}, nil
}

func (m *MangaListService) GetHomepageMangas() (*vo.HomepageMangaGroupList, error) {
	_, doc, err := m.httpService.HttpGetDocument(static.HOMEPAGE_URL, nil)
	if err != nil {
		return nil, err
	}

	sTopGroup, sGroups, sOtherGroups := m.getMangas(doc, 1, "serial")
	fTopGroup, fGroups, fOtherGroups := m.getMangas(doc, 2, "finish")
	lTopGroup, lGroups, lOtherGroups := m.getMangas(doc, 3, "latest")
	return &vo.HomepageMangaGroupList{
		Serial: &vo.MangaGroupList{Title: "热门连载", TopGroup: sTopGroup, Groups: sGroups, OtherGroups: sOtherGroups},
		Finish: &vo.MangaGroupList{Title: "经典完结", TopGroup: fTopGroup, Groups: fGroups, OtherGroups: fOtherGroups},
		Latest: &vo.MangaGroupList{Title: "最新上架", TopGroup: lTopGroup, Groups: lGroups, OtherGroups: lOtherGroups},
	}, nil
}

func (m *MangaListService) GetRecentUpdatedMangas(pa *param.PageParam) ([]*vo.TinyManga, int32, error) {
	_, doc, err := m.httpService.HttpGetDocument(static.MANGA_UPDATE_URL, nil)
	if err != nil {
		return nil, 0, err
	}

	latestLis := doc.Find("div.latest-list li")
	allMangas := make([]*vo.TinyManga, latestLis.Length())
	latestLis.Each(func(idx int, li *goquery.Selection) {
		allMangas[idx] = m.getTinyMangaPageFromLi(li, true)
	})
	totalLength := int32(len(allMangas))

	out := make([]*vo.TinyManga, 0)
	start := pa.Limit * (pa.Page - 1)
	end := start + pa.Limit
	for i := start; i < end && i < totalLength; i++ {
		out = append(out, allMangas[i])
	}
	return out, totalLength, nil
}
