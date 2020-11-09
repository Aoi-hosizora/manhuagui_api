package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/dto"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/param"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddRoutePaths(
		goapidoc.NewRoutePath("GET", "/v1/manga/{id}", "Get manga page").
			Tags("Manga").
			Params(
				goapidoc.NewPathParam("id", "integer#int64", true, "manga id"),
			).
			Responses(goapidoc.NewResponse(200, "MangaPageDto")),

		goapidoc.NewRoutePath("GET", "/v1/manga/{mid}/{cid}", "Get manga chapter").
			Tags("Manga").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
			).
			Responses(goapidoc.NewResponse(200, "MangaChapterDto")),
	)
}

type MangaController struct {
	mangaService *service.MangaService
}

func NewMangaController() *MangaController {
	return &MangaController{
		mangaService: xdi.GetByNameForce(sn.SMangaService).(*service.MangaService),
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
	}

	res := dto.BuildMangaChapterDto(chapter)
	return result.Ok().SetData(res)
}
