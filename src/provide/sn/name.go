package sn

import "github.com/Aoi-hosizora/ahlib/xdi"

const (
	// common
	SConfig xdi.ServiceName = "config" // *config.Config
	SLogger xdi.ServiceName = "logger" // *logrus.Logger

	// service
	SHttpService      xdi.ServiceName = "http-service"       // *service.HttpService
	SMangaService     xdi.ServiceName = "manga-service"      // *service.MangaService
	SMangaListService xdi.ServiceName = "manga-list-service" // *service.MangaListService
	SCategoryService  xdi.ServiceName = "category-service"   // *service.CategoryService
	SSearchService    xdi.ServiceName = "search-service"     // *service.SearchService
	SAuthorService    xdi.ServiceName = "author-service"     // *service.AuthorService
)
