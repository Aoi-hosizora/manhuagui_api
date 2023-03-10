package sn

import "github.com/Aoi-hosizora/ahlib/xmodule"

const (
	SConfig xmodule.ModuleName = "config" // *config.Config
	SLogger xmodule.ModuleName = "logger" // *logrus.Logger

	SHttpService      xmodule.ModuleName = "http-service"       // *service.HttpService
	SMangaService     xmodule.ModuleName = "manga-service"      // *service.MangaService
	SMangaListService xmodule.ModuleName = "manga-list-service" // *service.MangaListService
	SCategoryService  xmodule.ModuleName = "category-service"   // *service.CategoryService
	SSearchService    xmodule.ModuleName = "search-service"     // *service.SearchService
	SAuthorService    xmodule.ModuleName = "author-service"     // *service.AuthorService
	SRankService      xmodule.ModuleName = "rank-service"       // *service.RankService
	SCommentService   xmodule.ModuleName = "comment-service"    // *service.CommentService
	SUserService      xmodule.ModuleName = "user-service"       // *service.UserService
	SShelfService     xmodule.ModuleName = "shelf-service"      // *service.ShelfService
	SMessageService   xmodule.ModuleName = "message-service"    // *service.MessageService
)
