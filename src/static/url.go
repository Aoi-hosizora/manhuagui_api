package static

import (
	"strings"

	"github.com/Aoi-hosizora/ahlib/xnumber"
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
	MANGA_USER_URL     = "https://www.manhuagui.com/user/center/index"
	MANGA_SHELF_URL    = "https://www.manhuagui.com/user/book/shelf/%d"
	MANGA_RANDOM_URL   = "https://www.manhuagui.com/tools/random.ashx"

	MANGA_SOURCE_URL       = "https://%s.hamreus.com"
	MANGA_COVER_URL        = "https://cf.hamreus.com/cpic/g/%d.jpg"
	MANGA_COVER_S_URL      = "https://cf.hamreus.com/cpic/g/%s"
	MANGA_SCORE_URL        = "https://www.manhuagui.com/tools/vote.ashx?act=get&bid=%d"
	MANGA_COMMENT_URL      = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=comment_list&book_id=%d&page_index=%d"
	MANGA_LOGIN_URL        = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_login"
	MANGA_CHECK_LOGIN_URL  = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_check_login"
	MANGA_SHELF_CHECK_URL  = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_book_shelf_check&book_id=%d"
	MANGA_SHELF_ADD_URL    = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_book_shelf_add"
	MANGA_SHELF_DELETE_URL = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_book_shelf_delete"
	MANGA_COUNT_URL        = "https://www.manhuagui.com/tools/count.ashx?bookId=%d&chapterId=%d"

	MESSAGE_ISSUE_API        = "https://api.github.com/repos/Aoi-hosizora/manhuagui_flutter/issues/21"
	MESSAGE_COMMENTS_PERPAGE = 100
	MESSAGE_COMMENTS_API     = "https://api.github.com/repos/Aoi-hosizora/manhuagui_flutter/issues/21/comments?page=%d&per_page=%d"
)

// noinspection GoSnakeCaseUsage
const (
	USER_AGENT    = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"
	GITHUB_ACCEPT = "application/vnd.github+json"

	NOT_FOUND_TOKEN     = "您正访问的页面不存在"
	NOT_FOUND2_TOKEN    = "暂时没有此类别组合的漫画"
	NOT_FOUND3_TOKEN    = "很抱歉，没有查找到与"
	LOGIN_SUCCESS_TOKEN = "会员登录成功"
	UNAUTHORIZED_TOKEN  = "输入登录密码"
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
	url = strings.ReplaceAll(url, "/b/", "/g/") // 132x176
	url = strings.ReplaceAll(url, "/h/", "/g/") // 180x240
	url = strings.ReplaceAll(url, "/l/", "/g/") // 78x104
	url = strings.ReplaceAll(url, "/m/", "/g/") // 114x152
	url = strings.ReplaceAll(url, "/s/", "/g/") // 92x122
	return url                                  // 240x360
}

func ParseMid(url string) uint64 {
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	mid, _ := xnumber.Atou64(sp[len(sp)-1])
	return mid
}

func ParseAid(url string) uint64 {
	sp := strings.Split(strings.TrimSuffix(url, "/"), "/")
	aid, _ := xnumber.Atou64(sp[len(sp)-1])
	return aid
}
