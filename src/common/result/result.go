package result

import (
	"github.com/Aoi-hosizora/ahlib-web/xdto"
	"github.com/Aoi-hosizora/ahlib-web/xgin"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/exception"
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
	Status  int32          `json:"-"`
	Code    int32          `json:"code"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *xdto.ErrorDto `json:"error,omitempty"`
}

func Status(status int32) *Result {
	message := http.StatusText(int(status))
	if status == 200 {
		message = "success"
	} else if message == "" {
		message = "unknown"
	}
	return &Result{
		Status:  status,
		Code:    status,
		Message: strings.ToLower(message),
	}
}

func Ok() *Result {
	return Status(http.StatusOK)
}

func Error(e *exception.Error) *Result {
	return Status(e.Status).SetCode(e.Code).SetMessage(e.Message)
}

func (r *Result) SetStatus(status int32) *Result {
	r.Status = status
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

func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}

func (r *Result) SetPage(page int32, limit int32, total int32, data interface{}) *Result {
	r.Data = NewPage(page, limit, total, data)
	return r
}

func (r *Result) SetError(err error, c *gin.Context) *Result {
	rid := c.Writer.Header().Get("X-Request-Id")
	r.Error = xgin.BuildBasicErrorDto(err, c, "request_id", rid)
	return r
}

func (r *Result) JSON(c *gin.Context) {
	if gin.Mode() != gin.DebugMode {
		r.Error = nil
	}
	c.JSON(int(r.Status), r)
}

func (r *Result) XML(c *gin.Context) {
	if gin.Mode() != gin.DebugMode {
		r.Error = nil
	}
	c.XML(int(r.Status), r)
}
