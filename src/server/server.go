package server

import (
	"context"
	"fmt"
	"github.com/Aoi-hosizora/ahlib-web/xgin"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/docs"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/Aoi-hosizora/manhuagui-backend/src/middleware"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	goapidoc.SetDocument(
		"localhost:10018", "/",
		goapidoc.NewInfo("manhuagui-backend", "An unofficial backend for manhuagui written in golang/gin", "1.0").
			Contact(goapidoc.NewContact("Aoi-hosizora", "https://github.com/Aoi-hosizora", "aoihosizora@hotmail.com")),
	)

	goapidoc.SetTags(
		goapidoc.NewTag("Manga", "manga-controller"),
	)
}

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func NewServer() *Server {
	cfg := xdi.GetByNameForce(sn.SConfig).(*config.Config)
	gin.SetMode(cfg.Meta.RunMode)
	engine := gin.New()

	// mw
	engine.Use(middleware.RequestIdMiddleware())
	engine.Use(middleware.LoggerMiddleware())
	engine.Use(middleware.RecoveryMiddleware())
	engine.Use(middleware.CorsMiddleware())

	// route
	if gin.Mode() == gin.DebugMode {
		xgin.PprofWrap(engine)
	}
	docs.RegisterSwag()
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

	log.Printf("Listening and serving HTTP on %s", addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln("Failed to serve:", err)
	}

	<-closeCh
	log.Println("HTTP server exiting...")
}
