package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/vo"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewOperation("GET", "/v1/category", "Get categories").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<CategoryListDto>")),

		goapidoc.NewOperation("GET", "/v1/category/genre", "Get genres").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CategoryDto>>")),

		goapidoc.NewOperation("GET", "/v1/category/zone", "Get zones").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CategoryDto>>")),

		goapidoc.NewOperation("GET", "/v1/category/age", "Get ages").
			Tags("Category").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CategoryDto>>")),

		goapidoc.NewOperation("GET", "/v1/category/genre/{genre}", "Get genre mangas").
			Desc("order by popular / new / update").
			Tags("Category").
			Params(
				goapidoc.NewPathParam("genre", "string", true, "genre name, (all|...)"),
				goapidoc.NewQueryParam("zone", "string", false, "manga zone, (all|japan|hongkong|other|europe|china|korea)"),
				goapidoc.NewQueryParam("age", "string", false, "manga age, (all|shaonv|shaonian|qingnian|ertong|tongyong)"),
				goapidoc.NewQueryParam("status", "string", false, "manga status, (all|lianzai|wanjie)"),
				param.ParamPage, param.ParamOrder,
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
		config:          xmodule.MustGetByName(sn.SConfig).(*config.Config),
		categoryService: xmodule.MustGetByName(sn.SCategoryService).(*service.CategoryService),
	}
}

// GET /v1/category
func (ca *CategoryController) GetCategories(c *gin.Context) *result.Result {
	categories, err := ca.categoryService.GetAllCategories()
	if err != nil {
		return result.Error(errno.GetCategoriesError).SetError(err, c)
	}

	res := dto.BuildCategoryListDto(categories)
	return result.Ok().SetData(res)
}

// GET /v1/category/genre
func (ca *CategoryController) GetGenres(c *gin.Context) *result.Result {
	genres, err := ca.categoryService.GetGenres()
	if err != nil {
		return result.Error(errno.GetGenresError).SetError(err, c)
	}

	res := dto.BuildCategoryDtos(genres)
	return result.Ok().SetPage(1, int32(len(res)), int32(len(res)), res)
}

// GET /v1/category/zone
func (ca *CategoryController) GetZones(c *gin.Context) *result.Result {
	zones, err := ca.categoryService.GetZones()
	if err != nil {
		return result.Error(errno.GetZonesError).SetError(err, c)
	}

	res := dto.BuildCategoryDtos(zones)
	return result.Ok().SetPage(1, int32(len(res)), int32(len(res)), res)
}

// GET /v1/category/age
func (ca *CategoryController) GetAges(c *gin.Context) *result.Result {
	zones, err := ca.categoryService.GetAges()
	if err != nil {
		return result.Error(errno.GetAgesError).SetError(err, c)
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
		return result.Error(errno.GetGenreMangasError).SetError(err, c)
	} else if mangas == nil { // not found
		return result.Error(errno.GenreNotFoundError)
	} else if len(mangas) == 0 { // empty
		res := dto.BuildTinyMangaDtos([]*vo.TinyManga{})
		return result.Ok().SetPage(pa.Page, limit, 0, res)
	}

	res := dto.BuildTinyMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
