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
		goapidoc.NewOperation("GET", "/v1/author", "Get all authors").
			Desc("order by popular / comic / update").
			Tags("Author").
			Params(
				goapidoc.NewQueryParam("genre", "string", false, "author genre"),
				goapidoc.NewQueryParam("zone", "string", false, "author zone"),
				goapidoc.NewQueryParam("age", "string", false, "author age range, (shaonv|shaonian|qingnian|ertong|tongyong)"),
				param.ParamPage, param.ParamOrder,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallAuthorDto>>")),

		goapidoc.NewOperation("GET", "/v1/author/{aid}", "Get author").
			Tags("Author").
			Params(goapidoc.NewPathParam("aid", "integer#int64", false, "author id")).
			Responses(goapidoc.NewResponse(200, "_Result<AuthorDto>")),

		goapidoc.NewOperation("GET", "/v1/author/{aid}/manga", "Get author mangas").
			Desc("order by popular / new / update").
			Tags("Author").
			Params(
				goapidoc.NewPathParam("aid", "integer#int64", true, "author id"),
				param.ParamPage, param.ParamOrder,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallMangaDto>>")),
	)
}

type AuthorController struct {
	config        *config.Config
	authorService *service.AuthorService
}

func NewAuthorController() *AuthorController {
	return &AuthorController{
		config:        xmodule.MustGetByName(sn.SConfig).(*config.Config),
		authorService: xmodule.MustGetByName(sn.SAuthorService).(*service.AuthorService),
	}
}

// GET /v1/author
func (a *AuthorController) GetAllAuthors(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, a.config)
	genre := c.Query("genre")
	zone := c.Query("zone")
	age := c.Query("age")

	authors, limit, total, err := a.authorService.GetAllAuthors(genre, zone, age, pa.Page, pa.Order) // popular / comic / update
	if err != nil {
		return result.Error(errno.GetAllAuthorsError).SetError(err, c)
	}

	res := dto.BuildSmallAuthorDtos(authors)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}

// GET /v1/author/:aid
func (a *AuthorController) GetAuthor(c *gin.Context) *result.Result {
	aid, err := param.BindRouteId(c, "aid")
	if err != nil {
		return result.Error(errno.RequestParamError).SetError(err, c)
	}

	author, err := a.authorService.GetAuthor(aid)
	if err != nil {
		return result.Error(errno.GetAuthorError).SetError(err, c)
	} else if author == nil {
		return result.Error(errno.AuthorNotFound)
	}

	res := dto.BuildAuthorDto(author)
	return result.Ok().SetData(res)
}

// GET /v1/author/:aid/manga
func (a *AuthorController) GetAuthorMangas(c *gin.Context) *result.Result {
	aid, err := param.BindRouteId(c, "aid")
	if err != nil {
		return result.Error(errno.RequestParamError).SetError(err, c)
	}
	pa := param.BindPageOrder(c, a.config)

	mangas, limit, total, err := a.authorService.GetAuthorMangas(aid, pa.Page, pa.Order) // popular / new / update
	if err != nil {
		return result.Error(errno.GetAuthorMangasError).SetError(err, c)
	} else if mangas == nil {
		return result.Error(errno.AuthorNotFound)
	}

	res := dto.BuildSmallMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
