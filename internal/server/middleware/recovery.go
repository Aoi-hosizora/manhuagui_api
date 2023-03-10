package middleware

import (
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	logger := xmodule.MustGetByName(sn.SLogger).(*logrus.Logger)
	skip := 2

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				rid := c.Writer.Header().Get("X-Request-Id")
				r := result.Error(errno.ServerRecoveryError)
				r.Error = xgin.BuildErrorDto(err, c, skip, true, "request_id", rid)

				// TODO
				xrecovery.WithLogrus(logger, err,
					xrecovery.WithExtraString(rid),
					xrecovery.WithExtraFields(map[string]interface{}{"requestID": rid}),
				)
				r.JSON(c)
			}
		}()
		c.Next()
	}
}
