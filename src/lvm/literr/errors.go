package literr

import "strconv"

var errCode map[Code]string

type Code int

func (c *Code) StatusCode() int {
	i, err := strconv.Atoi((c.String())[0:3])
	if err != nil {
		return 500
	}

	return i
}

func (c *Code) String() string {
	return strconv.Itoa(int(*c))
}

const (
	ErrCodeBadRequest           Code = 4000000
	ErrCodeResourceNotFound     Code = 4000001
	ErrCodeInvalidRequestParams Code = 4000002
	ErrCodeInternalError        Code = 5000000
)

func initErrCodeMapping() {
	errCode = map[Code]string{
		ErrCodeBadRequest:           "Bad request",
		ErrCodeResourceNotFound:     "Resource not found",
		ErrCodeInvalidRequestParams: "Invalid request params",
		ErrCodeInternalError:        "Internal error",
	}
}

func init() {
	initErrCodeMapping()
}

func trans(code Code) string {
	return errCode[code]
}
