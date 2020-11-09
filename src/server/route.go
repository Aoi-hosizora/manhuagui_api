package server

import (
	"fmt"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/result"
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
	_ = engine.Group("v1")
}
