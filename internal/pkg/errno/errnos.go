package errno

var (
	errno4xx = int32(40000) - 1
	errno5xx = int32(50000) - 1
)

func new4(s int32, m string) *Error { errno4xx++; return New(s, errno4xx, m) }
func new5(s int32, m string) *Error { errno5xx++; return New(s, errno5xx, m) }

var (
	RequestParamError  = new4(400, "request parameter error")
	ServerUnknownError = new5(500, "server unknown error")
)

// manga
var (
	GetAllMangasError         = new5(500, "无法获取漫画列表")
	GetMangaError             = new5(500, "无法获取漫画信息")
	GetRandomMangaError       = new5(500, "无法获取随机漫画")
	VoteMangaError            = new5(500, "无法对漫画投票")
	MangaNotFoundError        = new4(404, "漫画未找到")
	GetMangaChapterError      = new5(500, "无法获取漫画章节")
	MangaChapterNotFoundError = new4(404, "漫画章节未找到")
	GetHotSerialMangasError   = new5(500, "无法获取热门连载漫画")
	GetFinishedMangasError    = new5(500, "无法获取经典完结漫画")
	GetLatestMangasError      = new5(500, "无法获取最新上架漫画")
	GetHomepageMangasError    = new5(500, "无法获取主页漫画列表")
	GetUpdatedMangasError     = new5(500, "无法获取最近更新漫画")
)

// category
var (
	GetGenresError      = new5(500, "无法获取类别信息")
	GetZonesError       = new5(500, "无法获取地区信息")
	GetAgesError        = new5(500, "无法获取受众信息")
	GetGenreMangasError = new5(500, "无法获取漫画分类结果")
	GenreNotFoundError  = new4(404, "类别未找到")
	GetCategoriesError  = new5(500, "无法获取漫画分类信息")
)

// search
var (
	SearchMangasError = new5(500, "无法获取漫画搜索结果")
)

// author
var (
	GetAllAuthorsError   = new5(500, "无法获取漫画家列表")
	GetAuthorError       = new5(500, "无法获取漫画家信息")
	AuthorNotFound       = new4(404, "漫画家未找到")
	GetAuthorMangasError = new5(500, "无法获取漫画家的漫画列表")
)

// rank
var (
	GetRankingError          = new5(500, "无法获取排行榜")
	RankingTypeNotFoundError = new4(404, "排行榜类别未找到")
)

// comment
var (
	GetMangaCommentsError = new5(500, "无法获取漫画评论")
	LikeCommentError      = new5(500, "无法点赞评论")
	EmptyCommentError     = new4(400, "评论内容为空")
	AddCommentError       = new5(500, "无法评论漫画")
	ReplyCommentError     = new5(500, "无法回复评论")
)

// user
var (
	LoginError                = new5(500, "无法登录")
	PasswordError             = new4(401, "用户名或密码错误")
	CheckLoginError           = new5(500, "无法检查登录状态")
	UnauthorizedError         = new4(401, "用户未登录")
	GetUserError              = new5(500, "无法获取用户信息")
	GetShelfMangasError       = new5(500, "无法获取书架列表")
	CheckMangaShelfError      = new5(500, "无法检查书架漫画状态")
	SaveMangaToShelfError     = new5(500, "无法添加订阅")
	RemoveMangaFromShelfError = new5(500, "无法删除订阅")
	MangaAlreadyInShelfError  = new4(409, "漫画已经被订阅")
	MangaNotInShelfYetError   = new4(404, "漫画还没有被订阅")
	CountMangaError           = new5(500, "无法记录漫画阅读历史")
)

// message
var (
	GetMessageError = new5(500, "无法获取消息")
)
