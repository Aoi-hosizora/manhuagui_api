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
		goapidoc.NewRoutePath("GET", "/v1/category/genre", "Get genres").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CategoryDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/category/zone", "Get zones").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CategoryDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/category/age", "Get ages").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CategoryDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/category/genre/{genre}", "Get genre mangas").
			Desc("order by popular / new / update").
			Tags("Category").
			Params(
				goapidoc.NewPathParam("genre", "string", true, "genre name, (all|...)"),
				goapidoc.NewQueryParam("zone", "string", false, "manga zone, (all|japan|hongkong|other|europe|china|korea)"),
				goapidoc.NewQueryParam("age", "string", false, "manga age, (all|shaonv|shaonian|qingnian|ertong|tongyong)"),
				goapidoc.NewQueryParam("status", "string", false, "manga status, (all|lianzai|wanjie)"),
				param.ADPage, param.ADOrder,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaDto>>")),
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

// GET /v1/category/zone
func (ca *CategoryController) GetZones(c *gin.Context) *result.Result {
	zones, err := ca.categoryService.GetZones()
	if err != nil {
		return result.Error(exception.GetZonesError).SetError(err, c)
	}

	res := dto.BuildCategoryDtos(zones)
	return result.Ok().SetPage(1, int32(len(res)), int32(len(res)), res)
}

// GET /v1/category/age
func (ca *CategoryController) GetAges(c *gin.Context) *result.Result {
	zones, err := ca.categoryService.GetAges()
	if err != nil {
		return result.Error(exception.GetAgesError).SetError(err, c)
	}

	res := dto.BuildCategoryDtos(zones)
	return result.Ok().SetPage(1, int32(len(res)), int32(len(res)), res)
}

// GET /v1/category/genre/:genre
func (ca *CategoryController) GetGenreMangas(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, ca.config)
	genre := c.Param("genre")
	zone := c.Query("zone")
	age := c.Query("age")
	status := c.Query("status")

	// zone > genre > age > status
	mangas, limit, total, err := ca.categoryService.GetGenreMangas(genre, zone, age, status, pa.Page, pa.Order) // popular / new / update
	if err != nil {
		return result.Error(exception.GetGenreMangasError).SetError(err, c)
	} else if mangas == nil { // not found
		return result.Error(exception.GenreNotFoundError)
	} else if len(mangas) == 0 { // empty
		res := dto.BuildTinyMangaDtos([]*vo.TinyManga{})
		return result.Ok().SetPage(pa.Page, limit, 0, res)
	}

	res := dto.BuildTinyMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
