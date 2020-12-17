package literr

var errCode map[int]string

const (
	ErrCodeBadRequest           = 40000
	ErrCodeResourceNotFound     = 40001
	ErrCodeInvalidRequestParams = 40002
	ErrCodeInternalError        = 50000
)

func initErrCodeMapping() {
	errCode = map[int]string{
		ErrCodeBadRequest:           "Bad request",
		ErrCodeResourceNotFound:     "Resource not found",
		ErrCodeInvalidRequestParams: "Invalid request params",
		ErrCodeInternalError:        "Internal error",
	}
}

func init() {
	initErrCodeMapping()
}

func trans(code int) string {
	return errCode[code]
}
