package result

func New(code Code, msg string) *Result {
	return &Result{
		Code:    code,
		Message: msg,
	}
}

func OK() *Result {
	return &Result{
		Code:    CodeOK,
		Message: "OK",
	}
}

func Error() *Result {
	return &Result{
		Code:    CodeInternalError,
		Message: "Error!",
	}
}

type Result struct {
	Code    Code
	Message string
}

func (r *Result) IsOK() bool {
	return r.Code == CodeOK
}
