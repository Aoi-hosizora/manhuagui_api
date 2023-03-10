package errno

type Error struct {
	Status  int32
	Code    int32
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func New(status int32, code int32, message string) *Error {
	return &Error{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
