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

type CategoryService struct {
	httpService *HttpService
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		httpService: xdi.GetByNameForce(sn.SHttpService).(*HttpService),
	}
}

func (c *CategoryService) GetGenres() ([]*vo.Category, error) {
	doc, err := c.httpService.HttpGetDocument(static.MANGA_CATEGORY_URL)
	if err != nil {
		return nil, err
	}

	categories := make([]*vo.Category, 0)
	genreLis := doc.Find("div.filter-nav div.filter.genre li")
	genreLis.Each(func(idx int, sel *goquery.Selection) {
		a := sel.Find("a")
		category := c.GetCategoryFromA(a)
		if category.Title != "全部" {
			categories = append(categories, category)
		}
	})

	return categories, nil
}

func (c *CategoryService) GetZones() ([]*vo.Category, error) {
	doc, err := c.httpService.HttpGetDocument(static.MANGA_CATEGORY_URL)
	if err != nil {
		return nil, err
	}

	categories := make([]*vo.Category, 0)
	genreLis := doc.Find("div.filter-nav div.filter.area li")
	genreLis.Each(func(idx int, sel *goquery.Selection) {
		a := sel.Find("a")
		category := c.GetCategoryFromA(a)
		if category.Title != "全部" {
			categories = append(categories, category)
		}
	})

	return categories, nil
}

func (c *CategoryService) GetAges() ([]*vo.Category, error) {
	doc, err := c.httpService.HttpGetDocument(static.MANGA_CATEGORY_URL)
	if err != nil {
		return nil, err
	}

	categories := make([]*vo.Category, 0)
	genreLis := doc.Find("div.filter-nav div.filter.age li")
	genreLis.Each(func(idx int, sel *goquery.Selection) {
		a := sel.Find("a")
		category := c.GetCategoryFromA(a)
		if category.Title != "全部" {
			categories = append(categories, category)
		}
	})

	return categories, nil
}

func (c *CategoryService) GetCategoryFromA(a *goquery.Selection) *vo.Category {
	title := a.Text()
	url := strings.TrimSuffix(a.AttrOr("href", ""), "/")
	sp := strings.Split(url, "/")
	name := sp[len(sp)-1]
	return &vo.Category{
		Name:  name,
		Title: title,
		Url:   static.HOMEPAGE_URL + url,
	}
}

func (c *CategoryService) GetGenreMangas(genre, zone, age, status string, page int32, orderByPopular bool) ([]*vo.TinyMangaPage, int32, int32, error) {
	url := static.MANGA_CATEGORY_URL + "/" // https://www.manhuagui.com/list/update_p1.html
	if zone != "" && zone != "all" {
		url += zone + "_" // https://www.manhuagui.com/list/japan/update.html
	}
	if genre != "" && genre != "all" {
		url += genre + "_" // https://www.manhuagui.com/list/japan_aiqing/update.html
	}
	if age != "" && age != "all" {
		url += age + "_" // https://www.manhuagui.com/list/japan_aiqing_shaonv/update.html
	}
	if status != "" && status != "all" {
		url += status + "_" // https://www.manhuagui.com/list/japan_aiqing_shaonv_lianzai/update_p1.html
	}
	url = strings.TrimSuffix(url, "_")
	if orderByPopular {
		url += fmt.Sprintf("/%s_p%d.html", "view", page)
	} else {
		url += fmt.Sprintf("/%s_p%d.html", "update", page)
	}

	doc, err := c.httpService.HttpGetDocument(url)
	if err != nil {
		return nil, 0, 0, err
	}
	if doc == nil {
		return nil, 0, 0, nil
	}

	limit := int32(42)
	pages, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(2)").Text())
	total, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(3)").Text())

	mangas := make([]*vo.TinyMangaPage, 0)
	if page <= pages {
		listLis := doc.Find("ul#contList li")
		listLis.Each(func(idx int, li *goquery.Selection) {
			mangas = append(mangas, c.getTinyMangaPageFromLi(li))
		})
	}

	return mangas, limit, total, nil
}

func (c *CategoryService) getTinyMangaPageFromLi(li *goquery.Selection) *vo.TinyMangaPage {
	url := li.Find("a").AttrOr("href", "")
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	mid, _ := xnumber.Atou64(sp[len(sp)-1])
	title := li.Find("a").AttrOr("title", "")
	cover := li.Find("a img").AttrOr("src", "")
	if cover == "" {
		cover = li.Find("a img").AttrOr("data-src", "")
	}
	tt := li.Find("span.tt").Text()
	newestChapter := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(tt, "更新至"), "共"), "[完]")
	score := li.Find("span.updateon em").Text()
	newestDate := strings.TrimPrefix(strings.TrimSuffix(li.Find("span.updateon").Text(), score), "更新于：")
	return &vo.TinyMangaPage{
		Mid:           mid,
		Title:         title,
		Cover:         cover,
		Url:           static.HOMEPAGE_URL + url,
		Finished:      strings.HasSuffix(tt, "[完]"),
		NewestChapter: newestChapter,
		NewestDate:    strings.TrimSpace(newestDate),
	}
}
