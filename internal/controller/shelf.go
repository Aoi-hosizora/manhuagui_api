package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewOperation("GET", "/v1/shelf", "Get shelf mangas").
			Tags("Shelf").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				param.ParamPage,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<ShelfMangaDto>>")),

		goapidoc.NewOperation("GET", "/v1/shelf/{mid}", "Check manga in shelf").
			Tags("Shelf").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<ShelfStatusDto>")),

		goapidoc.NewOperation("POST", "/v1/shelf/{mid}", "Save manga to shelf").
			Tags("Shelf").
			Params(
				goapidoc.NewHeaderParam("Authorization", "string", true, "access token"),
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "Result")),

		goapidoc.NewOperation("DELETE", "/v1/shelf/{mid}", "Remove manga from shelf").
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
		config:       xmodule.MustGetByName(sn.SConfig).(*config.Config),
		userService:  xmodule.MustGetByName(sn.SUserService).(*service.UserService),
		shelfService: xmodule.MustGetByName(sn.SShelfService).(*service.ShelfService),
	}
}

// GET /v1/shelf
func (s *ShelfController) GetShelfMangas(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	pa := param.BindQueryPage(c)

	mangas, limit, total, err := s.shelfService.GetShelfMangas(token, pa.Page)
	if err != nil {
		return result.Error(errno.GetShelfMangasError).SetError(err, c)
	} else if mangas == nil {
		return result.Error(errno.UnauthorizedError)
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
	mid, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}

	ok, _, err := s.userService.CheckLogin(token)
	if err != nil {
		return result.Error(errno.CheckLoginError).SetError(err, c)
	} else if !ok {
		return result.Error(errno.UnauthorizedError)
	}

	status, err := s.shelfService.CheckMangaInShelf(token, mid)
	if err != nil {
		return result.Error(errno.CheckMangaShelfError).SetError(err, c)
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
	mid, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}

	auth, existed, err := s.shelfService.SaveMangaToShelf(token, mid)
	if err != nil {
		return result.Error(errno.SaveMangaToShelfError).SetError(err, c)
	} else if !auth {
		return result.Error(errno.UnauthorizedError)
	}

	if existed {
		return result.Error(errno.MangaAlreadyInShelfError)
	}
	return result.Ok()
}

// DELETE /v1/shelf/:mid
func (s *ShelfController) RemoveMangaFromShelf(c *gin.Context) *result.Result {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	mid, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}

	auth, notFound, err := s.shelfService.RemoveMangaFromShelf(token, mid)
	if err != nil {
		return result.Error(errno.RemoveMangaFromShelfError).SetError(err, c)
	} else if !auth {
		return result.Error(errno.UnauthorizedError)
	}

	if notFound {
		return result.Error(errno.MangaNotInShelfYetError)
	}
	return result.Ok()
}
