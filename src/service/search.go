package service

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
	"github.com/PuerkitoBio/goquery"
	"math"
	"strings"
)

type SearchService struct {
	httpService *HttpService
}

func NewSearchService() *SearchService {
	return &SearchService{
		httpService: xdi.GetByNameForce(sn.SHttpService).(*HttpService),
	}
}

func (s *SearchService) SearchMangas(keyword string, page int32, orderByPopular bool) ([]*vo.TinyMangaPage, int32, int32, error) {
	url := ""
	if orderByPopular {
		url = fmt.Sprintf(static.MANGA_SEARCH_URL, fmt.Sprintf("%s_o1", keyword), page)
	} else {
		url = fmt.Sprintf(static.MANGA_SEARCH_URL, keyword, page)
	}

	doc, err := s.httpService.HttpGetDocument(url)
	if err != nil {
		return nil, 0, 0, err
	} else if doc == nil {
		return nil, 0, 0, nil
	}

	limit := int32(10)
	total, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(2)").Text())
	pages := int32(math.Ceil(float64(total) / float64(limit)))

	mangas := make([]*vo.TinyMangaPage, 0)
	if page <= pages {
		listLis := doc.Find("div.book-result li.cf")
		listLis.Each(func(idx int, li *goquery.Selection) {
			title := li.Find("dt a").AttrOr("title", "")
			url := li.Find("dt a").AttrOr("href", "")
			sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
			mid, _ := xnumber.Atou64(sp[len(sp)-1])
			cover := li.Find("div.book-cover img").AttrOr("src", "")
			statusSpan := li.Find("dd.tags.status")
			status := statusSpan.Find("span:nth-child(2)").Text()
			newestDate := statusSpan.Find("span:nth-child(3)").Text()
			newestChapter := statusSpan.Find("a").Text()
			mangas = append(mangas, &vo.TinyMangaPage{
				Mid:           mid,
				Title:         title,
				Cover:         cover,
				Url:           static.HOMEPAGE_URL + url,
				Finished:      status == "已完结",
				NewestChapter: newestChapter,
				NewestDate:    strings.Split(newestDate, " ")[0],
			})
		})
	}

	return mangas, limit, total, nil
}
