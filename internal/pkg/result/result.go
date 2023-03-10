package result

import (
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("Result", "global response").
			Properties(
				goapidoc.NewProperty("code", "integer#int32", true, "status code"),
				goapidoc.NewProperty("message", "string", true, "status message"),
			),

		goapidoc.NewDefinition("_Result", "global response").
			Generics("T").
			Properties(
				goapidoc.NewProperty("code", "integer#int32", true, "status code"),
				goapidoc.NewProperty("message", "string", true, "status message"),
				goapidoc.NewProperty("data", "T", true, "response data"),
			),
	)
}

type Result struct {
	status int32

	Code    int32           `json:"code"`
	Message string          `json:"message"`
	Data    any             `json:"data,omitempty"`
	Error   *errno.ErrorDto `json:"error,omitempty"`
}

func Status(status int32) *Result {
	message := http.StatusText(int(status))
	if status == 200 {
		message = "success"
	} else if message == "" {
		message = "unknown"
	}
	message = strings.ToLower(message)
	return &Result{status: status, Code: status, Message: message}
}

func Ok() *Result {
	return Status(http.StatusOK)
}

func Error(e *errno.Error) *Result {
	return Status(e.Status).SetCode(e.Code).SetMessage(e.Message)
}

func BindingError(err error, ctx *gin.Context) *Result {
	translated, need4xx := xgin.TranslateBindingError(err, xgin.WithUtTranslator(xgin.GetGlobalTranslator()))
	if need4xx {
		return Error(errno.RequestParamError).SetData(translated).SetError(err, nil) // no request info
	}
	return Error(errno.ServerUnknownError).SetError(err, ctx) // include request info
}

func (r *Result) SetStatus(status int32) *Result {
	r.status = status
	return r
}

func (r *Result) SetCode(code int32) *Result {
	r.Code = code
	return r
}

func (r *Result) SetMessage(message string) *Result {
	r.Message = strings.ToLower(message)
	return r
}

func (r *Result) SetData(data any) *Result {
	r.Data = data
	return r
}

func (r *Result) SetPage(page int32, limit int32, total int32, data any) *Result {
	r.Data = NewPage(page, limit, total, data)
	return r
}

func (r *Result) SetError(err error, ctx *gin.Context) *Result {
	if err != nil {
		r.Error = errno.BuildBasicErrorDto(err, ctx) // include request info, exclude trace info
	}
	return r
}

func (r *Result) prehandle() {
	if !config.IsDebugMode() && r.Error != nil {
		r.Error = r.Error.RequestOnly()
	}
}

func (r *Result) JSON(ctx *gin.Context) {
	r.prehandle()
	ctx.JSON(int(r.status), r) // application/json; charset=utf-8
}

func (r *Result) XML(ctx *gin.Context) {
	r.prehandle()
	ctx.XML(int(r.status), r) // application/xml; charset=utf-8
}
