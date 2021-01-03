package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/dto"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/param"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddRoutePaths(
		goapidoc.NewRoutePath("GET", "/v1/shelf", "Get shelf mangas").
			Tags("Shelf").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				param.ParamPage,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<ShelfMangaDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/shelf/{mid}", "Check manga in shelf").
			Tags("Shelf").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<ShelfStatusDto>")),

		goapidoc.NewRoutePath("POST", "/v1/shelf/{mid}", "Save manga to shelf").
			Tags("Shelf").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "Result")),

		goapidoc.NewRoutePath("DELETE", "/v1/shelf/{mid}", "Remove manga from shelf").
			Tags("Shelf").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "Result")),
	)
}

type ShelfController struct {
	config       *config.Config
	userService  *service.UserService
	shelfService *service.ShelfService
}

func NewShelfController() *ShelfController {
	return &ShelfController{
		config:       xdi.GetByNameForce(sn.SConfig).(*config.Config),
		userService:  xdi.GetByNameForce(sn.SUserService).(*service.UserService),
		shelfService: xdi.GetByNameForce(sn.SShelfService).(*service.ShelfService),
	}
}

// GET /v1/shelf
func (s *ShelfController) GetShelfMangas(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	pa := param.BindPage(c, s.config)

	mangas, limit, total, err := s.shelfService.GetShelfMangas(token, pa.Page)
	if err != nil {
		return result.Error(exception.GetShelfMangasError).SetError(err, c)
	} else if mangas == nil {
		return result.Error(exception.UnauthorizedError)
	}

	res := dto.BuildShelfMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}

// GET /v1/shelf/:mid
func (s *ShelfController) CheckMangaInShelf(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	ok, _, err := s.userService.CheckLogin(token)
	if err != nil {
		return result.Error(exception.CheckLoginError).SetError(err, c)
	} else if !ok {
		return result.Error(exception.UnauthorizedError)
	}

	status, err := s.shelfService.CheckMangaInShelf(token, mid)
	if err != nil {
		return result.Error(exception.CheckMangaShelfError).SetError(err, c)
	}

	res := dto.BuildShelfStatusDto(status)
	return result.Ok().SetData(res)
}

// POST /v1/shelf/:mid
func (s *ShelfController) SaveMangaToShelf(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	auth, existed, err := s.shelfService.SaveMangaToShelf(token, mid)
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

// DELETE /v1/shelf/:mid
func (s *ShelfController) RemoveMangaFromShelf(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	auth, notFound, err := s.shelfService.RemoveMangaFromShelf(token, mid)
	if err != nil {
		return result.Error(exception.RemoveMangaFromShelfError).SetError(err, c)
	} else if !auth {
		return result.Error(exception.UnauthorizedError)
	}

	if notFound {
		return result.Error(exception.MangaNotInShelfYetError)
	}
	return result.Ok()
}
