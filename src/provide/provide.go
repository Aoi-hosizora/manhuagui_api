package provide

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/logger"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
	"log"
)

func Provide(configPath string) error {
	// *config.Config
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}
	xdi.ProvideName(sn.SConfig, cfg)

	// *logrus.Logger
	lgr, err := logger.Setup()
	if err != nil {
		log.Fatalln("Failed to setup logger:", err)
	}
	xdi.ProvideName(sn.SLogger, lgr)

	// ///////////////////////////////////////////////////////////////////////

	// services
	xdi.ProvideName(sn.SHttpService, service.NewHttpService())           // *service.HttpService
	xdi.ProvideName(sn.SCategoryService, service.NewCategoryService())   // *service.CategoryService
	xdi.ProvideName(sn.SMangaService, service.NewMangaService())         // *service.MangaService
	xdi.ProvideName(sn.SMangaListService, service.NewMangaListService()) // *service.MangaListService

	return nil
}
