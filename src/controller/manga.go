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
		goapidoc.NewRoutePath("GET", "/v1/manga", "Get all mangas").
			Desc("order by popular / new / update").
			Tags("Manga").
			Params(param.ParamPage, param.ParamOrder).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/manga/{mid}", "Get manga").
			Tags("Manga").
			Params(goapidoc.NewPathParam("mid", "integer#int64", true, "manga id")).
			Responses(goapidoc.NewResponse(200, "_Result<MangaDto>")),

		goapidoc.NewRoutePath("GET", "/v1/manga/random", "Get random manga").
			Tags("Manga").
			Responses(goapidoc.NewResponse(200, "_Result<RandomMangaInfoDto>")),

		goapidoc.NewRoutePath("GET", "/v1/manga/{mid}/{cid}", "Get manga chapter").
			Tags("Manga").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<MangaChapterDto>")),
	)
}

type MangaController struct {
	config          *config.Config
	mangaService    *service.MangaService
	categoryService *service.CategoryService
}

func NewMangaController() *MangaController {
	return &MangaController{
		config:          xdi.GetByNameForce(sn.SConfig).(*config.Config),
		mangaService:    xdi.GetByNameForce(sn.SMangaService).(*service.MangaService),
		categoryService: xdi.GetByNameForce(sn.SCategoryService).(*service.CategoryService),
	}
}

// GET /v1/manga
func (m *MangaController) GetAllMangas(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, m.config)

	mangas, limit, total, err := m.categoryService.GetGenreMangas("all", "all", "all", "all", pa.Page, pa.Order) // popular / new / update
	if err != nil {
		return result.Error(exception.GetAllMangasError).SetError(err, c)
	}

	res := dto.BuildTinyMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}

// GET /v1/manga/...
func (m *MangaController) GetManga(c *gin.Context) *result.Result {
	if c.Param("mid") == "random" {
		return m.getRandomManga(c)
	}
	return m.getManga(c)
}

// GET /v1/manga/:mid
func (m *MangaController) getManga(c *gin.Context) *result.Result {
	id, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	manga, err := m.mangaService.GetMangaPage(id)
	if err != nil {
		return result.Error(exception.GetMangaError).SetError(err, c)
	} else if manga == nil {
		return result.Error(exception.MangaNotFoundError)
	}

	res := dto.BuildMangaDto(manga)
	return result.Ok().SetData(res)
}

// GET /v1/manga/random
func (m *MangaController) getRandomManga(c *gin.Context) *result.Result {
	info, err := m.mangaService.GetRandomMangaInfo()
	if err != nil || info == nil {
		return result.Error(exception.GetRandomMangaError).SetError(err, c)
	}

	res := dto.BuildRandomMangaInfoDto(info)
	return result.Ok().SetData(res)
}

// GET /v1/manga/:mid/:cid
func (m *MangaController) GetMangaChapter(c *gin.Context) *result.Result {
	id, err := param.BindRouteId(c, "mid")
	cid, err2 := param.BindRouteId(c, "cid")
	if err != nil || err2 != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	chapter, err := m.mangaService.GetMangaChapter(id, cid)
	if err != nil {
		return result.Error(exception.GetMangaChapterError).SetError(err, c)
	} else if chapter == nil {
		return result.Error(exception.MangaChapterNotFoundError)
	}

	res := dto.BuildMangaChapterDto(chapter)
	return result.Ok().SetData(res)
}
