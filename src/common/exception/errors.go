package exception

var (
	cerr = int32(40000) // client error code
	serr = int32(50000) // server error code
)

// Return ++ cerr.
func ce() int32 { cerr++; return cerr }

// Return ++ serr.
func se() int32 { serr++; return serr }

var (
	RequestParamError   = New(400, cerr, "request param error")
	ServerRecoveryError = New(500, serr, "server unknown error")
)

// manga
var (
	GetAllMangasError         = New(500, se(), "无法获取漫画列表")
	GetMangaError             = New(500, se(), "无法获取漫画信息")
	MangaNotFoundError        = New(404, ce(), "漫画未找到")
	GetMangaChapterError      = New(500, se(), "无法获取漫画章节")
	MangaChapterNotFoundError = New(404, ce(), "漫画章节未找到")
	GetHotSerialMangasError   = New(500, se(), "无法获取热门连载漫画")
	GetFinishedMangasError    = New(500, se(), "无法获取经典完结漫画")
	GetLatestMangasError      = New(500, se(), "无法获取最新上架漫画")
	GetHomepageMangasError    = New(500, se(), "无法获取主页漫画列表")
	GetUpdatedMangasError     = New(500, se(), "无法获取最近更新漫画")
)

// category
var (
	GetGenresError      = New(500, se(), "无法获取类别信息")
	GetZonesError       = New(500, se(), "无法获取地区信息")
	GetAgesError        = New(500, se(), "无法获取受众信息")
	GetGenreMangasError = New(500, se(), "无法获取漫画分类结果")
	GenreNotFoundError  = New(404, ce(), "类别未找到")
)

// search
var (
	SearchMangasError = New(500, se(), "无法获取漫画搜索结果")
)

// author
var (
	GetAllAuthorsError   = New(500, se(), "无法获取漫画家列表")
	GetAuthorError       = New(500, se(), "无法获取漫画家信息")
	AuthorNotFound       = New(404, ce(), "漫画家未找到")
	GetAuthorMangasError = New(500, se(), "无法获取漫画家的漫画列表")
)

// rank
var (
	GetRankingError          = New(500, se(), "无法获取排行榜")
	RankingTypeNotFoundError = New(404, ce(), "排行榜类别未找到")
)

// comment
var (
	GetMangaCommentsError = New(500, se(), "无法获取漫画评论")
)

// user
var (
	LoginError                = New(500, se(), "无法登录")
	PasswordError             = New(401, ce(), "用户名或密码错误")
	CheckLoginError           = New(500, se(), "无法检查登录状态")
	UnauthorizedError         = New(401, ce(), "用户未登录")
	GetUserError              = New(500, se(), "无法获取用户信息")
	GetShelfMangasError       = New(500, se(), "无法获取书架列表")
	CheckMangaShelfError      = New(500, se(), "无法检查书架漫画状态")
	SaveMangaToShelfError     = New(500, se(), "无法添加订阅")
	RemoveMangaFromShelfError = New(500, se(), "无法删除订阅")
	MangaAlreadyInShelfError  = New(409, ce(), "漫画已经被订阅")
	MangaNotInShelfYetError   = New(404, ce(), "漫画还没有被订阅")
	CountMangaError           = New(500, se(), "无法记录漫画阅读历史")
)
