package service

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/util"
	"github.com/PuerkitoBio/goquery"
	"log"
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

func (m *MangaService) GetMangaPage(id uint64) (*vo.MangaPage, error) {
	_ = fmt.Sprintf("https://www.manhuagui.com/comic/%d", id)
	return &vo.MangaPage{}, nil
}

func (m *MangaService) GetMangaChapter(id, cid uint64) (*vo.MangaChapter, error) {
	url := fmt.Sprintf("https://www.manhuagui.com/comic/%d/%d.html", id, cid)
	return m.getMangaChapter(url), nil
}

func (m *MangaService) getMangaChapter(url string) *vo.MangaChapter {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body := resp.Body
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	var script *goquery.Selection
	doc.Find("script").Each(func(i int, sel *goquery.Selection) {
		_, ok := sel.Attr("src")
		if !ok && strings.Contains(sel.Text(), `window["\x65\x76\x61\x6c"]`) {
			script = sel
		}
	})
	if script == nil {
		log.Fatalln("could not find decode script")
	}
	decodeScript := script.Text()
	found := regexp.MustCompile(`;return p;}\('(.+\(\);)',(.+?),(.+),'(.+)'\['\\x73`).FindAllStringSubmatch(decodeScript, 1)
	if len(found) == 0 {
		log.Fatalln("could not find decode text")
	}
	p := found[0][1]                  // XX.YY({...}).ZZ();
	a, _ := strconv.Atoi(found[0][2]) // 00
	c, _ := strconv.Atoi(found[0][3]) // 00
	k := found[0][4]                  // ~['\x73\x70\x6c\x69\x63']('\x7c')

	decodeString := m.decode(p, a, c, k)
	found = regexp.MustCompile(`SMH.imgData\((.+)\).preInit`).FindAllStringSubmatch(decodeString, 1)
	if len(found) == 0 {
		log.Fatalf("could not find decode json")
	}
	decodeJson := found[0][1]
	log.Println(decodeJson)
	obj := &vo.MangaChapter{}
	// TODO https://www.manhuagui.com/comic/34707/525399.html
	err = json.Unmarshal([]byte(decodeJson), &obj)
	if err != nil {
		log.Fatalln("failed to decode json:", err)
	}
	for idx := range obj.Files {
		obj.Files[idx] = fmt.Sprintf("%s%s%s?e=%d&m=%s", "https://i.hamreus.com", obj.Path, url, obj.Sl.E, obj.Sl.M)
	}

	return obj
}

func (m *MangaService) decode(p string, a, c int, k string) string {
	k, err := util.DecompressFromEncodedUriComponent(k)
	if err != nil {
		return ""
	}
	ks := strings.Split(k, "|")

	var e func(c int) string
	e = func(c int) string {
		p1 := ""
		if c < a {
			p1 = ""
		} else {
			p1 = e(c / a)
		}
		p2 := byte(0)
		c = c % a
		if c > 35 {
			p2 = byte(c + 29)
		} else {
			p2 = "0123456789abcdefghijklmnopqrstuvwxyz"[c]
		}
		return p1 + string(p2)
	}

	d := map[string]string{}
	for ; c >= 0; c-- {
		if len(ks) <= c || ks[c] == "" {
			d[e(c)] = e(c)
		} else {
			d[e(c)] = ks[c]
		}
	}
	p = regexp.MustCompile(`\b\w\b`).ReplaceAllStringFunc(p, func(s string) string {
		return d[s]
	})
	return p
}
