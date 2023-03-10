package result

import (
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xruntime"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteRegisterer struct {
	route gin.IRouter
}

type (
	GinHandlerFunc = gin.HandlerFunc
	HandlerFunc    = func(c *gin.Context) *Result
)

func NewRouteRegisterer(route gin.IRouter) *RouteRegisterer {
	return &RouteRegisterer{
		route: route,
	}
}

func (r *RouteRegisterer) Route() gin.IRouter {
	return r.route
}

func (r *RouteRegisterer) Group(relativePath string, handlers ...GinHandlerFunc) *RouteRegisterer {
	group := r.route.Group(relativePath, handlers...)
	return NewRouteRegisterer(group)
}

func (r *RouteRegisterer) Use(middleware ...GinHandlerFunc) *RouteRegisterer {
	r.route.Use(middleware...)
	return r
}

func (r *RouteRegisterer) Handle(httpMethod, relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	// !!!
	hackedMethod := xruntime.HackHideStringAfterString(httpMethod, xruntime.NameOfFunction(handler))
	handlers := append(preHandlers, func(c *gin.Context) {
		result := handler(c)
		if result != nil {
			result.JSON(c)
			c.Abort()
		}
	})
	r.route.Handle(hackedMethod, relativePath, handlers...)
	return r
}

func (r *RouteRegisterer) GET(relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	return r.Handle(http.MethodGet, relativePath, handler, preHandlers...)
}

func (r *RouteRegisterer) POST(relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	return r.Handle(http.MethodPost, relativePath, handler, preHandlers...)
}

func (r *RouteRegisterer) DELETE(relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	return r.Handle(http.MethodDelete, relativePath, handler, preHandlers...)
}

func (r *RouteRegisterer) PATCH(relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	return r.Handle(http.MethodPatch, relativePath, handler, preHandlers...)
}

func (r *RouteRegisterer) PUT(relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	return r.Handle(http.MethodPut, relativePath, handler, preHandlers...)
}

func (r *RouteRegisterer) OPTIONS(relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	return r.Handle(http.MethodOptions, relativePath, handler, preHandlers...)
}

func (r *RouteRegisterer) HEAD(relativePath string, handler HandlerFunc, preHandlers ...GinHandlerFunc) *RouteRegisterer {
	return r.Handle(http.MethodHead, relativePath, handler, preHandlers...)
}

func PrintRouteFunc(printRouteFunc xgin.DebugPrintRouteFuncType) xgin.DebugPrintRouteFuncType {
	if printRouteFunc == nil {
		printRouteFunc = xgin.DefaultPrintRouteFunc
	}
	return func(httpMethod, absolutePath, handlerName string, numHandlers int) {
		realHandlerName := xruntime.HackGetHiddenStringAfterString(httpMethod)
		if realHandlerName == "" {
			realHandlerName = handlerName
		}
		printRouteFunc(httpMethod, absolutePath, realHandlerName, numHandlers)
	}
}
