package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/static"
	"github.com/Aoi-hosizora/manhuagui-backend/src/util"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type MangaService struct {
}

func NewMangaService() *MangaService {
	return &MangaService{}
}

func (m *MangaService) GetMangaPage(mid uint64) (*vo.MangaPage, error) {
	url := fmt.Sprintf(static.MANGA_PAGE_URL, mid)

	// get html
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", static.USER_AGENT)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("response error: %v", err)
	}
	if bytes.Contains(bs, []byte(static.NOT_FOUND_TOKEN)) {
		return nil, nil
	}

	// get details
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
	if err != nil {
		return nil, fmt.Errorf("document error: %v", err)
	}

	bname := doc.Find("div.book-title").Text()
	bpic := doc.Find("p.hcover img").AttrOr("src", "")
	detailUl := doc.Find("ul.detail-list")
	publishYear := detailUl.Find("li:nth-child(1) span:nth-child(1) a").Text()
	zone := detailUl.Find("li:nth-child(1) span:nth-child(2) a").Text()
	alphabetIndex := detailUl.Find("li:nth-child(1) span:nth-child(3) a").Text()
	mangaType := detailUl.Find("li:nth-child(2) span:nth-child(1) a").Text()
	authorName := detailUl.Find("li:nth-child(2) span:nth-child(2) a").Text()
	alias := detailUl.Find("li:nth-child(3) span:nth-child(1)").Text()
	status := detailUl.Find("li:nth-child(4) span:nth-child(2)").Text()
	newestChapter := detailUl.Find("li:nth-child(4) a").Text()
	newestDate := detailUl.Find("li:nth-child(4) span:nth-child(3)").Text()
	introduction := doc.Find("div#intro-all").Text()
	rank := doc.Find("div.rank").Text()
	obj := &vo.MangaPage{
		Bid:           mid,
		Bname:         bname,
		Bpic:          bpic,
		Url:           url,
		PublishYear:   publishYear,
		Zone:          zone,
		AlphabetIndex: alphabetIndex,
		Type:          mangaType,
		AuthorName:    authorName,
		Alias:         strings.TrimPrefix(alias, "漫画别名："),
		Finished:      status == "已完结",
		NewestChapter: newestChapter,
		NewestDate:    newestDate,
		Introduction:  strings.TrimSpace(introduction),
		Rank:          rank,
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
		links := make([]*vo.MangaChapterLink, 0)
		sel.Find("ul").Each(func(idx int, sel *goquery.Selection) {
			linksInUl := make([]*vo.MangaChapterLink, 0)
			sel.Find("li").Each(func(idx int, sel *goquery.Selection) {
				cname := sel.Find("a").AttrOr("title", "")
				pages, _ := xnumber.Atoi32(strings.TrimSuffix(sel.Find("i").Text(), "p"))
				url := sel.Find("a").AttrOr("href", "")
				if url != "" {
					url = static.HOMEPAGE_URL + url
				}
				sp := strings.Split(url, "/")
				cid, _ := xnumber.Atou64(strings.TrimSuffix(sp[len(sp)-1], ".html"))
				em := sel.Find("em").AttrOr("class", "")
				linksInUl = append(linksInUl, &vo.MangaChapterLink{
					Cid:   cid,
					Cname: cname,
					Url:   url,
					Pages: pages,
					New:   em != "",
				})
			})
			links = append(linksInUl, links...)
		})
		groups[idx] = &vo.MangaChapterGroup{
			Title: groupTitles[idx],
			Links: links,
		}
	})
	obj.Chapters = groups

	// return
	return obj, nil
}

func (m *MangaService) GetMangaChapter(mid, cid uint64) (*vo.MangaChapter, error) {
	url := fmt.Sprintf(static.MANGA_CHAPTER_URL, mid, cid)

	// get html
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", static.USER_AGENT)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	body := resp.Body
	defer body.Close()
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("response error: %v", err)
	}
	if bytes.Contains(bs, []byte(static.NOT_FOUND_TOKEN)) {
		return nil, nil
	}

	// get script
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
	if err != nil {
		return nil, fmt.Errorf("document error: %v", err)
	}
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
	for idx := range obj.Files {
		// 自动 h: "i" (100) | "us" (1)
		// 电信 h: "eu" (100) | "i" (1) | "us" (1)
		// 联通 h: "us" (100) | "i" (1) | "eu" (1)
		obj.Files[idx] = fmt.Sprintf("%s%s%s?e=%d&m=%s", fmt.Sprintf(static.MANGA_SOURCE_URL, "i"), obj.Path, obj.Files[idx], obj.Sl.E, obj.Sl.M)
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
