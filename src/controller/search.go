package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-api/src/config"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/vo"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-api/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddRoutePaths(
		goapidoc.NewRoutePath("GET", "/v1/search/{keyword}", "Search mangas").
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
		config:        xdi.GetByNameForce(sn.SConfig).(*config.Config),
		searchService: xdi.GetByNameForce(sn.SSearchService).(*service.SearchService),
	}
}

// GET /v1/search/:keyword
func (s *SearchController) SearchMangas(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, s.config)
	keyword := c.Param("keyword")

	mangas, limit, total, err := s.searchService.SearchMangas(keyword, pa.Page, pa.Order) // popular / new / update
	if err != nil {
		return result.Error(exception.SearchMangasError).SetError(err, c)
	} else if mangas == nil { // empty
		res := dto.BuildSmallMangaDtos([]*vo.SmallManga{})
		return result.Ok().SetPage(pa.Page, limit, 0, res)
	}

	res := dto.BuildSmallMangaDtos(mangas)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
