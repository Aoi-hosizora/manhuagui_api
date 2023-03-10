package service

import (
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/object"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/static"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type CategoryService struct {
	httpService *HttpService
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		httpService: xmodule.MustGetByName(sn.SHttpService).(*HttpService),
	}
}

func (c *CategoryService) GetAllCategories() (*object.CategoryList, error) {
	_, doc, err := c.httpService.HttpGetDocument(fmt.Sprintf(static.MANGA_RANK_URL, ""), nil)
	if err != nil {
		return nil, err
	}
	return c.getAllCategories(doc), nil
}

func (c *CategoryService) getAllCategories(doc *goquery.Document) *object.CategoryList {
	zones := make([]*object.Category, 0)
	zoneLis := doc.Find("div.category-list ul:nth-of-type(2) li")
	zoneLis.Each(func(i int, sel *goquery.Selection) {
		a := sel.Find("a")
		category := c.getCategoryFromA(a)
		if category.Title != "全部" {
			zones = append(zones, category)
		}
	})

	ages := make([]*object.Category, 0)
	ageLis := doc.Find("div.category-list ul:nth-of-type(3) li")
	ageLis.Each(func(i int, sel *goquery.Selection) {
		a := sel.Find("a")
		category := c.getCategoryFromA(a)
		if category.Title != "全部" {
			ages = append(ages, category)
		}
	})

	genres := make([]*object.Category, 0)
	genresLis := doc.Find("div.category-list ul:nth-of-type(4) li")
	genresLis.Each(func(i int, sel *goquery.Selection) {
		a := sel.Find("a")
		category := c.getCategoryFromA(a)
		if category.Title != "全部" {
			genres = append(genres, category)
		}
	})

	out := &object.CategoryList{
		Genres: genres,
		Zones:  zones,
		Ages:   ages,
	}
	return out
}

func (c *CategoryService) GetGenres() ([]*object.Category, error) {
	_, doc, err := c.httpService.HttpGetDocument(fmt.Sprintf(static.MANGA_RANK_URL, ""), nil)
	if err != nil {
		return nil, err
	}
	return c.getAllCategories(doc).Genres, nil
}

func (c *CategoryService) GetZones() ([]*object.Category, error) {
	_, doc, err := c.httpService.HttpGetDocument(fmt.Sprintf(static.MANGA_RANK_URL, ""), nil)
	if err != nil {
		return nil, err
	}
	return c.getAllCategories(doc).Zones, nil
}

func (c *CategoryService) GetAges() ([]*object.Category, error) {
	_, doc, err := c.httpService.HttpGetDocument(fmt.Sprintf(static.MANGA_RANK_URL, ""), nil)
	if err != nil {
		return nil, err
	}
	return c.getAllCategories(doc).Ages, nil
}

func (c *CategoryService) getCategoryFromA(a *goquery.Selection) *object.Category {
	title := a.Text()
	url := strings.TrimSuffix(strings.TrimSuffix(a.AttrOr("href", ""), ".html"), "/")
	sp := strings.Split(url, "/")
	name := sp[len(sp)-1]
	return &object.Category{
		Name:  name,
		Title: title,
		Url:   static.HOMEPAGE_URL + url,
		Cover: c.matchCoverForCategory(name),
	}
}

func (c *CategoryService) GetGenreMangas(genre, zone, age, status string, page int32, order string) ([]*object.TinyManga, int32, int32, error) {
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
	if order == "popular" {
		url += fmt.Sprintf("/%s_p%d.html", "view", page)
	} else if order == "new" {
		url += fmt.Sprintf("/%s_p%d.html", "index", page)
	} else { // update
		url += fmt.Sprintf("/%s_p%d.html", "update", page)
	}

	bs, doc, err := c.httpService.HttpGetDocument(url, nil)
	if err != nil {
		return nil, 0, 0, err
	} else if doc == nil {
		return nil, 0, 0, nil
	} else if bytes.Contains(bs, []byte(static.NOT_FOUND2_TOKEN)) {
		return []*object.TinyManga{}, 0, 0, nil
	}

	limit := int32(42)
	pages, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(2)").Text())
	total, _ := xnumber.Atoi32(doc.Find("div.result-count strong:nth-child(3)").Text())

	mangas := make([]*object.TinyManga, 0)
	if page <= pages {
		listLis := doc.Find("ul#contList li")
		listLis.Each(func(idx int, li *goquery.Selection) {
			mangas = append(mangas, c.getTinyMangaPageFromLi(li))
		})
	}

	return mangas, limit, total, nil
}

