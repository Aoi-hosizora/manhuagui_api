package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddRoutePaths(
		goapidoc.NewRoutePath("GET", "/v1/rank/day", "Get day ranking").
			Tags("Rank").
			Params(goapidoc.NewPathParam("type", "string", true, "manga genre / zone / age, empty if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/rank/week", "Get week ranking").
			Tags("Rank").
			Params(goapidoc.NewPathParam("type", "string", true, "manga genre / zone / age, empty if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/rank/month", "Get month ranking").
			Tags("Rank").
			Params(goapidoc.NewPathParam("type", "string", true, "manga genre / zone / age, empty if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/rank/total", "Get total ranking").
			Tags("Rank").
			Params(goapidoc.NewPathParam("type", "string", true, "manga genre / zone / age, empty if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),
	)
}

type RankController struct {
	config      *config.Config
	rankService *service.RankService
}

func NewRankController() *RankController {
	return &RankController{
		config:      xdi.GetByNameForce(sn.SConfig).(*config.Config),
		rankService: xdi.GetByNameForce(sn.SRankService).(*service.RankService),
	}
}

// GET /v1/rank/day
func (r *RankController) GetDayRanking(c *gin.Context) *result.Result {
	return result.Ok()
}

// GET /v1/rank/week
func (r *RankController) GetWeekRanking(c *gin.Context) *result.Result {
	return result.Ok()
}

// GET /v1/rank/month
func (r *RankController) GetMonthRanking(c *gin.Context) *result.Result {
	return result.Ok()
}

// GET /v1/rank/total
func (r *RankController) GetTotalRanking(c *gin.Context) *result.Result {
	return result.Ok()
}
