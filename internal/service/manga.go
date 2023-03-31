package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xconstant/headers"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/object"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/lzstring"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/static"
	"github.com/PuerkitoBio/goquery"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type MangaService struct {
	httpService     *HttpService
	categoryService *CategoryService
	authorService   *AuthorService
}

func NewMangaService() *MangaService {
	return &MangaService{
		httpService:     xmodule.MustGetByName(sn.SHttpService).(*HttpService),
		categoryService: xmodule.MustGetByName(sn.SCategoryService).(*CategoryService),
		authorService:   xmodule.MustGetByName(sn.SAuthorService).(*AuthorService),
	}
}

func (m *MangaService) GetMangaPage(mid uint64) (*object.Manga, error) {
	// get document
	url := fmt.Sprintf(static.MANGA_PAGE_URL, mid)
	bs, doc, err := m.httpService.HttpGetDocument(url, nil)
	if err != nil {
		return nil, err
	}
	if doc == nil {
		return nil, nil
	}

	// get basic information
	title := doc.Find("div.book-title h1").Text()
	aliasTitle := doc.Find("div.book-title h2").Text()
	cover := doc.Find("p.hcover img").AttrOr("src", "")
	detailUl := doc.Find("ul.detail-list")
	publishYear := detailUl.Find("li:nth-child(1) span:nth-child(1) a").Text()
	mangaZone := detailUl.Find("li:nth-child(1) span:nth-child(2) a").Text()
	alphabetIndex := detailUl.Find("li:nth-child(1) span:nth-child(3) a").Text()
	mangaAliases := make([]string, 0)
	detailUl.Find("li:nth-child(3) span:nth-child(1) a").Each(func(i int, sel *goquery.Selection) {
		mangaAliases = append(mangaAliases, sel.Text())
	})
	aliases := make([]string, 0) // <= 别名
	if aliasTitle != "" {
		aliases = append(aliases, aliasTitle)
	}
	for _, alias := range mangaAliases {
		if alias != "暂无" {
			aliases = append(aliases, alias)
		}
	}
	status := detailUl.Find("li:nth-child(4) span:nth-child(2)").Text()
	newestChapter := detailUl.Find("li:nth-child(4) a").Text()
	newestDate := detailUl.Find("li:nth-child(4) span:nth-child(3)").Text()
	briefIntroduction := doc.Find("div#intro-cut").Text()
	introduction := doc.Find("div#intro-all").Text()
	mangaRank := doc.Find("div.rank").Text()

	genreA := detailUl.Find("li:nth-child(2) span:nth-child(1) a")
	genres := make([]*object.Category, 0)
	genreA.Each(func(idx int, sel *goquery.Selection) {
		genres = append(genres, m.categoryService.getCategoryFromA(sel))
	})
	authorA := detailUl.Find("li:nth-child(2) span:nth-child(2) a")
	authors := make([]*object.TinyAuthor, 0)
	authorA.Each(func(idx int, sel *goquery.Selection) {
		authors = append(authors, m.authorService.getAuthorFromA(sel))
	})
	obj := &object.Manga{
		Mid:               mid,
		Title:             title,
		Cover:             static.ParseCoverUrl(cover),
		Url:               url,
		PublishYear:       publishYear,
		MangaZone:         mangaZone,
		AlphabetIndex:     alphabetIndex,
		Genres:            genres,
		Authors:           authors,
		Alias:             strings.Join(mangaAliases, ", "),
		AliasTitle:        aliasTitle,
		Aliases:           aliases,
		Finished:          status == "已完结",
		NewestChapter:     newestChapter,
		NewestDate:        newestDate,
		BriefIntroduction: strings.TrimSpace(briefIntroduction),
		Introduction:      strings.TrimSpace(introduction),
		MangaRank:         mangaRank,
		Copyright:         true,
	}

	// get score
	scoreUrl := fmt.Sprintf(static.MANGA_SCORE_URL, mid)
	scoreJsonBs, _, err := m.httpService.HttpGet(scoreUrl, nil)
	if err != nil {
		return nil, err
	}
	scoreMap := make(map[string]any)
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
	} else if scoreJson, ok := scoreJsonItf.(map[string]any); !ok {
		return nil, scoreErr
	} else {
		s1 := scoreJson["s1"].(float64)
		s2 := scoreJson["s2"].(float64)
		s3 := scoreJson["s3"].(float64)
		s4 := scoreJson["s4"].(float64)
		s5 := scoreJson["s5"].(float64)
		tot := s1 + s2 + s3 + s4 + s5
		if tot > 0 {
			avg := (s1*1 + s2*2 + s3*3 + s4*4 + s5*5) / (tot * 5)
			per1 := fmt.Sprintf("%.01f%%", s1/tot*100) // 0.0%
			per2 := fmt.Sprintf("%.01f%%", s2/tot*100)
			per3 := fmt.Sprintf("%.01f%%", s3/tot*100)
			per4 := fmt.Sprintf("%.01f%%", s4/tot*100)
			per5 := fmt.Sprintf("%.01f%%", s5/tot*100)
			obj.ScoreCount = int32(tot)
			obj.AverageScore = float32(math.Round(avg*100) / 10) // 0.0
			obj.PerScores = [6]string{"", per1, per2, per3, per4, per5}
		} else {
			obj.ScoreCount = 0
			obj.AverageScore = 0.0
			obj.PerScores = [6]string{"", "0.0%", "0.0%", "0.0%", "0.0%", "0.0%"}
		}
	}

	obj.Copyright = !bytes.Contains(bs, []byte("版权方的要求"))
	obj.Banned = doc.Find("a#checkAdult").Length() != 0

	// get chapter groups
	newDoc := doc
	if vs := doc.Find("input#__VIEWSTATE"); vs.Length() != 0 {
		value := vs.AttrOr("value", "")
		hiddenHtml, err := lzstring.DecompressLZStringFromBase64(value)
		if err == nil {
			hiddenHtml = `<div class="chapter cf mt16">` + hiddenHtml + `</div>`
			newDoc, err = goquery.NewDocumentFromReader(strings.NewReader(hiddenHtml))
			if err != nil {
				return nil, err
			}
		}
	}

	groupTitleH4s := newDoc.Find("div.chapter h4").Children()
	groupListDivs := newDoc.Find("div.chapter div.chapter-list")
	groupTitles := make([]string, groupTitleH4s.Length())
	groups := make([]*object.MangaChapterGroup, len(groupTitles))
	groupTitleH4s.Each(func(idx int, sel *goquery.Selection) {
		groupTitles[idx] = sel.Text()
	})
	groupListDivs.Each(func(idx int, sel *goquery.Selection) {
		chapters := make([]*object.TinyMangaChapter, 0)
		sel.Find("ul").Each(func(idx int, sel *goquery.Selection) {
			chaptersInUl := make([]*object.TinyMangaChapter, 0)
			sel.Find("li").Each(func(idx int, sel *goquery.Selection) {
				title := sel.Find("a").AttrOr("title", "")
				pageCount, _ := xnumber.Atoi32(strings.TrimSuffix(sel.Find("i").Text(), "p"))
				url := sel.Find("a").AttrOr("href", "")
				sp := strings.Split(url, "/")
				cid, _ := xnumber.Atou64(strings.TrimSuffix(sp[len(sp)-1], ".html"))
				em := sel.Find("em").AttrOr("class", "")
				chaptersInUl = append(chaptersInUl, &object.TinyMangaChapter{
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
		for i := 0; i < len(chapters); i++ {
			chapters[i].Group = groupTitles[idx]
			chapters[i].Number = int32(len(chapters) - i)
		}
		groups[idx] = &object.MangaChapterGroup{
			Title:    groupTitles[idx],
			Chapters: chapters,
		}
	})
	obj.ChapterGroups = groups

	// return
	return obj, nil
}

func (m *MangaService) GetRandomMangaInfo() (*object.RandomMangaInfo, error) {
	resp, err := m.httpService.HttpHeadNoRedirect(static.MANGA_RANDOM_URL, nil) // 302 Found
	if err != nil {
		return nil, err
	}

	location := resp.Header.Get("Location")
	sp := strings.Split(strings.Trim(location, "/"), "/") // /comic/25882/
	if len(sp) == 0 {
		return nil, errors.New("failed to get random manga")
	}
	mid, err := xnumber.Atou64(sp[len(sp)-1])
	if err != nil {
		return nil, errors.New("failed to get random manga")
	}

	info := &object.RandomMangaInfo{
		Mid: mid,
		Url: fmt.Sprintf(static.MANGA_PAGE_URL, mid),
	}
	return info, nil
}

func (m *MangaService) VoteManga(mid uint64, score uint8) error {
	url := fmt.Sprintf(static.MANGA_VOTE_URL, mid, score) // score: 1-5
	bs, _, err := m.httpService.HttpGet(url, func(r *http.Request) {
		r.Header.Set(headers.Referer, fmt.Sprintf(static.MANGA_PAGE_URL, mid))
	})
	if err != nil {
		return err
	}

	type model struct {
		Success bool `json:"success"`
	}
	mm := &model{}
	err = json.Unmarshal(bs, mm)
	if err != nil {
		return err
	}

	// { "success": true }
	if !mm.Success {
		return fmt.Errorf("can not vote manga %d", mid)
	}
	return nil
}

func (m *MangaService) GetMangaChapter(mid, cid uint64) (*object.MangaChapter, error) {
	// get document
	url := fmt.Sprintf(static.MANGA_CHAPTER_URL, mid, cid)
	bs, doc, err := m.httpService.HttpGetDocument(url, nil)
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
	obj := &object.MangaChapter{Copyright: true}
	err = json.Unmarshal([]byte(decodeJson), &obj)
	if err != nil {
		if bytes.Contains(bs, []byte("版权方的要求")) {
			obj.Copyright = false
			return obj, nil
		}
		return nil, fmt.Errorf("chapter script error: %v", err)
	}
	obj.MangaCover = fmt.Sprintf(static.MANGA_COVER_S_URL, obj.MangaCover)
	obj.MangaUrl = fmt.Sprintf(static.MANGA_PAGE_URL, mid)
	obj.Url = url
	for idx := range obj.Pages {
		// https://cf.hamreus.com/scripts/core_2C5AD3BA009F5A0F5CCE4B6875F17FF70D5663A9.js
		// 自动 h: "i" (100) | "us" (1)
		// 电信 h: "eu" (100) | "i" (1) | "us" (1)
		// 联通 h: "us" (100) | "i" (1) | "eu" (1)
		obj.Pages[idx] = fmt.Sprintf("%s%s%s?e=%d&m=%s", fmt.Sprintf(static.MANGA_SOURCE_URL, "i"), obj.Path, obj.Pages[idx], obj.Sl.E, obj.Sl.M)
	}

	// return
	return obj, nil
}

func (m *MangaService) decodeChapterScript(p string, a, c int, k string) string {
	k, err := lzstring.DecompressLZStringFromBase64(k)
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
