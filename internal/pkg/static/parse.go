package static

import (
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"strings"
)

/*
	Demo url:

	https://www.manhuagui.com/comic/34707
	https://www.manhuagui.com/comic/34707/472931.html
	https://www.manhuagui.com/author/5802/
*/

func ParseCoverUrl(url string) string {
	url = strings.TrimPrefix(url, "//")
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	url = strings.ReplaceAll(url, "/b/", "/g/") // b: 132x176
	url = strings.ReplaceAll(url, "/h/", "/g/") // h: 180x240
	url = strings.ReplaceAll(url, "/l/", "/g/") // l: 78x104
	url = strings.ReplaceAll(url, "/m/", "/g/") // m: 114x152
	url = strings.ReplaceAll(url, "/s/", "/g/") // s: 92x122
	return url                                  // g: 240x360
}

func ParseMid(url string) uint64 {
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	if len(sp) > 0 {
		mid, _ := xnumber.Atou64(sp[len(sp)-1])
		return mid
	}
	return 0
}

func ParseAid(url string) uint64 {
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	if len(sp) > 0 {
		aid, _ := xnumber.Atou64(sp[len(sp)-1])
		return aid
	}
	return 0
}
