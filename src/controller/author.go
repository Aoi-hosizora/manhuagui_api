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
		goapidoc.NewRoutePath("GET", "/v1/author", "Get all authors").
			Tags("Author").
			Params(param.ADPage, param.ADOrder).
			Responses(goapidoc.NewResponse(200, "_Result<_Page<SmallAuthorDto>>")),
	)
}

type AuthorController struct {
	config        *config.Config
	authorService *service.AuthorService
}

func NewAuthorController() *AuthorController {
	return &AuthorController{
		config:        xdi.GetByNameForce(sn.SConfig).(*config.Config),
		authorService: xdi.GetByNameForce(sn.SAuthorService).(*service.AuthorService),
	}
}

// GET /v1/author
func (a *AuthorController) GetAllAuthors(c *gin.Context) *result.Result {
	pa := param.BindPageOrder(c, a.config)

	authors, limit, total, err := a.authorService.GetAllAuthors(pa.Page, pa.Order == "popular")
	if err != nil {
		return result.Error(exception.GetAllAuthorsError).SetError(err, c)
	}

	res := dto.BuildSmallAuthorDtos(authors)
	return result.Ok().SetPage(pa.Page, limit, total, res)
}