func (c *CategoryService) getTinyMangaPageFromLi(li *goquery.Selection) *object.TinyManga {
	url := li.Find("a").AttrOr("href", "")
	title := li.Find("a").AttrOr("title", "")
	cover := li.Find("a img").AttrOr("src", "")
	if cover == "" {
		cover = li.Find("a img").AttrOr("data-src", "")
	}
	tt := li.Find("span.tt").Text()
	newestChapter := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(tt, "更新至"), "共"), "[完]")
	score := li.Find("span.updateon em").Text()
	newestDate := strings.TrimPrefix(strings.TrimSuffix(li.Find("span.updateon").Text(), score), "更新于：")
	return &object.TinyManga{
		Mid:           static.ParseMid(url),
		Title:         title,
		Cover:         static.ParseCoverUrl(cover),
		Url:           static.HOMEPAGE_URL + url,
		Finished:      strings.HasSuffix(tt, "[完]"),
		NewestChapter: newestChapter,
		NewestDate:    strings.TrimSpace(newestDate),
	}
}

var _categoryCovers = map[string]string{
	"all": "", // 全部
	//
	"rexue":     "https://cf.hamreus.com/cpic/g/1128.jpg",     // 热血 https://www.manhuagui.com/comic/1128/
	"maoxian":   "https://cf.hamreus.com/cpic/g/19430.jpg",    // 冒险 https://www.manhuagui.com/comic/19430/
	"mohuan":    "https://cf.hamreus.com/cpic/g/17023_24.jpg", // 魔幻 https://www.manhuagui.com/comic/17023/
	"shengui":   "https://cf.hamreus.com/cpic/g/19397_39.jpg", // 神鬼 https://www.manhuagui.com/comic/19397/
	"gaoxiao":   "https://cf.hamreus.com/cpic/g/7507_21.jpg",  // 搞笑 https://www.manhuagui.com/comic/7507/
	"mengxi":    "https://cf.hamreus.com/cpic/g/8904.jpg",     // 萌系 https://www.manhuagui.com/comic/8904/
	"aiqing":    "https://cf.hamreus.com/cpic/g/17332_52.jpg", // 爱情 https://www.manhuagui.com/comic/17332/
	"kehuan":    "https://cf.hamreus.com/cpic/g/1676_14.jpg",  // 科幻 https://www.manhuagui.com/comic/1676/
	"mofa":      "https://cf.hamreus.com/cpic/g/29168.jpg",    // 魔法 https://www.manhuagui.com/comic/29168/
	"gedou":     "https://cf.hamreus.com/cpic/g/7580.jpg",     // 格斗 https://www.manhuagui.com/comic/7580/
	"wuxia":     "https://cf.hamreus.com/cpic/g/7620.jpg",     // 武侠 https://www.manhuagui.com/comic/7620/
	"jizhan":    "https://cf.hamreus.com/cpic/g/20562.jpg",    // 机战 https://www.manhuagui.com/comic/20562/
	"zhanzheng": "https://cf.hamreus.com/cpic/g/10528.jpg",    // 战争 https://www.manhuagui.com/comic/10528/
	"jingji":    "https://cf.hamreus.com/cpic/g/32303.jpg",    // 竞技 https://www.manhuagui.com/comic/32303/
	"tiyu":      "https://cf.hamreus.com/cpic/g/4721.jpg",     // 体育 https://www.manhuagui.com/comic/4721/
	"xiaoyuan":  "https://cf.hamreus.com/cpic/g/22942_25.jpg", // 校园 https://www.manhuagui.com/comic/22942/
	"shenghuo":  "https://cf.hamreus.com/cpic/g/25882.jpg",    // 生活 https://www.manhuagui.com/comic/25882/
	"lizhi":     "https://cf.hamreus.com/cpic/g/4779.jpg",     // 励志 https://www.manhuagui.com/comic/4779/
	"lishi":     "https://cf.hamreus.com/cpic/g/1147.jpg",     // 历史 https://www.manhuagui.com/comic/1147/
	"weiniang":  "https://cf.hamreus.com/cpic/g/26682.jpg",    // 伪娘 https://www.manhuagui.com/comic/26682/
	"zhainan":   "https://cf.hamreus.com/cpic/g/17217_02.jpg", // 宅男 https://www.manhuagui.com/comic/17217/
	"funv":      "https://cf.hamreus.com/cpic/g/20135.jpg",    // 腐女 https://www.manhuagui.com/comic/20135/
	"danmei":    "https://cf.hamreus.com/cpic/g/15667_09.jpg", // 耽美 https://www.manhuagui.com/comic/15667/
	"baihe":     "https://cf.hamreus.com/cpic/g/17201.jpg",    // 百合 https://www.manhuagui.com/comic/17201/
	"hougong":   "https://cf.hamreus.com/cpic/g/17596_35.jpg", // 后宫 https://www.manhuagui.com/comic/17596/
	"zhiyu":     "https://cf.hamreus.com/cpic/g/31550_19.jpg", // 治愈 https://www.manhuagui.com/comic/31550/
	"meishi":    "https://cf.hamreus.com/cpic/g/2863.jpg",     // 美食 https://www.manhuagui.com/comic/2863/
	"tuili":     "https://cf.hamreus.com/cpic/g/4383.jpg",     // 推理 https://www.manhuagui.com/comic/4383/
	"xuanyi":    "https://cf.hamreus.com/cpic/g/17726.jpg",    // 悬疑 https://www.manhuagui.com/comic/17726/
	"kongbu":    "https://cf.hamreus.com/cpic/g/29499_74.jpg", // 恐怖 https://www.manhuagui.com/comic/29499/
	"sige":      "https://cf.hamreus.com/cpic/g/24846.jpg",    // 四格 https://www.manhuagui.com/comic/24846/
	"zhichang":  "https://cf.hamreus.com/cpic/g/35634.jpg",    // 职场 https://www.manhuagui.com/comic/35634/
	"zhentan":   "https://cf.hamreus.com/cpic/g/15581.jpg",    // 侦探 https://www.manhuagui.com/comic/15581/
	"shehui":    "https://cf.hamreus.com/cpic/g/44866.jpg",    // 社会 https://www.manhuagui.com/comic/44866/
	"yinyue":    "https://cf.hamreus.com/cpic/g/32263.jpg",    // 音乐 https://www.manhuagui.com/comic/32263/
	"wudao":     "https://cf.hamreus.com/cpic/g/9426.jpg",     // 舞蹈 https://www.manhuagui.com/comic/9426/
	"zazhi":     "https://cf.hamreus.com/cpic/g/7891.jpg",     // 杂志 https://www.manhuagui.com/comic/7891/
	"heidao":    "https://cf.hamreus.com/cpic/g/8928.jpg",     // 黑道 https://www.manhuagui.com/comic/8928/
	//
	"shaonv":   "https://cf.hamreus.com/cpic/g/883.jpg",      // 少女 https://www.manhuagui.com/comic/883/
	"shaonian": "https://cf.hamreus.com/cpic/g/23394.jpg",    // 少年 https://www.manhuagui.com/comic/23394/
	"qingnian": "https://cf.hamreus.com/cpic/g/4740.jpg",     // 青年 https://www.manhuagui.com/comic/4740/
	"ertong":   "https://cf.hamreus.com/cpic/g/9623.jpg",     // 儿童 https://www.manhuagui.com/comic/9623/
	"tongyong": "https://cf.hamreus.com/cpic/g/21269_34.jpg", // 通用 https://www.manhuagui.com/comic/21269/
	//
	"japan":    "https://cf.hamreus.com/cpic/g/24591_13.jpg", // 日本 https://www.manhuagui.com/comic/24591/
	"hongkong": "https://cf.hamreus.com/cpic/g/20001.jpg",    // 港台 https://www.manhuagui.com/comic/20001/
	"other":    "https://cf.hamreus.com/cpic/g/16240.jpg",    // 其它 https://www.manhuagui.com/comic/16240/
	"europe":   "https://cf.hamreus.com/cpic/g/37514.jpg",    // 欧美 https://www.manhuagui.com/comic/37514/
	"china":    "https://cf.hamreus.com/cpic/g/7382.jpg",     // 内地 https://www.manhuagui.com/comic/7382/
	"korea":    "https://cf.hamreus.com/cpic/g/32706.jpg",    // 韩国 https://www.manhuagui.com/comic/32706/
}

func (c *CategoryService) matchCoverForCategory(name string) string {
	cover, ok := _categoryCovers[name]
	if !ok {
		return ""
	}
	return cover
}
