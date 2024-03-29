package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/apidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewGetOperation("/v1/manga", "Get all mangas").
			Desc("order by popular / new / update").
			Tags("Manga").
			Params(apidoc.ParamPage, apidoc.ParamOrder).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<TinyMangaDto>>")),

		goapidoc.NewGetOperation("/v1/manga/{mid}", "Get manga").
			Tags("Manga").
			Params(goapidoc.NewPathParam("mid", "integer#int64", true, "manga id")).
			Responses(goapidoc.NewResponse(200, "_Result<MangaDto>")),

		goapidoc.NewGetOperation("/v1/manga/random", "Get random manga").
			Tags("Manga").
			Responses(goapidoc.NewResponse(200, "_Result<RandomMangaInfoDto>")),

		goapidoc.NewPostOperation("/v1/manga/{mid}/vote", "Vote manga").
			Tags("Manga").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewQueryParam("score", "integer#int32", true, "manga score").ValueRange(1, 5),
				apidoc.ParamToken,
			).
			Responses(goapidoc.NewResponse(200, "_Result<RandomMangaInfoDto>")),

		goapidoc.NewGetOperation("/v1/manga/{mid}/{cid}", "Get manga chapter").
			Tags("Manga").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "manga chapter id"),
			).
			Responses(goapidoc.NewResponse(200, "_Result<MangaChapterDto>")),
	)
}

type MangaController struct {
	mangaService    *service.MangaService
	categoryService *service.CategoryService
}

func NewMangaController() *MangaController {
	return &MangaController{
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

// GET /v1/manga/:mid
func (m *MangaController) GetManga(c *gin.Context) *result.Result {
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
func (m *MangaController) GetRandomManga(c *gin.Context) *result.Result {
	info, err := m.mangaService.GetRandomMangaInfo()
	if err != nil || info == nil {
		return result.Error(errno.GetRandomMangaError).SetError(err, c)
	}

	res := dto.BuildRandomMangaInfoDto(info)
	return result.Ok().SetData(res)
}

// POST /v1/manga/:mid/vote
func (m *MangaController) VoteManga(c *gin.Context) *result.Result {
	id, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}
	token := param.BindToken(c)
	if token == "" {
		return result.Error(errno.UnauthorizedError)
	}
	score := c.Query("score")
	scoreValue, err := xnumber.Atou8(score)
	if err != nil {
		return result.BindingError(err, c)
	}
	if scoreValue < 1 {
		scoreValue = 1
	}
	if scoreValue > 5 {
		scoreValue = 5
	}

	err = m.mangaService.VoteManga(token, id, scoreValue)
	if err != nil {
		return result.Error(errno.VoteMangaError).SetError(err, c)
	}

	return result.Ok()
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
