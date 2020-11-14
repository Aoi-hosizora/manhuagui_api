package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
	"github.com/Aoi-hosizora/manhuagui-backend/src/util"
	"github.com/PuerkitoBio/goquery"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type MangaService struct {
	httpService     *HttpService
	categoryService *CategoryService
}

func NewMangaService() *MangaService {
	return &MangaService{
		httpService:     xdi.GetByNameForce(sn.SHttpService).(*HttpService),
		categoryService: xdi.GetByNameForce(sn.SCategoryService).(*CategoryService),
	}
}

func (m *MangaService) GetMangaPage(mid uint64) (*vo.MangaPage, error) {
	// get document
	url := fmt.Sprintf(static.MANGA_PAGE_URL, mid)
	doc, err := m.httpService.HttpGetDocument(url)
	if err != nil {
		return nil, err
	}
	if doc == nil {
		return nil, nil
	}

	// get basic information
	title := doc.Find("div.book-title").Text()
	cover := doc.Find("p.hcover img").AttrOr("src", "")
	detailUl := doc.Find("ul.detail-list")
	publishYear := detailUl.Find("li:nth-child(1) span:nth-child(1) a").Text()
	mangaZone := detailUl.Find("li:nth-child(1) span:nth-child(2) a").Text()
	alphabetIndex := detailUl.Find("li:nth-child(1) span:nth-child(3) a").Text()
	authorName := detailUl.Find("li:nth-child(2) span:nth-child(2) a").Text()
	alias := detailUl.Find("li:nth-child(3) span:nth-child(1)").Text()
	status := detailUl.Find("li:nth-child(4) span:nth-child(2)").Text()
	newestChapter := detailUl.Find("li:nth-child(4) a").Text()
	newestDate := detailUl.Find("li:nth-child(4) span:nth-child(3)").Text()
	introduction := doc.Find("div#intro-all").Text()
	mangaRank := doc.Find("div.rank").Text()
	genreA := detailUl.Find("li:nth-child(2) span:nth-child(1) a")
	genres := make([]*vo.Category, 0)
	genreA.Each(func(idx int, sel *goquery.Selection) {
		genres = append(genres, m.categoryService.GetCategoryFromA(sel))
	})
	obj := &vo.MangaPage{
		Mid:           mid,
		Title:         title,
		Cover:         cover,
		Url:           url,
		PublishYear:   publishYear,
		MangaZone:     mangaZone,
		AlphabetIndex: alphabetIndex,
		Genres:        genres,
		AuthorName:    authorName,
		Alias:         strings.TrimPrefix(alias, "漫画别名："),
		Finished:      status == "已完结",
		NewestChapter: newestChapter,
		NewestDate:    newestDate,
		Introduction:  strings.TrimSpace(introduction),
		MangaRank:     mangaRank,
	}

	// get score
	scoreUrl := fmt.Sprintf(static.MANGA_SCORE_URL, mid)
	scoreJsonBs, err := m.httpService.HttpGet(scoreUrl)
	if err != nil {
		return nil, err
	}
	scoreMap := make(map[string]interface{})
	err = json.Unmarshal(scoreJsonBs, &scoreMap)
	if err != nil {
		return nil, err
	}
	scoreErr := fmt.Errorf("failed to get score result")
	if resultItf, ok := scoreMap["success"]; !ok {
		return nil, scoreErr
	} else if result, ok := resultItf.(bool); !ok || !result {
		return nil, scoreErr
	}
	if scoreJsonItf, ok := scoreMap["data"]; !ok {
		return nil, scoreErr
	} else if scoreJson, ok := scoreJsonItf.(map[string]interface{}); !ok {
		return nil, scoreErr
	} else {
		s1 := scoreJson["s1"].(float64)
		s2 := scoreJson["s2"].(float64)
		s3 := scoreJson["s3"].(float64)
		s4 := scoreJson["s4"].(float64)
		s5 := scoreJson["s5"].(float64)
		tot := s1 + s2 + s3 + s4 + s5
		avg := (s1*1 + s2*2 + s3*3 + s4*4 + s5*5) / (tot * 5)
		per1 := float32(math.Round((s1/tot)*1000) / 1000) // 00.0%
		per2 := float32(math.Round((s2/tot)*1000) / 1000)
		per3 := float32(math.Round((s3/tot)*1000) / 1000)
		per4 := float32(math.Round((s4/tot)*1000) / 1000)
		per5 := float32(math.Round((s5/tot)*1000) / 1000)

		obj.ScoreCount = int32(tot)
		obj.AverageScore = float32(math.Round(avg*100) / 10) // 0.0
		obj.PerScores = [6]float32{0, per1, per2, per3, per4, per5}
	}

	// get chapter groups
	groupTitleH4s := doc.Find("div.chapter h4").Children()
	groupListDivs := doc.Find("div.chapter div.chapter-list")
	groupTitles := make([]string, groupTitleH4s.Length())
	groups := make([]*vo.MangaChapterGroup, len(groupTitles))
	groupTitleH4s.Each(func(idx int, sel *goquery.Selection) {
		groupTitles[idx] = sel.Text()
	})
	groupListDivs.Each(func(idx int, sel *goquery.Selection) {
		chapters := make([]*vo.TinyMangaChapter, 0)
		sel.Find("ul").Each(func(idx int, sel *goquery.Selection) {
			chaptersInUl := make([]*vo.TinyMangaChapter, 0)
			sel.Find("li").Each(func(idx int, sel *goquery.Selection) {
				title := sel.Find("a").AttrOr("title", "")
				pageCount, _ := xnumber.Atoi32(strings.TrimSuffix(sel.Find("i").Text(), "p"))
				url := sel.Find("a").AttrOr("href", "")
				sp := strings.Split(url, "/")
				cid, _ := xnumber.Atou64(strings.TrimSuffix(sp[len(sp)-1], ".html"))
				em := sel.Find("em").AttrOr("class", "")
				chaptersInUl = append(chaptersInUl, &vo.TinyMangaChapter{
					Cid:       cid,
					Title:     title,
					Mid:       mid,
					Url:       static.HOMEPAGE_URL + url,
					PageCount: pageCount,
					IsNew:     em != "",
				})
			})
			chapters = append(chaptersInUl, chapters...)
		})
		groups[idx] = &vo.MangaChapterGroup{
			Title:    groupTitles[idx],
			Chapters: chapters,
		}
	})
	obj.ChapterGroups = groups

	// return
	return obj, nil
}

