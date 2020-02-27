package constant

const (
	// Success ...
	Success = 200
	// BadRequest ...
	BadRequest = 400
	// Unauthorized ...
	Unauthorized = 401
	// Forbidden ...
	Forbidden = 403
	// NotFound ...
	NotFound = 404
	// InternalError ...
	InternalError = 500
)

var codeMsg = map[int]string{
	Success:       "success",
	BadRequest:    "bad request",
	Unauthorized:  "not authorized",
	Forbidden:     "forbidden request",
	NotFound:      "not found",
	InternalError: "internal server error",
}

// GetMsg converts code to message
func GetMsg(code int) string {
	msg, ok := codeMsg[code]
	if ok {
		return msg
	}
	return "failed to get message code"
}
