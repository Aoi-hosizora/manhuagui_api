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
		goapidoc.NewRoutePath("GET", "/v1/manga/{mid}", "Get manga page").
			Tags("Manga").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<MangaPageDto>")),

		goapidoc.NewRoutePath("GET", "/v1/manga/{mid}/{cid}", "Get manga chapter").
			Tags("Manga").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<MangaChapterDto>")),

		goapidoc.NewRoutePath("GET", "/v1/list/serial", "Get hot serial mangas").
			Tags("Manga").
			Responses(goapidoc.NewResponse(200, "_Result<MangaPageGroupListDto>")),

		goapidoc.NewRoutePath("GET", "/v1/list/finish", "Get finished mangas").
			Tags("Manga").
			Responses(goapidoc.NewResponse(200, "_Result<MangaPageGroupListDto>")),

		goapidoc.NewRoutePath("GET", "/v1/list/latest", "Get latest mangas").
			Tags("Manga").
			Responses(goapidoc.NewResponse(200, "_Result<MangaPageGroupListDto>")),

		goapidoc.NewRoutePath("GET", "/v1/list/updated", "Get latest mangas").
			Tags("Manga").
			Params(param.ADPage, param.ADLimit).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaPageDto>>")),
	)
}

type MangaController struct {
	config           *config.Config
	mangaService     *service.MangaService
	mangaListService *service.MangaListService
}

func NewMangaController() *MangaController {
	return &MangaController{
		config:           xdi.GetByNameForce(sn.SConfig).(*config.Config),
		mangaService:     xdi.GetByNameForce(sn.SMangaService).(*service.MangaService),
		mangaListService: xdi.GetByNameForce(sn.SMangaListService).(*service.MangaListService),
	}
}

// GET /v1/manga/:mid
func (m *MangaController) GetMangaPage(c *gin.Context) *result.Result {
	id, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	page, err := m.mangaService.GetMangaPage(id)
	if err != nil {
		return result.Error(exception.GetMangaPageError).SetError(err, c)
	} else if page == nil {
		return result.Error(exception.MangaPageNotFoundError)
	}

	res := dto.BuildMangaPageDto(page)
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

// GET /v1/list/serial
func (m *MangaController) GetHotSerialMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetHotSerialMangas()
	if err != nil {
		return result.Error(exception.GetHotSerialMangasError).SetError(err, c)
	}

	res := dto.BuildMangaPageGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/finish
func (m *MangaController) GetFinishedMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetFinishedMangas()
	if err != nil {
		return result.Error(exception.GetFinishedMangasError).SetError(err, c)
	}

	res := dto.BuildMangaPageGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/latest
func (m *MangaController) GetLatestMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetLatestMangas()
	if err != nil {
		return result.Error(exception.GetLatestMangasError).SetError(err, c)
	}

	res := dto.BuildMangaPageGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/updated
func (m *MangaController) GetRecentUpdatedMangas(c *gin.Context) *result.Result {
	pa := param.BindPage(c, m.config)
	pages, tot, err := m.mangaListService.GetRecentUpdatedMangas(pa)
	if err != nil {
		return result.Error(exception.GetUpdatedMangasError).SetError(err, c)
	}

	res := dto.BuildTinyMangaPageDtos(pages)
	return result.Ok().SetPage(pa.Page, pa.Limit, tot, res)
}
