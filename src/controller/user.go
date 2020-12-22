package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/dto"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/param"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
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
			Responses(goapidoc.NewResponse(200, "Result")),

		goapidoc.NewRoutePath("GET", "/v1/user/info", "Get authorized user information").
			Tags("User").
			Params(goapidoc.NewHeaderParam("Authorization", "string", true, "access token")).
			Responses(goapidoc.NewResponse(200, "_Result<UserDto>")),

		goapidoc.NewRoutePath("GET", "/v1/user/shelf", "Get shelf mangas").
			Tags("User").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				param.ParamPage,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<ShelfMangaDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/user/shelf/{mid}", "Check manga in shelf").
			Tags("User").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<ShelfStatusDto>")),

		goapidoc.NewRoutePath("POST", "/v1/user/shelf/{mid}", "Save manga to shelf").
			Tags("User").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "Result")),

		goapidoc.NewRoutePath("DELETE", "/v1/user/shelf/{mid}", "Remove manga from shelf").
			Tags("User").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
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
	ok, err := u.userService.CheckLogin(token)
	if err != nil {
		return result.Error(exception.CheckLoginError).SetError(err, c)
	} else if !ok {
		return result.Error(exception.UnauthorizedError)
	}

	return result.Ok()
}

// GET /v1/user/info
func (u *UserController) GetUser(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	user, err := u.userService.GetUser(token)
	if err != nil {
		return result.Error(exception.GetUserError).SetError(err, c)
	} else if user == nil {
		return result.Error(exception.UnauthorizedError)
	}

	res := dto.BuildUserDto(user)
	return result.Ok().SetData(res)
}

// GET /v1/user/shelf
func (u *UserController) GetShelfMangas(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	pa := param.BindPage(c, u.config)

	mangas, limit, total, err := u.userService.GetShelfMangas(token, pa.Page)
	if err != nil {
		return result.Error(exception.GetShelfMangasError).SetError(err, c)
	} else if mangas == nil {
		return result.Error(exception.UnauthorizedError)
	}

	res := dto.BuildShelfMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}

// GET /v1/user/shelf/:mid
func (u *UserController) CheckMangaInShelf(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	auth, in, err := u.userService.CheckMangaInShelf(token, mid)
	if err != nil {
		return result.Error(exception.CheckMangaShelfError).SetError(err, c)
	} else if !auth {
		return result.Error(exception.UnauthorizedError)
	}

	sts := &vo.ShelfStatus{In: in}
	res := dto.BuildShelfStatusDto(sts)
	return result.Ok().SetData(res)
}

// POST /v1/user/shelf/:mid
func (u *UserController) SaveMangaToShelf(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	auth, existed, err := u.userService.SaveMangaToShelf(token, mid)
	if err != nil {
		return result.Error(exception.SaveMangaToShelfError).SetError(err, c)
	} else if !auth {
		return result.Error(exception.UnauthorizedError)
	}

	if existed {
		return result.Error(exception.MangaAlreadyInShelfError)
	}
	return result.Ok()
}

// DELETE /v1/user/shelf/:mid
func (u *UserController) RemoveMangaFromShelf(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	auth, notFound, err := u.userService.RemoveMangaFromShelf(token, mid)
	if err != nil {
		return result.Error(exception.RemoveMangaFromShelfError).SetError(err, c)
	}else if !auth {
		return result.Error(exception.UnauthorizedError)
	}

	if notFound {
		return result.Error(exception.MangaNotInShelfYetError)
	}
	return result.Ok()
}
