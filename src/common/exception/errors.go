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
	GetAllMangaPagesError     = New(500, se(), "failed to get all manga pages")
	GetMangaPageError         = New(500, se(), "failed to get manga page")
	MangaPageNotFoundError    = New(404, ce(), "manga page not found")
	GetMangaChapterError      = New(500, se(), "failed to get manga chapter")
	MangaChapterNotFoundError = New(404, ce(), "manga chapter not found")
	GetHotSerialMangasError   = New(500, se(), "failed to get hot serial mangas")
	GetFinishedMangasError    = New(500, se(), "failed to get finished mangas")
	GetLatestMangasError      = New(500, se(), "failed to get latest mangas")
	GetUpdatedMangasError     = New(500, se(), "failed to get updated mangas")
)

// category
var (
	GetGenresError      = New(500, se(), "failed to get genres")
	GetZonesError       = New(500, se(), "failed to get zones")
	GetAgesError        = New(500, se(), "failed to get ages")
	GetGenreMangasError = New(500, se(), "failed to get genre mangas")
	GenreNotFoundError  = New(404, ce(), "genre not found")
)

// search
var (
	SearchMangasError   = New(500, se(), "failed to search mangas")
	SearchNotFoundError = New(404, ce(), "search result not found")
)

func WrapValidationError(err error) *Error {
	if xvalidator.ValidationRequiredError(err) {
		return RequestParamError
	}
	return RequestFormatError
}
