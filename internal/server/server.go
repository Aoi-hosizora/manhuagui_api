package server

import (
	"context"
	"fmt"
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/api"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/server/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	goapidoc.SetDocument(
		"localhost:10018", "/",
		goapidoc.NewInfo("manhuagui-backend", "An unofficial backend for manhuagui written in golang/gin", "1.0").
			Contact(goapidoc.NewContact("Aoi-hosizora", "https://github.com/Aoi-hosizora", "aoihosizora@hotmail.com")),
	)

	goapidoc.SetOption(goapidoc.NewOption().
		AddTags(
			goapidoc.NewTag("Manga", "manga-controller"),
			goapidoc.NewTag("MangaList", "manga-list-controller"),
			goapidoc.NewTag("Category", "category-controller"),
			goapidoc.NewTag("Search", "search-controller"),
			goapidoc.NewTag("Author", "author-controller"),
			goapidoc.NewTag("Rank", "rank-controller"),
			goapidoc.NewTag("Comment", "comment-controller"),
			goapidoc.NewTag("User", "user-controller"),
			goapidoc.NewTag("Shelf", "shelf-controller"),
			goapidoc.NewTag("Message", "message-controller"),
		),
	)
}

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func NewServer() *Server {
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config)
	gin.SetMode(cfg.Meta.RunMode)
	engine := gin.New()

	// mw
	engine.Use(middleware.RequestIdMiddleware())
	engine.Use(middleware.LoggerMiddleware())
	engine.Use(middleware.RecoveryMiddleware())
	engine.Use(middleware.LimiterMiddleware())
	engine.Use(middleware.CorsMiddleware())

	// route
	if gin.Mode() == gin.DebugMode {
		xgin.WrapPprofSilently(engine)
	}
	api.RegisterSwag()
	engine.GET("/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("doc.json")))
	engine.GET("/v1/swagger", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/v1/swagger/index.html") })
	initRoute(engine)

	return &Server{engine: engine, config: cfg}
}

func (s *Server) Serve() {
	addr := fmt.Sprintf("0.0.0.0:%d", s.config.Meta.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	closeCh := make(chan int)
	go func() {
		signalCh := make(chan os.Signal)
		signal.Notify(signalCh, os.Interrupt)
		sig := <-signalCh
		log.Printf("Shutdown server by %s(%#x)", sig.String(), int(sig.(syscall.Signal)))

		err := server.Shutdown(context.Background())
		if err != nil {
			log.Fatalln("Failed to shutdown HTTP server:", err)
		}
		closeCh <- 0
	}()

	time.Sleep(500 * time.Millisecond)
	if e := os.Getenv("HTTP_PROXY"); e != "" {
		log.Printf("Using env HTTP_PROXY as %s", e)
	}
	if e := os.Getenv("HTTPS_PROXY"); e != "" {
		log.Printf("Using env HTTPS_PROXY as %s", e)
	}
	log.Printf("Listening and serving HTTP on %s", addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln("Failed to serve:", err)
	}

	<-closeCh
	log.Println("HTTP server exiting...")
}
