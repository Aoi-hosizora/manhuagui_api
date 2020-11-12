package param

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/config"
	"github.com/gin-gonic/gin"
)

var (
	ADPage  = goapidoc.NewQueryParam("page", "integer#int32", false, "current page")
	ADLimit = goapidoc.NewQueryParam("limit", "integer#int32", false, "page size")
)

type PageParam struct {
	Page  int32
	Limit int32
}

// Bind ?page&limit
func BindPage(c *gin.Context, config *config.Config) *PageParam {
	page, err := xnumber.Atoi32(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := xnumber.Atoi32(c.DefaultQuery("limit", "0"))
	if def := config.Meta.DefLimit; err != nil || limit <= 0 {
		limit = def
	} else if max := config.Meta.MaxLimit; limit > max {
		limit = max
	}

	return &PageParam{Page: page, Limit: limit}
}

// Bind :xid
func BindRouteId(c *gin.Context, field string) (uint64, error) {
	uid, err := xnumber.Atou64(c.Param(field))
	if err != nil {
		return 0, err
	}
	if uid <= 0 {
		return 0, fmt.Errorf("id shoule larger than 0")
	}

	return uid, nil
}
