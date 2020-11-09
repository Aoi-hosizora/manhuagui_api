package exception

var (
	RequestParamError   = New(400, 40000, "request param error")
	ServerRecoveryError = New(500, 50000, "server unknown error")
	DataNotFoundError   = New(404, 40001, "data not found")
)
