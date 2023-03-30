package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewGetOperation("/v1/rank/day", "Get day ranking").
			Tags("Rank").
			Params(goapidoc.NewQueryParam("type", "string", false, "manga genre / zone / age, empty or all if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewGetOperation("/v1/rank/week", "Get week ranking").
			Tags("Rank").
			Params(goapidoc.NewQueryParam("type", "string", false, "manga genre / zone / age, empty or all if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewGetOperation("/v1/rank/month", "Get month ranking").
			Tags("Rank").
			Params(goapidoc.NewQueryParam("type", "string", false, "manga genre / zone / age, empty or all if all")).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MangaRankDto>>")),

		goapidoc.NewGetOperation("/v1/rank/total", "Get total ranking").
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
		rankService: xmodule.MustGetByName(sn.SRankService).(*service.RankService),
	}
}

// GET /v1/rank/day
func (r *RankController) GetDayRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetDayRanking(typ)
	if err != nil {
		return result.Error(errno.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(errno.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}

// GET /v1/rank/week
func (r *RankController) GetWeekRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetWeekRanking(typ)
	if err != nil {
		return result.Error(errno.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(errno.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}

// GET /v1/rank/month
func (r *RankController) GetMonthRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetMonthRanking(typ)
	if err != nil {
		return result.Error(errno.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(errno.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}

// GET /v1/rank/total
func (r *RankController) GetTotalRanking(c *gin.Context) *result.Result {
	typ := c.Query("type")
	ranking, err := r.rankService.GetTotalRanking(typ)
	if err != nil {
		return result.Error(errno.GetRankingError).SetError(err, c)
	} else if ranking == nil {
		return result.Error(errno.RankingTypeNotFoundError)
	}

	res := dto.BuildMangaRankDtos(ranking)
	return result.Ok().SetPage(1, int32(len(ranking)), int32(len(ranking)), res)
}
