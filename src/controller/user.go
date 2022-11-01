package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-api/src/config"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-api/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddRoutePaths(
		goapidoc.NewRoutePath("POST", "/v1/user/login", "Login").
			Tags("User").
			Params(
				goapidoc.NewQueryParam("username", "string", true, "login username"),
				goapidoc.NewQueryParam("password", "string", true, "login password"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<TokenDto>")),

		goapidoc.NewRoutePath("POST", "/v1/user/check_login", "Check login").
			Tags("User").
			Params(goapidoc.NewHeaderParam("Authorization", "string", true, "access token")).
			Responses(goapidoc.NewResponse(200, "_Result<UsernameDto>")),

		goapidoc.NewRoutePath("GET", "/v1/user/info", "Get authorized user information").
			Tags("User").
			Params(goapidoc.NewHeaderParam("Authorization", "string", true, "access token")).
			Responses(goapidoc.NewResponse(200, "_Result<UserDto>")),

		goapidoc.NewRoutePath("GET", "/v1/user/manga/{mid}/{cid}", "Record manga for the authorized user").
			Tags("User").
			Deprecated(true).
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
			).
			Responses(goapidoc.NewResponse(200, "Result")),

		goapidoc.NewRoutePath("POST", "/v1/user/manga/{mid}/{cid}", "Record manga for the authorized user").
			Tags("User").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
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
		config:      xdi.GetByNameForce(sn.SConfig).(*config.Config),
		userService: xdi.GetByNameForce(sn.SUserService).(*service.UserService),
	}
}

// POST /v1/user/login
func (u *UserController) Login(c *gin.Context) *result.Result {
	username := c.Query("username")
	password := c.Query("password")
	token, err := u.userService.Login(username, password)
	if err != nil {
		return result.Error(exception.LoginError).SetError(err, c)
	} else if token == "" {
		return result.Error(exception.PasswordError)
	}

	return result.Ok().SetData(&dto.TokenDto{Token: token})
}

// POST /v1/user/check_login
func (u *UserController) CheckLogin(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	ok, username, err := u.userService.CheckLogin(token)
	if err != nil {
		return result.Error(exception.CheckLoginError).SetError(err, c)
	} else if !ok {
		return result.Error(exception.UnauthorizedError)
	}

	return result.Ok().SetData(&dto.UsernameDto{Username: username})
}

// GET /v1/user/info
func (u *UserController) GetUser(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	user, err := u.userService.GetUser(token)
	if err != nil {
		return result.Error(exception.GetUserError).SetError(err, c)
	} else if user == nil {
		return result.Error(exception.UnauthorizedError)
	}

	res := dto.BuildUserDto(user)
	return result.Ok().SetData(res)
}

// POST/GET /v1/user/manga/:mid/:cid
func (u *UserController) RecordManga(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}
	cid, err := param.BindRouteId(c, "cid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	ok, _, err := u.userService.CheckLogin(token)
	if err != nil {
		return result.Error(exception.CheckLoginError).SetError(err, c)
	} else if !ok {
		return result.Error(exception.UnauthorizedError)
	}

	err = u.userService.RecordManga(token, mid, cid)
	if err != nil {
		return result.Error(exception.CountMangaError).SetError(err, c)
	}
	return result.Ok()
}
