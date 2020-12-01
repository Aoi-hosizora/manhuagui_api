package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
)

type UserController struct {
	config      *config.Config
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		config:      xdi.GetByNameForce(sn.SConfig).(*config.Config),
		userService: xdi.GetByNameForce(sn.SUserService).(*service.UserService),
	}
}
