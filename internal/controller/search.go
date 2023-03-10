package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/object"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewOperation("GET", "/v1/search/{keyword}", "Search mangas").
			Desc("order by popular / new / update").
			Tags("Search").
			Params(
				goapidoc.NewPathParam("keyword", "string", true, "search keyword"),
				param.ParamPage, param.ParamOrder,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallMangaDto>>")),
	)
}

type SearchController struct {
	config        *config.Config
	searchService *service.SearchService
}

func NewSearchController() *SearchController {
	return &SearchController{
		config:        xmodule.MustGetByName(sn.SConfig).(*config.Config),
		searchService: xmodule.MustGetByName(sn.SSearchService).(*service.SearchService),
	}
}

// GET /v1/search/:keyword
func (s *SearchController) SearchMangas(c *gin.Context) *result.Result {
	pa := param.BindQueryPageOrder(c)
	keyword := c.Param("keyword")

	mangas, limit, total, err := s.searchService.SearchMangas(keyword, pa.Page, pa.Order) // popular / new / update
	if err != nil {
		return result.Error(errno.SearchMangasError).SetError(err, c)
	} else if mangas == nil { // empty
		res := dto.BuildSmallMangaDtos([]*object.SmallManga{})
		return result.Ok().SetPage(pa.Page, limit, 0, res)
	}

	res := dto.BuildSmallMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
