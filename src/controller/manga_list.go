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
		goapidoc.NewRoutePath("GET", "/v1/list/serial", "Get hot serial mangas").
			Tags("MangaList").
			Responses(goapidoc.NewResponse(200, "_Result<MangaGroupListDto>")),

		goapidoc.NewRoutePath("GET", "/v1/list/finish", "Get finished mangas").
			Tags("MangaList").
			Responses(goapidoc.NewResponse(200, "_Result<MangaGroupListDto>")),

		goapidoc.NewRoutePath("GET", "/v1/list/latest", "Get latest mangas").
			Tags("MangaList").
			Responses(goapidoc.NewResponse(200, "_Result<MangaGroupListDto>")),

		goapidoc.NewRoutePath("GET", "/v1/list/updated", "Get recent update mangas").
			Tags("MangaList").
			Params(param.ADPage, param.ADLimit).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaDto>>")),
	)
}

type MangaListController struct {
	config           *config.Config
	mangaListService *service.MangaListService
}

func NewMangaListController() *MangaListController {
	return &MangaListController{
		config:           xdi.GetByNameForce(sn.SConfig).(*config.Config),
		mangaListService: xdi.GetByNameForce(sn.SMangaListService).(*service.MangaListService),
	}
}

// GET /v1/list/serial
func (m *MangaListController) GetHotSerialMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetHotSerialMangas()
	if err != nil {
		return result.Error(exception.GetHotSerialMangasError).SetError(err, c)
	}

	res := dto.BuildMangaGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/finish
func (m *MangaListController) GetFinishedMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetFinishedMangas()
	if err != nil {
		return result.Error(exception.GetFinishedMangasError).SetError(err, c)
	}

	res := dto.BuildMangaGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/latest
func (m *MangaListController) GetLatestMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetLatestMangas()
	if err != nil {
		return result.Error(exception.GetLatestMangasError).SetError(err, c)
	}

	res := dto.BuildMangaGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/updated
func (m *MangaListController) GetRecentUpdatedMangas(c *gin.Context) *result.Result {
	pa := param.BindPage(c, m.config)
	pages, tot, err := m.mangaListService.GetRecentUpdatedMangas(pa) // categoryService.GetGenreMangas
	if err != nil {
		return result.Error(exception.GetUpdatedMangasError).SetError(err, c)
	}

	res := dto.BuildTinyMangaDtos(pages)
	return result.Ok().SetPage(pa.Page, pa.Limit, tot, res)
}
