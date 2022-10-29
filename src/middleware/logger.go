package middleware

import (
	"github.com/Aoi-hosizora/ahlib-web/xgin"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := xdi.GetByNameForce(sn.SLogger).(*logrus.Logger)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		rid := c.Writer.Header().Get("X-Request-Id")
		xgin.WithLogrus(logger, start, c,
			xgin.WithExtraString(rid),
			xgin.WithExtraFields(map[string]interface{}{"requestID": rid}),
		)
	}
}
