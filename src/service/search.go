package service

import (
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
	"github.com/PuerkitoBio/goquery"
	"math"
)

type SearchService struct {
	httpService     *HttpService
	categoryService *CategoryService
	authorService   *AuthorService
}

func NewSearchService() *SearchService {
	return &SearchService{
		httpService:     xdi.GetByNameForce(sn.SHttpService).(*HttpService),
		categoryService: xdi.GetByNameForce(sn.SCategoryService).(*CategoryService),
		authorService:   xdi.GetByNameForce(sn.SAuthorService).(*AuthorService),
	}
}

func (s *SearchService) SearchMangas(keyword string, page int32, order string) ([]*vo.SmallManga, int32, int32, error) {
	url := ""
	if order == "popular" {
		url = fmt.Sprintf(static.MANGA_SEARCH_URL, fmt.Sprintf("%s_o1", keyword), page)
	} else if order == "new" {
		url = fmt.Sprintf(static.MANGA_SEARCH_URL, fmt.Sprintf("%s_o2", keyword), page)
	} else { // update
		url = fmt.Sprintf(static.MANGA_SEARCH_URL, keyword, page)
	}

	bs, doc, err := s.httpService.HttpGetDocument(url)
	if err != nil {
		return nil, 0, 0, err
	} else if doc == nil {
		return nil, 0, 0, nil
	} else if bytes.Contains(bs, []byte(static.NOT_FOUND3_TOKEN)) {
		return nil, 0, 0, nil
	}

	limit := int32(10)
	total, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(2)").Text())
	pages := int32(math.Ceil(float64(total) / float64(limit)))

	mangas := make([]*vo.SmallManga, 0)
	if page <= pages {
		listLis := doc.Find("div.book-result li.cf")
		listLis.Each(func(idx int, li *goquery.Selection) {
			mangas = append(mangas, s.authorService.GetSmallMangaPageFromLi(li))
		})
	}

	return mangas, limit, total, nil
}
