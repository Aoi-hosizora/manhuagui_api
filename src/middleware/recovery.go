package middleware

import (
	"github.com/Aoi-hosizora/ahlib-web/xgin"
	"github.com/Aoi-hosizora/ahlib-web/xrecovery"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	logger := xdi.GetByNameForce(sn.SLogger).(*logrus.Logger)
	skip := 2

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				rid := c.Writer.Header().Get("X-Request-Id")
				r := result.Error(exception.ServerRecoveryError)
				r.Error = xgin.BuildErrorDto(err, c, skip, true, "request_id", rid)

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
