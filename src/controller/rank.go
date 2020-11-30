package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-backend/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/dto"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-backend/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddRoutePaths(
		goapidoc.NewRoutePath("GET", "/v1/rank/day", "Get day ranking").
			Tags("Rank").
			Params(goapidoc.NewQueryParam("type", "string", false, "manga genre / zone / age, empty or all if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/rank/week", "Get week ranking").
			Tags("Rank").
			Params(goapidoc.NewQueryParam("type", "string", false, "manga genre / zone / age, empty or all if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/rank/month", "Get month ranking").
			Tags("Rank").
			Params(goapidoc.NewQueryParam("type", "string", false, "manga genre / zone / age, empty or all if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/rank/total", "Get total ranking").
			Tags("Rank").
			Params(goapidoc.NewQueryParam("type", "string", false, "manga genre / zone / age, empty or all if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),
	)
}

type RankController struct {
	rankService *service.RankService
}

func NewRankController() *RankController {
	return &RankController{
		rankService: xdi.GetByNameForce(sn.SRankService).(*service.RankService),
	}
}

// GET /v1/rank/day
func (r *RankController) GetDayRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetDayRanking(typ)
	if err != nil {
		return result.Error(exception.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(exception.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}

// GET /v1/rank/week
func (r *RankController) GetWeekRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetWeekRanking(typ)
	if err != nil {
		return result.Error(exception.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(exception.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}

// GET /v1/rank/month
func (r *RankController) GetMonthRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetMonthRanking(typ)
	if err != nil {
		return result.Error(exception.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(exception.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}

// GET /v1/rank/total
func (r *RankController) GetTotalRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetTotalRanking(typ)
	if err != nil {
		return result.Error(exception.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(exception.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}
