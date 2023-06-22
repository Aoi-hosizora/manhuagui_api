package static

const (
	HOMEPAGE_URL       = "https://www.manhuagui.com"
	MANGA_PAGE_URL     = "https://www.manhuagui.com/comic/%d"
	MANGA_CHAPTER_URL  = "https://www.manhuagui.com/comic/%d/%d.html"
	MANGA_UPDATE_URL   = "https://www.manhuagui.com/update/d30.html"
	MANGA_UPDATE_MURL  = "https://m.manhuagui.com/update/?page=%d&ajax=1&order="
	MANGA_OVERALL_MURL = "https://m.manhuagui.com/list/?page=%d&catid=0&ajax=1&order="
	MANGA_CATEGORY_URL = "https://www.manhuagui.com/list"
	MANGA_SEARCH_URL   = "https://www.manhuagui.com/s/%s_p%d.html"
	MANGA_AUTHORS_URL  = "https://www.manhuagui.com/alist"
	MANGA_AUTHOR_URL   = "https://www.manhuagui.com/author/%d/%s_p%d.html"
	MANGA_RANK_URL     = "https://www.manhuagui.com/rank/%s"
	MANGA_USER_URL     = "https://www.manhuagui.com/user/center/index"
	MANGA_SHELF_URL    = "https://www.manhuagui.com/user/book/shelf/%d"
	MANGA_RANDOM_URL   = "https://www.manhuagui.com/tools/random.ashx"
	MANGA_VOTE_URL     = "https://www.manhuagui.com/tools/vote.ashx?act=vote&bid=%d&s=%d"

	MANGA_SOURCE_URL       = "https://%s.hamreus.com"
	MANGA_COVER_URL        = "https://cf.hamreus.com/cpic/g/%d.jpg"
	MANGA_COVER_S_URL      = "https://cf.hamreus.com/cpic/g/%s"
	MANGA_SCORE_URL        = "https://www.manhuagui.com/tools/vote.ashx?act=get&bid=%d"
	MANGA_COMMENT_URL      = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=comment_list&book_id=%d&page_index=%d"
	MANGA_LIKE_COMMENT_URL = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=comment_support&comment_id=%d"
	MANGA_ADD_COMMENT_URL  = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=comment_add"
	MANGA_LOGIN_URL        = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_login"
	MANGA_CHECK_LOGIN_URL  = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_check_login"
	MANGA_SHELF_CHECK_URL  = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_book_shelf_check&book_id=%d"
	MANGA_SHELF_ADD_URL    = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_book_shelf_add"
	MANGA_SHELF_DELETE_URL = "https://www.manhuagui.com/tools/submit_ajax.ashx?action=user_book_shelf_delete"
	MANGA_COUNT_URL        = "https://www.manhuagui.com/tools/count.ashx?bookId=%d&chapterId=%d"

	DEFAULT_USER_AVATAR_URL  = "https://cf.hamreus.com/images/default.png"
	DEFAULT_AUTHOR_COVER_URL = "https://cf.hamreus.com/zpic/none.jpg"

	MESSAGE_ISSUE_API        = "https://api.github.com/repos/Aoi-hosizora/manhuagui_flutter/issues/21"
	MESSAGE_COMMENTS_PERPAGE = 100
	MESSAGE_COMMENTS_API     = "https://api.github.com/repos/Aoi-hosizora/manhuagui_flutter/issues/21/comments?page=%d&per_page=%d"
)

const (
	USER_AGENT    = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"
	USER_AGENT_M  = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36"
	REFERER       = "https://www.manhuagui.com/"
	GITHUB_ACCEPT = "application/vnd.github+json"

	NOT_FOUND_TOKEN     = "您正访问的页面不存在"
	NOT_FOUND2_TOKEN    = "暂时没有此类别组合的漫画"
	NOT_FOUND3_TOKEN    = "很抱歉，没有查找到与"
	LOGIN_SUCCESS_TOKEN = "会员登录成功"
	UNAUTHORIZED_TOKEN  = "输入登录密码"
)
