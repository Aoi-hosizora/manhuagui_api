package server

import (
	"fmt"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-backend/src/controller"
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
		mangaController = controller.NewMangaController()
	)

	mangaGroup := v1.Group("manga") // /v1/manga
	{
		mangaGroup.GET(":mid", j(mangaController.GetMangaPage))
		mangaGroup.GET(":mid/:cid", j(mangaController.GetMangaChapter))
	}

	listGroup := v1.Group("list") // /v1/list
	{
		listGroup.GET("serial", j(mangaController.GetHotSerialMangas))
		listGroup.GET("finish", j(mangaController.GetFinishedMangas))
		listGroup.GET("latest", j(mangaController.GetLatestMangas))
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
