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
		goapidoc.NewRoutePath("GET", "/v1/comment/manga/{mid}", "Get manga comments").
			Tags("Comment").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				param.ParamPage,
			).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CommentDto>>")),
	)
}

type CommentController struct {
	config         *config.Config
	commentService *service.CommentService
}

func NewCommentService() *CommentController {
	return &CommentController{
		config:         xdi.GetByNameForce(sn.SConfig).(*config.Config),
		commentService: xdi.GetByNameForce(sn.SCommentService).(*service.CommentService),
	}
}

// /v1/comment/manga/:mid
func (co *CommentController) GetComments(c *gin.Context) *result.Result {
	pa := param.BindPage(c, co.config)
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(exception.RequestParamError).SetError(err, c)
	}

	comments, total, err := co.commentService.GetComments(mid, pa.Page)
	if err != nil {
		return result.Error(exception.GetMangaCommentsError).SetError(err, c)
	} else if comments == nil {
		return result.Error(exception.MangaNotFoundError) // unreachable
	}

	res := dto.BuildCommentDtos(comments)
	return result.Ok().SetPage(pa.Page, int32(len(comments)), total, res)
}
