package exception

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
	ServerRecoveryError = New(500, serr, "server unknown error")
)

// manga
var (
	GetAllMangasError         = New(500, se(), "failed to get all mangas")
	GetMangaError             = New(500, se(), "failed to get manga")
	MangaNotFoundError        = New(404, ce(), "manga not found")
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
	SearchMangasError = New(500, se(), "failed to search mangas")
)

// author
var (
	GetAllAuthorsError   = New(500, se(), "failed to get all authors")
	GetAuthorError       = New(500, se(), "failed to get author")
	AuthorNotFound       = New(404, ce(), "author not found")
	GetAuthorMangasError = New(500, se(), "failed to get author mangas")
)

// rank
var (
	GetRankingError          = New(500, se(), "failed to get ranking list")
	RankingTypeNotFoundError = New(404, ce(), "ranking type not found")
)

// comment
var (
	GetMangaCommentsError = New(500, se(), "failed to get manga comments")
)

// user
var (
	LoginError                = New(500, se(), "failed to login")
	PasswordError             = New(401, ce(), "username or password error")
	CheckLoginError           = New(500, se(), "failed to check login")
	UnauthorizedError         = New(401, ce(), "Unauthorized")
	GetUserError              = New(500, se(), "failed to get user")
	GetShelfMangasError       = New(500, se(), "failed to get shelf mangas")
	CheckMangaShelfError      = New(500, se(), "failed to check manga in shelf")
	SaveMangaToShelfError     = New(500, se(), "failed to save manga to shelf")
	RemoveMangaFromShelfError = New(500, se(), "failed to remove manga from shelf")
	MangaAlreadyInShelfError  = New(409, ce(), "manga is already in shelf")
	MangaNotInShelfYetError   = New(404, ce(), "manga is not in shelf yet")
)
