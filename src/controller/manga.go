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
		goapidoc.NewRoutePath("GET", "/v1/manga", "Get all manga pages").
			Desc("order by popular / new / update").
			Tags("Manga").
			Params(param.ADPage, param.ADOrder).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/manga/{mid}", "Get manga page").
			Tags("Manga").
			Params(goapidoc.NewPathParam("mid", "integer#int64", true, "manga id")).
			Responses(goapidoc.NewResponse(200, "_Result<MangaDto>")),

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
func (m *MangaController) GetAllMangaPages(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, m.config)

	mangas, limit, total, err := m.categoryService.GetGenreMangas("all", "all", "all", "all", pa.Page, pa.Order) // popular / new / update
	if err != nil {
		return result.Error(exception.GetAllMangaPagesError).SetError(err, c)
	}

	res := dto.BuildTinyMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}

// GET /v1/manga/:mid
func (m *MangaController) GetMangaPage(c *gin.Context) *result.Result {
	id, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	mangas, err := m.mangaService.GetMangaPage(id)
	if err != nil {
		return result.Error(exception.GetMangaPageError).SetError(err, c)
	} else if mangas == nil {
		return result.Error(exception.MangaPageNotFoundError)
	}

	res := dto.BuildMangaDto(mangas)
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
