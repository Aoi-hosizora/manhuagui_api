package module

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/logger"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"log"
)

func Provide(configPath string) error {
	// *config.Config
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}
	xmodule.ProvideByName(sn.SConfig, cfg)

	// *logrus.Logger
	lgr, err := logger.Setup()
	if err != nil {
		log.Fatalln("Failed to setup logger:", err)
	}
	xmodule.ProvideByName(sn.SLogger, lgr)

	// ///////////////////////////////////////////////////////////////////////

	// services
	xmodule.ProvideByName(sn.SHttpService, service.NewHttpService())           // *service.HttpService
	xmodule.ProvideByName(sn.SCategoryService, service.NewCategoryService())   // *service.CategoryService
	xmodule.ProvideByName(sn.SAuthorService, service.NewAuthorService())       // *service.AuthorService
	xmodule.ProvideByName(sn.SMangaService, service.NewMangaService())         // *service.MangaService
	xmodule.ProvideByName(sn.SRankService, service.NewRankService())           // *service.RankService
	xmodule.ProvideByName(sn.SMangaListService, service.NewMangaListService()) // *service.MangaListService
	xmodule.ProvideByName(sn.SSearchService, service.NewSearchService())       // *service.SearchService
	xmodule.ProvideByName(sn.SCommentService, service.NewCommentService())     // *service.CommentService
	xmodule.ProvideByName(sn.SUserService, service.NewUserService())           // *service.UserService
	xmodule.ProvideByName(sn.SShelfService, service.NewShelfService())         // *service.ShelfService
	xmodule.ProvideByName(sn.SMessageService, service.NewMessageService())     // *service.MessageService

	return nil
}
