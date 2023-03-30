package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/apidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewOperation("POST", "/v1/user/login", "Login").
			Tags("User").
			Params(
				goapidoc.NewQueryParam("username", "string", true, "login username"),
				goapidoc.NewQueryParam("password", "string", true, "login password"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<TokenDto>")),

		goapidoc.NewOperation("POST", "/v1/user/check_login", "Check login").
			Tags("User").
			Params(apidoc.ParamToken).
			Responses(goapidoc.NewResponse(200, "_Result<UsernameDto>")),

		goapidoc.NewOperation("GET", "/v1/user/info", "Get authorized user information").
			Tags("User").
			Params(apidoc.ParamToken).
			Responses(goapidoc.NewResponse(200, "_Result<UserDto>")),

		goapidoc.NewOperation("GET", "/v1/user/manga/{mid}/{cid}", "Record manga for the authorized user").
			Tags("User").
			Deprecated(true).
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
				apidoc.ParamToken,
			).
			Responses(goapidoc.NewResponse(200, "Result")),

		goapidoc.NewOperation("POST", "/v1/user/manga/{mid}/{cid}", "Record manga for the authorized user").
			Tags("User").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
				apidoc.ParamToken,
			).
			Responses(goapidoc.NewResponse(200, "Result")),
	)
}

type UserController struct {
	config      *config.Config
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		config:      xmodule.MustGetByName(sn.SConfig).(*config.Config),
		userService: xmodule.MustGetByName(sn.SUserService).(*service.UserService),
	}
}

// POST /v1/user/login
func (u *UserController) Login(c *gin.Context) *result.Result {
	username := c.Query("username")
	password := c.Query("password")
	token, err := u.userService.Login(username, password)
	if err != nil {
		return result.Error(errno.LoginError).SetError(err, c)
	} else if token == "" {
		return result.Error(errno.PasswordError)
	}

	return result.Ok().SetData(&dto.TokenDto{Token: token})
}

// POST /v1/user/check_login
func (u *UserController) CheckLogin(c *gin.Context) *result.Result {
	token := param.BindToken(c)
	ok, username, err := u.userService.CheckLogin(token)
	if err != nil {
		return result.Error(errno.CheckLoginError).SetError(err, c)
	} else if !ok {
		return result.Error(errno.UnauthorizedError)
	}

	return result.Ok().SetData(&dto.UsernameDto{Username: username})
}

// GET /v1/user/info
func (u *UserController) GetUser(c *gin.Context) *result.Result {
	token := param.BindToken(c)
	user, err := u.userService.GetUser(token)
	if err != nil {
		return result.Error(errno.GetUserError).SetError(err, c)
	} else if user == nil {
		return result.Error(errno.UnauthorizedError)
	}

	res := dto.BuildUserDto(user)
	return result.Ok().SetData(res)
}

// POST /v1/user/manga/:mid/:cid
// GET /v1/user/manga/:mid/:cid (deprecated)
func (u *UserController) RecordManga(c *gin.Context) *result.Result {
	token := param.BindToken(c)
	mid, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}
	cid, err := param.BindRouteID(c, "cid")
	if err != nil {
		return result.BindingError(err, c)
	}

	ok, _, err := u.userService.CheckLogin(token)
	if err != nil {
		return result.Error(errno.CheckLoginError).SetError(err, c)
	} else if !ok {
		return result.Error(errno.UnauthorizedError)
	}

	err = u.userService.RecordManga(token, mid, cid)
	if err != nil {
		return result.Error(errno.CountMangaError).SetError(err, c)
	}
	return result.Ok()
}
