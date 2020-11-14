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

func (s *SearchService) SearchMangas(keyword string, page int32, orderByPopular bool) ([]*vo.SmallMangaPage, int32, int32, error) {
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

	mangas := make([]*vo.SmallMangaPage, 0)
	if page <= pages {
		listLis := doc.Find("div.book-result li.cf")
		listLis.Each(func(idx int, li *goquery.Selection) {
			mangas = append(mangas, s.getSmallMangaPageFromLi(li))
		})
	}

	return mangas, limit, total, nil
}

func (s *SearchService) getSmallMangaPageFromLi(li *goquery.Selection) *vo.SmallMangaPage {
	title := li.Find("dt a").AttrOr("title", "")
	url := li.Find("dt a").AttrOr("href", "")
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	mid, _ := xnumber.Atou64(sp[len(sp)-1])
	cover := li.Find("div.book-cover img").AttrOr("src", "")
	statusDD := li.Find("dd.status")
	status := statusDD.Find("span:nth-child(2)").Text()
	newestDate := statusDD.Find("span:nth-child(3)").Text()
	newestChapter := statusDD.Find("a").Text()
	categoryDD := li.Find("div.book-detail dl dd.tags:nth-child(3)")
	publishYear := categoryDD.Find("span:nth-child(1) a").Text()
	mangaZone := categoryDD.Find("span:nth-child(2) a").AttrOr("title", "")
	genreA := categoryDD.Find("span:nth-child(3) a")
	genres := make([]*vo.Category, 0)
	genreA.Each(func(idx int, sel *goquery.Selection) {
		genres = append(genres, s.categoryService.GetCategoryFromA(sel))
	})
	authorA := li.Find("div.book-detail dl dd.tags:nth-child(4) a")
	authors := make([]*vo.TinyAuthor, 0)
	authorA.Each(func(idx int, sel *goquery.Selection) {
		authors = append(authors, s.authorService.GetAuthorFromA(sel))
	})
	briefIntroduction := strings.TrimSuffix(strings.TrimPrefix(li.Find("dd.intro").Text(), "简介："), "[详情]")

	return &vo.SmallMangaPage{
		Mid:               mid,
		Title:             title,
		Cover:             cover,
		Url:               static.HOMEPAGE_URL + url,
		PublishYear:       publishYear,
		MangaZone:         mangaZone,
		Genres:            genres,
		Authors:           authors,
		Finished:          status == "已完结",
		NewestChapter:     newestChapter,
		NewestDate:        strings.Split(newestDate, " ")[0],
		BriefIntroduction: briefIntroduction,
	}
}
