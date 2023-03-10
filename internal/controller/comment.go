package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/param"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewOperation("GET", "/v1/comment/manga/{mid}", "Get manga comments").
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
		config:         xmodule.MustGetByName(sn.SConfig).(*config.Config),
		commentService: xmodule.MustGetByName(sn.SCommentService).(*service.CommentService),
	}
}

// /v1/comment/manga/:mid
func (co *CommentController) GetComments(c *gin.Context) *result.Result {
	pa := param.BindPage(c, co.config)
	mid, err := param.BindRouteId(c, "mid")
	if err != nil {
		return result.Error(errno.RequestParamError).SetError(err, c)
	}

	comments, total, err := co.commentService.GetComments(mid, pa.Page)
	if err != nil {
		return result.Error(errno.GetMangaCommentsError).SetError(err, c)
	} else if comments == nil {
		return result.Error(errno.MangaNotFoundError) // unreachable
	}

	res := dto.BuildCommentDtos(comments)
	return result.Ok().SetPage(pa.Page, int32(len(comments)), total, res)
}
