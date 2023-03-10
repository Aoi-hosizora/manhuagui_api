package param

import (
	"errors"
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/gin-gonic/gin"
)

var (
	ParamPage  = goapidoc.NewQueryParam("page", "integer#int32", false, "current page")
	ParamLimit = goapidoc.NewQueryParam("limit", "integer#int32", false, "page size")
	ParamOrder = goapidoc.NewQueryParam("order", "string", false, "order string")
)

type PageParam struct {
	Page  int32
	Limit int32
}

type PageOrderParam struct {
	Page  int32  `json:"page"`
	Limit int32  `json:"limit"`
	Order string `json:"order"`
}

// Bind ?page&limit
func BindQueryPage(c *gin.Context) *PageParam {
	page, err := xnumber.Atoi32(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := xnumber.Atoi32(c.Query("limit"))
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Meta
	if err != nil || limit <= 0 {
		limit = cfg.DefLimit
	} else if limit > cfg.MaxLimit {
		limit = cfg.MaxLimit
	}

	return &PageParam{Page: page, Limit: limit}
}

// Bind ?page&limit&order
func BindQueryPageOrder(c *gin.Context) *PageOrderParam {
	page := BindQueryPage(c)
	order := c.DefaultQuery("order", "")
	return &PageOrderParam{Page: page.Page, Limit: page.Limit, Order: order}
}

// Bind :xid
func BindRouteID(c *gin.Context, name string) (uint64, error) {
	s := c.Param(name)
	id, err := xnumber.Atou64(s)
	if err != nil {
		return 0, xgin.NewRouterDecodeError(name, s, err, "")
	}
	if id <= 0 {
		err = errors.New("non-positive number")
		return 0, xgin.NewRouterDecodeError(name, s, err, "must be a positive number")
	}
	return id, nil
}

// Bind body
func BindBody[T any](c *gin.Context, obj T) (T, error) {
	err := c.ShouldBind(obj)
	if err != nil {
		var zero T
		return zero, err
	}
	return obj, nil
}
