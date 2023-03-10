package middleware

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"reflect"
	"time"
)

func LimiterMiddleware() gin.HandlerFunc {
	// TODO ip limiter
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Meta
	interval := time.Second * time.Duration(cfg.BucketPrd)
	limiter := ratelimit.NewBucketWithQuantum(interval, cfg.BucketCap, cfg.BucketQua)
	startTime := xreflect.GetUnexportedField(reflect.ValueOf(limiter).Elem().FieldByName("startTime")).(time.Time)
	return func(c *gin.Context) {
		available := xnumber.I64toa(limiter.Available())
		capacity := xnumber.I64toa(limiter.Capacity())
		reset := (interval - (time.Now().Sub(startTime) % interval)).String()
		c.Header("X-RateLimit-Remaining", available)
		c.Header("X-RateLimit-Limit", capacity)
		c.Header("X-RateLimit-Reset", reset)

		if limiter.TakeAvailable(1) == 0 {
			r := gin.H{"remaining": available, "limit": capacity /* always 0 here */, "reset": reset}
			result.Status(http.StatusTooManyRequests).SetData(r).JSON(c)
			c.Abort()
		}
	}
}
