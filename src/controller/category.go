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
		goapidoc.NewRoutePath("GET", "/v1/category/genre", "Get genres").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CategoryDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/category/genre/{name}", "Get genre mangas").
			Tags("Category").
			Params(
				goapidoc.NewPathParam("name", "string", true, "genre name"),
				param.ADPage, param.ADOrder,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaPageDto>>")),
	)
}

type CategoryController struct {
	config          *config.Config
	categoryService *service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		config:          xdi.GetByNameForce(sn.SConfig).(*config.Config),
		categoryService: xdi.GetByNameForce(sn.SCategoryService).(*service.CategoryService),
	}
}

// GET /v1/category/genre
func (ca *CategoryController) GetGenres(c *gin.Context) *result.Result {
	genres, err := ca.categoryService.GetGenres()
	if err != nil {
		return result.Error(exception.GetGenresError).SetError(err, c)
	}

	res := dto.BuildCategoryDtos(genres)
	return result.Ok().SetPage(1, int32(len(res)), int32(len(res)), res)
}

// GET /v1/category/genre/:name
func (ca *CategoryController) GetGenreMangas(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, ca.config)
	name := c.Param("name")
	mangas, limit, total, err := ca.categoryService.GetGenreMangas(name, pa.Page, pa.Order == "popular")
	if err != nil {
		return result.Error(exception.GetGenreMangasError).SetError(err, c)
	} else if mangas == nil {
		return result.Error(exception.GenreNotFoundError)
	}

	res := dto.BuildTinyMangaPageDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
