package result

import "net/http"

type Code int

const (
	CodeUnspecified Code = iota
	CodeOK
	CodeNotFound
	CodeBadRequest
	CodeForbidden
	CodeInternalError
)

var codeStringMap = map[Code]string{
	CodeNotFound:      "not found",
	CodeBadRequest:    "bad request",
	CodeForbidden:     "forbidden",
	CodeInternalError: "error",
}

func (c Code) String() string {
	s, ok := codeStringMap[c]
	if !ok {
		return "error"
	}
	return s
}

var statusCodeMap = map[Code]int{
	CodeOK:            http.StatusOK,
	CodeNotFound:      http.StatusNotFound,
	CodeBadRequest:    http.StatusBadRequest,
	CodeForbidden:     http.StatusForbidden,
	CodeInternalError: http.StatusInternalServerError,
}

func (c Code) ToStatusCode() int {
	status, ok := statusCodeMap[c]
	if !ok {
		return http.StatusInternalServerError
	}
	return status
}
