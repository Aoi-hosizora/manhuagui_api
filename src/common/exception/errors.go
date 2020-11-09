package exception

import "github.com/Aoi-hosizora/ahlib-web/xvalidator"

var (
	cerr = int32(40000) // client error code
	serr = int32(50000) // server error code
)

// Return ++ cerr.
func ce() int32 { cerr++; return cerr }

// Return ++ serr.
func se() int32 { serr++; return serr }

var (
	RequestParamError   = New(400, cerr, "request param error")
	RequestFormatError  = New(400, ce(), "request format error")
	ServerRecoveryError = New(500, serr, "server unknown error")
)

// manga
var (
	GetMangaPageError    = New(500, se(), "failed to get manga page")
	GetMangaChapterError = New(500, se(), "failed to get manga chapter")
)

func WrapValidationError(err error) *Error {
	if xvalidator.ValidationRequiredError(err) {
		return RequestParamError
	}
	return RequestFormatError
}