func (m *MangaService) GetMangaChapter(mid, cid uint64) (*vo.MangaChapter, error) {
	// get document
	url := fmt.Sprintf(static.MANGA_CHAPTER_URL, mid, cid)
	doc, err := m.httpService.HttpGetDocument(url)
	if err != nil {
		return nil, err
	}
	if doc == nil {
		return nil, nil
	}

	// get script
	var script *goquery.Selection
	doc.Find("script").Each(func(i int, sel *goquery.Selection) {
		_, ok := sel.Attr("src")
		if !ok && strings.Contains(sel.Text(), `window["\x65\x76\x61\x6c"]`) {
			script = sel
		}
	})
	if script == nil {
		return nil, fmt.Errorf("script error: script not found")
	}

	// decode json
	decodeScript := script.Text()
	decodeParameterResults := regexp.MustCompile(`;return p;}\('(.+\(\);)',(.+?),(.+),'(.+)'\['\\x73`).FindAllStringSubmatch(decodeScript, 1)
	if len(decodeParameterResults) == 0 {
		return nil, fmt.Errorf("script error: could not find parameter")
	}
	p := decodeParameterResults[0][1]                  // XX.YY({...}).ZZ();
	a, _ := strconv.Atoi(decodeParameterResults[0][2]) // 00
	c, _ := strconv.Atoi(decodeParameterResults[0][3]) // 00
	k := decodeParameterResults[0][4]                  // ~['\x73\x70\x6c\x69\x63']('\x7c')
	decodeString := m.decodeChapterScript(p, a, c, k)
	decodeJsonResults := regexp.MustCompile(`SMH.imgData\((.+)\).preInit`).FindAllStringSubmatch(decodeString, 1)
	if len(decodeJsonResults) == 0 {
		return nil, fmt.Errorf("script error: invalid script")
	}
	decodeJson := decodeJsonResults[0][1]

	// unmarshal
	obj := &vo.MangaChapter{}
	err = json.Unmarshal([]byte(decodeJson), &obj)
	if err != nil {
		return nil, fmt.Errorf("chapter script error: %v", err)
	}
	obj.Url = url
	for idx := range obj.Pages {
		// 自动 h: "i" (100) | "us" (1)
		// 电信 h: "eu" (100) | "i" (1) | "us" (1)
		// 联通 h: "us" (100) | "i" (1) | "eu" (1)
		obj.Pages[idx] = fmt.Sprintf("%s%s%s?e=%d&m=%s", fmt.Sprintf(static.MANGA_SOURCE_URL, "i"), obj.Path, obj.Pages[idx], obj.Sl.E, obj.Sl.M)
	}

	// return
	return obj, nil
}

func (m *MangaService) decodeChapterScript(p string, a, c int, k string) string {
	k, err := util.DecompressLZStringFromBase64(k)
	if err != nil {
		return ""
	}
	ks := strings.Split(k, "|")

	// e=function(c){return(c<a?"":e(parseInt(c/a)))+((c=c%a)>35?String.fromCharCode(c+29):c.toString(36))};
	var e func(c int) string
	e = func(c int) string {
		p1 := ""
		if c >= a {
			p1 = e(c / a)
		}
		c = c % a
		p2 := byte(c + 29)
		if c <= 35 {
			p2 = "0123456789abcdefghijklmnopqrstuvwxyz"[c]
		}
		return p1 + string(p2)
	}

	// while(c--)d[e(c)]=k[c]||e(c);k=[function(e){return d[e]}];e=function(){return'\\w+'};c=1;
	d := map[string]string{}
	for ; c >= 0; c-- {
		if len(ks) > c && ks[c] != "" {
			d[e(c)] = ks[c]
		} else {
			d[e(c)] = e(c)
		}
	}

	// while(c--)if(k[c])p=p.replace(new RegExp('\\b'+e(c)+'\\b','g'),k[c]);return p;}
	p = regexp.MustCompile(`\b\w+\b`).ReplaceAllStringFunc(p, func(e string) string {
		return d[e]
	})
	return p
}
