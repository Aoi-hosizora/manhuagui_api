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
		goapidoc.NewOperation("GET", "/v1/comment/manga/{mid}", "Get manga comments").
			Tags("Comment").
			Params(goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"), apidoc.ParamPage).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<CommentDto>>")),

		goapidoc.NewOperation("POST", "/v1/comment/{cid}/like", "Like comment").
			Tags("Comment").
			Params(goapidoc.NewPathParam("cid", "integer#int64", true, "comment id")).
			Responses(goapidoc.NewResponse(200, "Result")),

		goapidoc.NewOperation("POST", "/v1/comment/manga/{mid}", "Add manga comment").
			Tags("Comment").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewQueryParam("text", "string", true, "comment content"),
				apidoc.ParamToken,
			).
			Responses(goapidoc.NewResponse(200, "_Result<AddedCommentDto>")),

		goapidoc.NewOperation("POST", "/v1/comment/manga/{mid}/{cid}", "Reply manga comment").
			Tags("Comment").
			Params(
				goapidoc.NewPathParam("mid", "integer#int64", true, "manga id"),
				goapidoc.NewPathParam("cid", "integer#int64", true, "comment id"),
				goapidoc.NewQueryParam("text", "string", true, "comment content"),
				apidoc.ParamToken,
			).
			Responses(goapidoc.NewResponse(200, "_Result<AddedCommentDto>")),
	)
}

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentService() *CommentController {
	return &CommentController{
		commentService: xmodule.MustGetByName(sn.SCommentService).(*service.CommentService),
	}
}

// GET /v1/comment/manga/:mid
func (co *CommentController) GetComments(c *gin.Context) *result.Result {
	pa := param.BindQueryPage(c)
	mid, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
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

// POST /v1/comment/:cid/like
func (co *CommentController) ListComment(c *gin.Context) *result.Result {
	cid, err := param.BindRouteID(c, "cid")
	if err != nil {
		return result.BindingError(err, c)
	}

	err = co.commentService.LikeComment(cid)
	if err != nil {
		return result.Error(errno.LikeCommentError).SetError(err, c)
	}

	return result.Ok()
}

// POST /v1/comment/manga/:mid
func (co *CommentController) AddComment(c *gin.Context) *result.Result {
	mid, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}
	text := strings.TrimSpace(c.Query("text"))
	if text == "" {
		return result.Error(errno.EmptyCommentError)
	}
	token := param.BindToken(c)

	comment, auth, err := co.commentService.AddComment(token, mid, text)
	if err != nil {
		return result.Error(errno.AddCommentError).SetError(err, c)
	}

	if !auth {
		return result.Error(errno.UnauthorizedError)
	}
	res := dto.BuildAddedCommentDto(comment)
	return result.Ok().SetData(res)
}

// POST /v1/comment/manga/:mid/:cid
func (co *CommentController) ReplyComment(c *gin.Context) *result.Result {
	mid, err := param.BindRouteID(c, "mid")
	if err != nil {
		return result.BindingError(err, c)
	}
	cid, err := param.BindRouteID(c, "cid")
	if err != nil {
		return result.BindingError(err, c)
	}
	text := strings.TrimSpace(c.Query("text"))
	if text == "" {
		return result.Error(errno.EmptyCommentError)
	}
	token := param.BindToken(c)

	comment, auth, err := co.commentService.ReplyComment(token, mid, cid, text)
	if err != nil {
		return result.Error(errno.ReplyCommentError).SetError(err, c)
	}

	if !auth {
		return result.Error(errno.UnauthorizedError)
	}
	res := dto.BuildAddedCommentDto(comment)
	return result.Ok().SetData(res)
}
