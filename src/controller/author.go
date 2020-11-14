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
		goapidoc.NewRoutePath("GET", "/v1/author", "Get all authors").
			Tags("Author").
			Params(
				goapidoc.NewQueryParam("genre", "string", false, "author genre"),
				goapidoc.NewQueryParam("zone", "string", false, "author zone"),
				goapidoc.NewQueryParam("age", "string", false, "author age range, (shaonv|shaonian|qingnian|ertong|tongyong)"),
				param.ADPage, param.ADOrder,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallAuthorDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/author/{aid}", "Get author").
			Tags("Author").
			Params(goapidoc.NewPathParam("aid", "integer#int64", false, "author id")).
			Responses(goapidoc.NewResponse(200, "_Result<AuthorDto>")),

		goapidoc.NewRoutePath("GET", "/v1/author/{aid}/manga", "Get author").
			Tags("Author").
			Params(
				goapidoc.NewPathParam("aid", "integer#int64", true, "author id"),
				param.ADPage, param.ADOrder,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallMangaPageDto>>")),
	)
}

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

// GET /v1/author
func (a *AuthorController) GetAllAuthors(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, a.config)
	genre := c.Query("genre")
	zone := c.Query("zone")
	age := c.Query("age")

	authors, limit, total, err := a.authorService.GetAllAuthors(genre, zone, age, pa.Page, pa.Order == "popular")
	if err != nil {
		return result.Error(exception.GetAllAuthorsError).SetError(err, c)
	}

	res := dto.BuildSmallAuthorDtos(authors)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}

// GET /v1/author/:aid
func (a *AuthorController) GetAuthor(c *gin.Context) *result.Result {
	aid, err := param.BindRouteId(c, "aid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	author, err := a.authorService.GetAuthor(aid)
	if err != nil {
		return result.Error(exception.GetAuthorError).SetError(err, c)
	} else if author == nil {
		return result.Error(exception.AuthorNotFound)
	}

	res := dto.BuildAuthorDto(author)
	return result.Ok().SetData(res)
}

// GET /v1/author/:aid/manga
func (a *AuthorController) GetAuthorMangas(c *gin.Context) *result.Result {
	aid, err := param.BindRouteId(c, "aid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}
	pa := param.BindPageOrder(c, a.config)

	mangas, limit, total, err := a.authorService.GetAuthorMangas(aid, pa.Page, pa.Order == "popular")
	if err != nil {
		return result.Error(exception.GetAuthorMangasError).SetError(err, c)
	} else if mangas == nil {
		return result.Error(exception.AuthorNotFound)
	}

	res := dto.BuildSmallMangaPageDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
