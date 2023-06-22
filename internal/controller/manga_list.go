package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/apidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewGetOperation("/v1/list/serial", "Get hot serial mangas").
			Tags("MangaList").
			Responses(goapidoc.NewResponse(200, "_Result<MangaGroupListDto>")),

		goapidoc.NewGetOperation("/v1/list/finish", "Get finished mangas").
			Tags("MangaList").
			Responses(goapidoc.NewResponse(200, "_Result<MangaGroupListDto>")),

		goapidoc.NewGetOperation("/v1/list/latest", "Get latest mangas").
			Tags("MangaList").
			Responses(goapidoc.NewResponse(200, "_Result<MangaGroupListDto>")),

		goapidoc.NewGetOperation("/v1/list/homepage", "Get homepage mangas").
			Tags("MangaList").
			Responses(goapidoc.NewResponse(200, "_Result<HomepageMangaGroupListDto>")),

		goapidoc.NewGetOperation("/v1/list/updated", "Get recent update mangas").
			Tags("MangaList").
			Params(apidoc.ParamPage, apidoc.ParamLimit).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaDto>>")),

		goapidoc.NewGetOperation("/v1/list/updated_v2", "Get recent update mangas (new version)").
			Tags("MangaList").
			Params(apidoc.ParamPage, goapidoc.NewQueryParam("need_total", "boolean", false, "query total whether needed").Default(true)).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallerMangaDto>>")),

		goapidoc.NewGetOperation("/v1/list/published_v2", "Get recent published mangas (new version)").
			Tags("MangaList").
			Params(apidoc.ParamPage, goapidoc.NewQueryParam("need_total", "boolean", false, "query total whether needed").Default(true)).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallerMangaDto>>")),
	)
}

type MangaListController struct {
	mangaListService *service.MangaListService
}

func NewMangaListController() *MangaListController {
	return &MangaListController{
		mangaListService: xmodule.MustGetByName(sn.SMangaListService).(*service.MangaListService),
	}
}

// GET /v1/list/serial
func (m *MangaListController) GetHotSerialMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetHotSerialMangas()
	if err != nil {
		return result.Error(errno.GetHotSerialMangasError).SetError(err, c)
	}

	res := dto.BuildMangaGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/finish
func (m *MangaListController) GetFinishedMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetFinishedMangas()
	if err != nil {
		return result.Error(errno.GetFinishedMangasError).SetError(err, c)
	}

	res := dto.BuildMangaGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/latest
func (m *MangaListController) GetLatestMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetLatestMangas()
	if err != nil {
		return result.Error(errno.GetLatestMangasError).SetError(err, c)
	}

	res := dto.BuildMangaGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/homepage
func (m *MangaListController) GetHomepageMangas(c *gin.Context) *result.Result {
	list, err := m.mangaListService.GetHomepageMangaGroupList()
	if err != nil {
		return result.Error(errno.GetHomepageMangasError).SetError(err, c)
	}

	res := dto.BuildHomepageMangaGroupListDto(list)
	return result.Ok().SetData(res)
}

// GET /v1/list/updated
func (m *MangaListController) GetRecentUpdatedMangas(c *gin.Context) *result.Result {
	pa := param.BindQueryPage(c)
	pages, tot, err := m.mangaListService.GetRecentUpdatedMangas(pa)
	if err != nil {
		return result.Error(errno.GetUpdatedMangasError).SetError(err, c)
	}

	res := dto.BuildTinyMangaDtos(pages)
	return result.Ok().SetPage(pa.Page, pa.Limit, tot, res)
}

// GET /v1/list/updated_v2
func (m *MangaListController) GetRecentUpdatedMangasV2(c *gin.Context) *result.Result {
	pa := param.BindQueryPage(c)
	needTotalVar := strings.ToLower(strings.TrimSpace(c.Query("need_total")))
	needTotal := needTotalVar == "1" || needTotalVar == "t" || needTotalVar == "true"
	mangas, tot, err := m.mangaListService.GetRecentUpdatedMangasFromMobile(pa, needTotal)
	if err != nil {
		return result.Error(errno.GetUpdatedMangasError).SetError(err, c)
	}

	res := dto.BuildSmallerMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, int32(len(mangas)), tot, res)
}

// GET /v1/list/published_v2
func (m *MangaListController) GetRecentPublishedMangasV2(c *gin.Context) *result.Result {
	pa := param.BindQueryPage(c)
	needTotalVar := strings.ToLower(strings.TrimSpace(c.Query("need_total")))
	needTotal := needTotalVar == "1" || needTotalVar == "t" || needTotalVar == "true"

	mangas, total, err := m.mangaListService.GetRecentPublishedMangasFromMobile(pa, needTotal)
	if err != nil {
		return result.Error(errno.GetAllMangasError).SetError(err, c)
	}

	res := dto.BuildSmallerMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, int32(len(mangas)), total, res)
}
