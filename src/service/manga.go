package service

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	return &vo.MangaPage{Bid: mid}, nil
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
	for idx := range obj.Files {
		// 自动 h: "i" (100) | "us" (1)
		// 电信 h: "eu" (100) | "i" (1) | "us" (1)
		// 联通 h: "us" (100) | "i" (1) | "eu" (1)
		obj.Files[idx] = fmt.Sprintf("%s%s%s?e=%d&m=%s", fmt.Sprintf(static.MANGA_SOURCE_URL, "i"), obj.Path, url, obj.Sl.E, obj.Sl.M)
	}

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
