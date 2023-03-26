package middleware

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"github.com/Aoi-hosizora/ahlib/xconstant/headers"
	"github.com/Aoi-hosizora/ahlib/xgeneric/xsugar"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xnumber"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/gincache"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/iplimiter"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/ratelimit"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}

func getRequestID(c *gin.Context) string {
	return c.Writer.Header().Get(headers.XRequestID)
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := strings.TrimSpace(getRequestID(c))
		if rid == "" {
			rid = uuid.New().String()
			c.Header(headers.XRequestID, rid)
		}
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	logger := xmodule.MustGetByName(sn.SLogger).(*logrus.Logger)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()

		rid := getRequestID(c)
		options := []xgin.LoggerOption{xgin.WithExtraText(fmt.Sprintf(" | %s", rid)), xgin.WithExtraFieldsV("request_id", rid)}
		if ok, basedOn := isCached(c); ok {
			options = append(options, xgin.WithMoreExtraText(" | cached"), xgin.WithMoreExtraFieldsV("base_request_id", basedOn))
		}
		xgin.LogResponseToLogrus(logger, c, start, end, options...)
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	logger := xmodule.MustGetByName(sn.SLogger).(*logrus.Logger)
	const skip = 2
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errDto, stack := errno.BuildFullErrorDto(err, c, skip) // include request info and trace info
				xcolor.BrightRed.Printf("\n%s\n\n", stack.String())

				rid := getRequestID(c)
				options := []xgin.LoggerOption{xgin.WithExtraText(fmt.Sprintf(" | %s", rid)), xgin.WithExtraFieldsV("request_id", rid)}
				xgin.LogRecoveryToLogrus(logger, err, stack, options...)

				r := result.Error(errno.ServerUnknownError)
				r.Error = errDto
				r.JSON(c)
			}
		}()
		c.Next()
	}
}

func LimiterMiddleware() gin.HandlerFunc {
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Server
	fillInterval := time.Second * time.Duration(cfg.BucketPeriod)
	cleanupDuration := time.Second * time.Duration(cfg.BucketCleanup)
	ipLimiter := iplimiter.NewIPRateLimiter(cleanupDuration, int16(cfg.BucketSurvived), func() *ratelimit.Bucket {
		return ratelimit.NewBucketWithQuantum(fillInterval, int64(cfg.BucketCap), int64(cfg.BucketQua))
	})
	return func(c *gin.Context) {
		limiter, startTime, takeAvailable := ipLimiter.GetOrCreate(c.ClientIP())
		available := xnumber.I64toa(limiter.Available())
		capacity := xnumber.I64toa(limiter.Capacity())
		reset := (fillInterval - (time.Now().Sub(startTime) % fillInterval)).String()
		c.Header(headers.XRateLimitRemaining, available)
		c.Header(headers.XRateLimitLimit, capacity)
		c.Header(headers.XRateLimitReset, reset)

		if takeAvailable(1) == 0 {
			r := gin.H{"remaining": available, "limit": capacity /* always 0 here */, "reset": reset}
			result.Status(http.StatusTooManyRequests).SetData(r).JSON(c)
			c.Abort()
		}
	}
}

func isCached(c *gin.Context) (bool, string) {
	baseRid := c.Writer.Header().Get("X-Base-Request-ID")
	if baseRid == "" {
		return false, ""
	}
	return true, baseRid
}

func CacheMiddleware() gin.HandlerFunc {
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Server
	cacheExpiration := time.Second * time.Duration(cfg.CacheExpire)
	storage := gincache.NewCacheStorage(int16(cfg.CacheSize), cacheExpiration)
	return func(c *gin.Context) {
		forceRefresh := xsugar.Let(strings.ToLower(c.Query("force_refresh")), func(t string) bool { return t == "true" || t == "t" || t == "1" })
		if cfg.DisableCache || c.Request.Method != "GET" || forceRefresh || strings.Contains(c.FullPath(), "swagger") {
			return
		}

		key := fmt.Sprintf("%s-%s", c.Request.URL.RequestURI(), c.GetHeader(headers.Authorization))
		cached, expiration, _ := storage.GetWithExpiration(key)
		if cached == nil || (expiration == time.Time{}) || !gincache.Is2XXStatus(cached.Status()) {
			cw := gincache.NewCachedWriter(c.Writer, storage, key)
			c.Writer = cw
			c.Next()
			c.Writer = cw.Origin()
			if gincache.Is2XXStatus(cw.Status()) {
				_ = cw.WriteCache() // ignore error
			}
		} else {
			for k, vs := range cached.Header() {
				if c.Writer.Header().Get(k) == "" { // header existed => ignore cached header
					for _, v := range vs {
						c.Writer.Header().Add(k, v)
					}
				}
			}
			c.Writer.Header().Set(headers.CacheControl, fmt.Sprintf("private, max-age=%d%s", cacheExpiration/time.Second, xsugar.If(cfg.ClientCache, "", ", no-store")))
			c.Writer.Header().Set(headers.Expires, expiration.UTC().Format(http.TimeFormat))
			c.Writer.Header().Set("X-Base-Request-ID", cached.Header().Get(headers.XRequestID))
			c.Writer.Flush()
			c.Writer.WriteHeader(cached.Status())
			_, _ = c.Writer.Write(cached.Data())
			c.Abort()
		}
	}
}
