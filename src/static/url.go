package static

import (
	"strings"
)

// noinspection GoSnakeCaseUsage
const (
	HOMEPAGE_URL       = "https://www.manhuagui.com"
	MANGA_PAGE_URL     = "https://www.manhuagui.com/comic/%d"
	MANGA_CHAPTER_URL  = "https://www.manhuagui.com/comic/%d/%d.html"
	MANGA_UPDATE_URL   = "https://www.manhuagui.com/update/d30.html"
	MANGA_CATEGORY_URL = "https://www.manhuagui.com/list"
	MANGA_SEARCH_URL   = "https://www.manhuagui.com/s/%s_p%d.html"
	MANGA_AUTHORS_URL  = "https://www.manhuagui.com/alist"
	MANGA_AUTHOR_URL   = "https://www.manhuagui.com/author/%d/%s_p%d.html"
	MANGA_RANK_URL     = "https://www.manhuagui.com/rank/%s"

	MANGA_SCORE_URL  = "https://www.manhuagui.com/tools/vote.ashx?act=get&bid=%d"
	MANGA_SOURCE_URL = "https://%s.hamreus.com"
)

// noinspection GoSnakeCaseUsage
const (
	USER_AGENT       = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"
	NOT_FOUND_TOKEN  = "您正访问的页面不存在"
	NOT_FOUND2_TOKEN = "暂时没有此类别组合的漫画"
	NOT_FOUND3_TOKEN = "很抱歉，没有查找到与"
)

/*
Demo url:

https://www.manhuagui.com/comic/34707
https://www.manhuagui.com/comic/34707/472931.html
https://www.manhuagui.com/author/5802/
*/

func ParseCoverUrl(url string) string {
	url = strings.ReplaceAll(url, "/b/", "/g/") // 132x176
	url = strings.ReplaceAll(url, "/h/", "/g/") // 180x240
	url = strings.ReplaceAll(url, "/l/", "/g/") // 78x104
	url = strings.ReplaceAll(url, "/m/", "/g/") // 114x152
	url = strings.ReplaceAll(url, "/s/", "/g/") // 92x122
	return url                                  // 240x360
}
