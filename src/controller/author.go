package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
)

type AuthorController struct {
	config        *config.Config
	authorService *service.AuthorService
}

func NewAuthorController() *AuthorController {
	return &AuthorController{
		config:        xdi.GetByNameForce(sn.SConfig).(*config.Config),
		authorService: xdi.GetByNameForce(sn.SAuthorService).(*service.AuthorService),
	}
}
