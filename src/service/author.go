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

type AuthorService struct {
	httpService     *HttpService
	categoryService *CategoryService
}

func NewAuthorService() *AuthorService {
	return &AuthorService{
		httpService:     xdi.GetByNameForce(sn.SHttpService).(*HttpService),
		categoryService: xdi.GetByNameForce(sn.SCategoryService).(*CategoryService),
	}
}

func (a *AuthorService) GetAuthorFromA(sel *goquery.Selection) *vo.TinyAuthor {
	name := sel.AttrOr("title", "")
	url := strings.TrimSuffix(sel.AttrOr("href", ""), "/")
	return &vo.TinyAuthor{
		Aid:  static.ParseAid(url),
		Name: name,
		Url:  strings.TrimSuffix(static.HOMEPAGE_URL+url, "/"),
	}
}

func (a *AuthorService) GetAllAuthors(genre, zone, age string, page int32, order string) ([]*vo.SmallAuthor, int32, int32, error) {
	url := static.MANGA_AUTHORS_URL + "/" // https://www.manhuagui.com/alist/index_p1.html
	if zone != "" && zone != "all" {
		url += zone + "_" // https://www.manhuagui.com/alist/japan/update.html
	}
	if genre != "" && genre != "all" {
		url += genre + "_" // https://www.manhuagui.com/alist/japan_aiqing/update.html
	}
	if age != "" && age != "all" {
		url += age + "_" // https://www.manhuagui.com/alist/japan_aiqing_shaonv/update.html
	}
	url = strings.TrimSuffix(url, "_")
	if order == "popular" {
		url += fmt.Sprintf("/%s_p%d.html", "view", page)
	} else if order == "comic" {
		url += fmt.Sprintf("/%s_p%d.html", "comic", page)
	} else { // update
		url += fmt.Sprintf("/%s_p%d.html", "index", page)
	}

	_, doc, err := a.httpService.HttpGetDocument(url, nil)
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
	zone := li.Find("font:nth-child(2)").Text()
	mangaCount, _ := xnumber.Atoi32(li.Find("font:nth-child(3)").Text())
	score := li.Find("span.updateon em").Text()
	newestDate := strings.TrimPrefix(strings.TrimSuffix(li.Find("span.updateon").Text(), score), "更新于：")
	return &vo.SmallAuthor{
		Aid:        static.ParseAid(url),
		Name:       name,
		Zone:       zone,
		Cover:      "https://cf.hamreus.com/zpic/none.jpg", // <<<
		Url:        strings.TrimSuffix(static.HOMEPAGE_URL+url, "/"),
		MangaCount: mangaCount,
		NewestDate: newestDate,
	}
}

func (a *AuthorService) GetAuthor(aid uint64) (*vo.Author, error) {
	url := fmt.Sprintf(static.MANGA_AUTHOR_URL, aid, "index", 1)
	_, doc, err := a.httpService.HttpGetDocument(url, nil)
	if err != nil {
		return nil, err
	} else if doc == nil {
		return nil, nil
	}

	name := doc.Find("div.title h1").Text()
	infoDiv := doc.Find("div.info")
	alias := strings.TrimPrefix(infoDiv.Find("p:nth-child(1)").Text(), "作者别名：")
	zone := strings.TrimPrefix(infoDiv.Find("p:nth-child(2)").Text(), "所属地区：")
	cover := doc.Find("div.info_cover img").AttrOr("src", "")
	newestP := infoDiv.Find("p:nth-child(4)")
	newestMangaUrl := newestP.Find("a").AttrOr("href", "")
	newestMangaTitle := strings.TrimSuffix(strings.TrimPrefix(newestP.Find("a").AttrOr("title", ""), "【"), "】")
	newestDate := newestP.Find("span").Text()
	mangaCount, _ := xnumber.Atoi32(infoDiv.Find("p:nth-child(5) font").Text())
	averageScore, _ := xnumber.Atof32(infoDiv.Find("p:nth-child(7) span").Text())
	introduction := doc.Find("div#intro-all h2").Text()
	out := &vo.Author{
		Aid:              aid,
		Name:             name,
		Zone:             zone,
		Cover:            cover,
		Url:              url,
		MangaCount:       mangaCount,
		NewestMangaId:    static.ParseMid(newestMangaUrl),
		NewestMangaTitle: newestMangaTitle,
		NewestDate:       newestDate,
		Alias:            alias,
		AverageScore:     averageScore,
		Introduction:     introduction,
	}

	return out, nil
}

func (a *AuthorService) GetAuthorMangas(aid uint64, page int32, order string) ([]*vo.SmallManga, int32, int32, error) {
	url := ""
	if order == "popular" {
		url = fmt.Sprintf(static.MANGA_AUTHOR_URL, aid, "view", page)
	} else if order == "new" {
		url = fmt.Sprintf(static.MANGA_AUTHOR_URL, aid, "new", page)
	} else { // update
		url = fmt.Sprintf(static.MANGA_AUTHOR_URL, aid, "index", page)
	}

	_, doc, err := a.httpService.HttpGetDocument(url, nil)
	if err != nil {
		return nil, 0, 0, err
	} else if doc == nil {
		return nil, 0, 0, nil
	}

	limit := int32(10)
	total, _ := xnumber.Atoi32(doc.Find("div.result-count strong").Text())
	pages := int32(math.Ceil(float64(total) / float64(limit)))

	mangas := make([]*vo.SmallManga, 0)
	if page <= pages {
		listLis := doc.Find("div.book-result li.cf")
		listLis.Each(func(idx int, li *goquery.Selection) {
			mangas = append(mangas, a.GetSmallMangaPageFromLi(li))
		})
	}

	return mangas, limit, total, nil
}

func (a *AuthorService) GetSmallMangaPageFromLi(li *goquery.Selection) *vo.SmallManga {
	title := li.Find("dt a").AttrOr("title", "")
	url := li.Find("dt a").AttrOr("href", "")
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
		genres = append(genres, a.categoryService.GetCategoryFromA(sel))
	})
	authorA := li.Find("div.book-detail dl dd.tags:nth-child(4) a")
	authors := make([]*vo.TinyAuthor, 0)
	authorA.Each(func(idx int, sel *goquery.Selection) {
		authors = append(authors, a.GetAuthorFromA(sel))
	})
	briefIntroduction := strings.TrimSuffix(strings.TrimPrefix(li.Find("dd.intro").Text(), "简介："), "[详情]")

	return &vo.SmallManga{
		Mid:               static.ParseMid(url),
		Title:             title,
		Cover:             static.ParseCoverUrl(cover),
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
