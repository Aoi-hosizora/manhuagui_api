package server

import (
	"fmt"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-api/src/controller"
	"github.com/gin-gonic/gin"
	"strings"
)

func initRoute(engine *gin.Engine) {
	engine.HandleMethodNotAllowed = true
	engine.NoRoute(func(c *gin.Context) {
		result.Status(404).SetMessage(fmt.Sprintf("route %s is not found", c.Request.URL.Path)).JSON(c)
	})
	engine.NoMethod(func(c *gin.Context) {
		result.Status(405).SetMessage(fmt.Sprintf("method %s is not allowed", strings.ToUpper(c.Request.Method))).JSON(c)
	})
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, &gin.H{"ping": "pong"})
	})
	engine.GET("", func(c *gin.Context) {
		c.JSON(200, &gin.H{"message": "Welcome to manhuagui-backend."})
	})

	// controller
	v1 := engine.Group("v1")

	var (
		mangaController     = controller.NewMangaController()
		mangaListController = controller.NewMangaListController()
		categoryController  = controller.NewCategoryController()
		searchController    = controller.NewSearchController()
		authorController    = controller.NewAuthorController()
		rankController      = controller.NewRankController()
		commentController   = controller.NewCommentService()
		userController      = controller.NewUserController()
		shelfController     = controller.NewShelfController()
		messageController   = controller.NewMessageController()
	)

	mangaGroup := v1.Group("manga") // /v1/manga/...
	{
		mangaGroup.GET("", j(mangaController.GetAllMangas))
		mangaGroup.GET(":mid", j(mangaController.GetManga))
		mangaGroup.GET(":mid/:cid", j(mangaController.GetMangaChapter))
	}

	listGroup := v1.Group("list") // /v1/list/...
	{
		listGroup.GET("serial", j(mangaListController.GetHotSerialMangas))
		listGroup.GET("finish", j(mangaListController.GetFinishedMangas))
		listGroup.GET("latest", j(mangaListController.GetLatestMangas))
		listGroup.GET("homepage", j(mangaListController.GetHomepageMangas))
		listGroup.GET("updated", j(mangaListController.GetRecentUpdatedMangas))
	}

	categoryGroup := v1.Group("category") // /v1/category/...
	{
		categoryGroup.GET("", j(categoryController.GetCategories))
		categoryGroup.GET("genre", j(categoryController.GetGenres))
		categoryGroup.GET("zone", j(categoryController.GetZones))
		categoryGroup.GET("age", j(categoryController.GetAges))
		categoryGroup.GET("genre/:genre", j(categoryController.GetGenreMangas))
	}

	searchGroup := v1.Group("search") // /v1/search/...
	{
		searchGroup.GET(":keyword", j(searchController.SearchMangas))
	}

	authorGroup := v1.Group("author") // /v1/author/...
	{
		authorGroup.GET("", j(authorController.GetAllAuthors))
		authorGroup.GET(":aid", j(authorController.GetAuthor))
		authorGroup.GET(":aid/manga", j(authorController.GetAuthorMangas))
	}

	rankGroup := v1.Group("rank") // /v1/rank/...
	{
		rankGroup.GET("day", j(rankController.GetDayRanking))
		rankGroup.GET("week", j(rankController.GetWeekRanking))
		rankGroup.GET("month", j(rankController.GetMonthRanking))
		rankGroup.GET("total", j(rankController.GetTotalRanking))
	}

	commentGroup := v1.Group("comment") // /v1/comment/...
	{
		commentGroup.GET("manga/:mid", j(commentController.GetComments))
	}

	userGroup := v1.Group("user") // /v1/user/...
	{
		userGroup.POST("login", j(userController.Login))
		userGroup.POST("check_login", j(userController.CheckLogin))
		userGroup.GET("info", j(userController.GetUser))
		userGroup.GET("manga/:mid/:cid", j(userController.RecordManga)) // deprecated
		userGroup.POST("manga/:mid/:cid", j(userController.RecordManga))
	}

	shelfGroup := v1.Group("shelf") // /v1/shelf/...
	{
		shelfGroup.GET("", j(shelfController.GetShelfMangas))
		shelfGroup.GET(":mid", j(shelfController.CheckMangaInShelf))
		shelfGroup.POST(":mid", j(shelfController.SaveMangaToShelf))
		shelfGroup.DELETE(":mid", j(shelfController.RemoveMangaFromShelf))
	}

	messageGroup := v1.Group("message") // /v1/message/...
	{
		messageGroup.GET("", j(messageController.GetMessages))
		messageGroup.GET("latest", j(messageController.GetLatestMessage))
	}
}

// j Simplify controller's functions.
func j(fn func(c *gin.Context) *result.Result) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.IsAborted() {
			return
		}
		r := fn(c)
		if r != nil {
			r.JSON(c)
		}
	}
}
