package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	req, err := http.NewRequest("GET", "https://www.manhuagui.com/comic/34707/472931.html", nil)
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

	log.Println(p)
	log.Println(a)
	log.Println(c)
	log.Println(k)

	log.Println(decode(p, a, c, k))
}

func decode(p string, a, c int, k string) string {
	k, err := DecompressFromEncodedUriComponent(k)
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
	for c--; c > 0; c-- {
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

// https://github.com/Aoi-hosizora/recent_works/issues/57

/*
// https://www.manhuagui.com/comic/34707/472931.html

<script type="text/javascript" src="https://cf.hamreus.com/scripts/core_5649DEA1194321210A9977D8C49E535033C95CAA.js"></script>
<script type="text/javascript">
	window["\x65\x76\x61\x6c"](
		function (p, a, c, k, e, d) {
			e = function (c) {
				return (c < a ? "" : e(parseInt(c / a))) + ((c = c % a) > 35 ? String.fromCharCode(c + 29) : c.toString(36));
			};
			if (!''.replace(/^/, String)) {
				while (c--) {
					d[e(c)] = k[c] || e(c);
				}
				k = [function (e) { return d[e] }];
				e = function () { return '\\w+' };
				c = 1;
			};
			while (c--) {
				if (k[c]) {
					p = p.replace(new RegExp('\\b' + e(c) + '\\b', 'g'), k[c]);
					// =>
					p = p.replace(/\b\w\b/g, function (e) { return d[e] });
				}
			}
			return p;
		}(
			'E.h({"s":f,"r":"q","p":"f.a","o":n,"l":"c.d","k":["0.a.b","1.a.b","2.a.b","3.a.b","4.a.b","5.a.b","6.a.b","7.a.b","8.a.b","9.a.b","j.a.b"],"i":v,"t":u,"D":"/w/0-9/I/H/c.d/","G":1,"F":"","C":B,"A":0,"z":{"e":y,"m":"x"}}).g();',
			45,
			45,
			'D41hWAODmwO4FMBGliBpvAjMDhd6JAZgBYB2ABmOEgCcEBJAOwEsAXYRgW2gBEBDZn4ADNGTAM4ALBABNspIYwA2CUcADG9HuwQgSAJgCc+LKsYyUjVcEAAcoDh0q4Eag4Eg1anp4EvrYsgngtHakKL4wABCcACiPADiAIoAWgByjAD60ACCAML0ABwAEnkACroA8gCOsdgAbKQArIa1ORi6wKIKlDQAbrQyevrELfQIAB7MPZR84sAAygCyeU4KAPaqANYpqpai/MwArioYohhT+voYGEA='
				['\x73\x70\x6c\x69\x63']('\x7c'),
			0,
			{},
		);
	);
</script>
*/

/*

'D41hWAODmwO4FMBGliBpvAjMDhd6JAZgBYB2ABmOEgCcEBJAOwEsAXYRgW2gBEBDZn4ADNGTAM4ALBABNspIYwA2CUcADG9HuwQgSAJgCc+LKsYyUjVcEAAcoDh0q4Eag4Eg1anp4EvrYsgngtHakKL4wABCcACiPADiAIoAWgByjAD60ACCAML0ABwAEnkACroA8gCOsdgAbKQArIa1ORi6wKIKlDQAbrQyevrELfQIAB7MPZR84sAAygCyeU4KAPaqANYpqpai/MwArioYohhT+voYGEA='
.splic('|')

LZString.decompressFromBase64(
	'D41hWAODmwO4FMBGliBpvAjMDhd6JAZgBYB2ABmOEgCcEBJAOwEsAXYRgW2gBEBDZn4ADNGTAM4ALBABNspIYwA2CUcADG9HuwQgSAJgCc+LKsYyUjVcEAAcoDh0q4Eag4Eg1anp4EvrYsgngtHakKL4wABCcACiPADiAIoAWgByjAD60ACCAML0ABwAEnkACroA8gCOsdgAbKQArIa1ORi6wKIKlDQAbrQyevrELfQIAB7MPZR84sAAygCyeU4KAPaqANYpqpai/MwArioYohhT+voYGEA='
).split('|')
'||||||||||jpg|webp|第1|1话||34707|preInit|imgData|finished|10|files|cname||472931|cid|bpic|一霎一花|bname|bid|len|11|false|ps3|BwEaGQZNi_gACn8HHP2OqQ|1605935812|sl|prevId|472972|nextId|path|SMH|block_cc|status|1s1h|9911'

*/