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
		goapidoc.NewOperation("GET", "/v1/manga", "Get all mangas").
			Desc("order by popular / new / update").
			Tags("Manga").
			Params(param.ParamPage, param.ParamOrder).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaDto>>")),

		goapidoc.NewOperation("GET", "/v1/manga/{mid}", "Get manga").
			Tags("Manga").
			Params(goapidoc.NewPathParam("mid", "integer#int64", true, "manga id")).
			Responses(goapidoc.NewResponse(200, "_Result<MangaDto>")),

		goapidoc.NewOperation("GET", "/v1/manga/random", "Get random manga").
			Tags("Manga").
			Responses(goapidoc.NewResponse(200, "_Result<RandomMangaInfoDto>")),

		goapidoc.NewOperation("GET", "/v1/manga/{mid}/{cid}", "Get manga chapter").
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
		config:          xmodule.MustGetByName(sn.SConfig).(*config.Config),
		mangaService:    xmodule.MustGetByName(sn.SMangaService).(*service.MangaService),
		categoryService: xmodule.MustGetByName(sn.SCategoryService).(*service.CategoryService),
	}
}

// GET /v1/manga
func (m *MangaController) GetAllMangas(c *gin.Context) *result.Result {
	pa := param.BindQueryPageOrder(c)

	mangas, limit, total, err := m.categoryService.GetGenreMangas("all", "all", "all", "all", pa.Page, pa.Order) // popular / new / update
	if err != nil {
		return result.Error(errno.GetAllMangasError).SetError(err, c)
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
	id, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}

	manga, err := m.mangaService.GetMangaPage(id)
	if err != nil {
		return result.Error(errno.GetMangaError).SetError(err, c)
	} else if manga == nil {
		return result.Error(errno.MangaNotFoundError)
	}

	res := dto.BuildMangaDto(manga)
	return result.Ok().SetData(res)
}

// GET /v1/manga/random
func (m *MangaController) getRandomManga(c *gin.Context) *result.Result {
	info, err := m.mangaService.GetRandomMangaInfo()
	if err != nil || info == nil {
		return result.Error(errno.GetRandomMangaError).SetError(err, c)
	}

	res := dto.BuildRandomMangaInfoDto(info)
	return result.Ok().SetData(res)
}

// GET /v1/manga/:mid/:cid
func (m *MangaController) GetMangaChapter(c *gin.Context) *result.Result {
	id, err := param.BindRouteID(c, "mid")
	cid, err2 := param.BindRouteID(c, "cid")
	if err != nil || err2 != nil {
		return result.BindingError(err, c)
	}

	chapter, err := m.mangaService.GetMangaChapter(id, cid)
	if err != nil {
		return result.Error(errno.GetMangaChapterError).SetError(err, c)
	} else if chapter == nil {
		return result.Error(errno.MangaChapterNotFoundError)
	}

	res := dto.BuildMangaChapterDto(chapter)
	return result.Ok().SetData(res)
}
