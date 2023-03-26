package server

import (
	"fmt"
	"github.com/Aoi-hosizora/manhuagui-api/internal/controller"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/gin-gonic/gin"
	"strings"
)

func setupRoutes(engine *gin.Engine) {
	// ===========
	// meta routes
	// ===========

	engine.NoRoute(func(c *gin.Context) {
		msg := fmt.Sprintf("route '%s' is not found", c.Request.URL.Path)
		result.Status(404).SetMessage(msg).JSON(c)
	})
	engine.NoMethod(func(c *gin.Context) {
		msg := fmt.Sprintf("method '%s' is not allowed", strings.ToUpper(c.Request.Method))
		result.Status(405).SetMessage(msg).JSON(c)
	})
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})
	engine.GET("", func(c *gin.Context) {
		c.JSON(200, &gin.H{"message": "Here is manhuagui-backend."})
	})

	// ===========
	// controllers
	// ===========

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

	// ============
	// route groups
	// ============

	// v1 route group
	v1Group := engine.Group("v1")
	v1 := result.NewRouteRegisterer(v1Group)

	mangaGroup := v1.Group("manga")
	mangaGroup.GET("", mangaController.GetAllMangas)
	mangaGroup.GET(":mid", mangaController.GetManga)
	mangaGroup.GET("random", mangaController.GetRandomManga)
	mangaGroup.GET(":mid/:cid", mangaController.GetMangaChapter)

	listGroup := v1.Group("list")
	listGroup.GET("serial", mangaListController.GetHotSerialMangas)
	listGroup.GET("finish", mangaListController.GetFinishedMangas)
	listGroup.GET("latest", mangaListController.GetLatestMangas)
	listGroup.GET("homepage", mangaListController.GetHomepageMangas)
	listGroup.GET("updated", mangaListController.GetRecentUpdatedMangas)

	categoryGroup := v1.Group("category")
	categoryGroup.GET("", categoryController.GetCategories)
	categoryGroup.GET("genre", categoryController.GetGenres)
	categoryGroup.GET("zone", categoryController.GetZones)
	categoryGroup.GET("age", categoryController.GetAges)
	categoryGroup.GET("genre/:genre", categoryController.GetGenreMangas)

	searchGroup := v1.Group("search")
	searchGroup.GET("", searchController.SearchMangas)
	searchGroup.GET(":keyword", searchController.SearchMangas) // deprecated

	authorGroup := v1.Group("author")
	authorGroup.GET("", authorController.GetAllAuthors)
	authorGroup.GET(":aid", authorController.GetAuthor)
	authorGroup.GET(":aid/manga", authorController.GetAuthorMangas)

	rankGroup := v1.Group("rank")
	rankGroup.GET("day", rankController.GetDayRanking)
	rankGroup.GET("week", rankController.GetWeekRanking)
	rankGroup.GET("month", rankController.GetMonthRanking)
	rankGroup.GET("total", rankController.GetTotalRanking)

	commentGroup := v1.Group("comment")
	commentGroup.GET("manga/:mid", commentController.GetComments)

	userGroup := v1.Group("user")
	userGroup.POST("login", userController.Login)
	userGroup.POST("check_login", userController.CheckLogin)
	userGroup.GET("info", userController.GetUser)
	userGroup.GET("manga/:mid/:cid", userController.RecordManga) // deprecated
	userGroup.POST("manga/:mid/:cid", userController.RecordManga)

	shelfGroup := v1.Group("shelf")
	shelfGroup.GET("", shelfController.GetShelfMangas)
	shelfGroup.GET(":mid", shelfController.CheckMangaInShelf)
	shelfGroup.POST(":mid", shelfController.AddMangaToShelf)
	shelfGroup.DELETE(":mid", shelfController.RemoveMangaFromShelf)

	messageGroup := v1.Group("message")
	messageGroup.GET("", messageController.GetMessages)
	messageGroup.GET("latest", messageController.GetLatestMessage)
}
