package result

type Code int

const (
	CodeUnspecified Code = iota
	CodeOK
	CodeNotFound
	CodeInternalError
)
