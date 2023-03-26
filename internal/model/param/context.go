package param

import (
	"errors"
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/gin-gonic/gin"
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
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Server
	if err != nil || limit <= 0 {
		limit = int32(cfg.DefLimit)
	} else if limit > int32(cfg.MaxLimit) {
		limit = int32(cfg.MaxLimit)
	}

	return &PageParam{Page: page, Limit: limit}
}

// Bind Authorization / ?token
func BindToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}
	return token
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
