package middleware

import (
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := xmodule.MustGetByName(sn.SLogger).(*logrus.Logger)

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
