package e

var MsgFlags = map[int]string{
	SUCCESS:        "success",
	INVALID_PARAMS: "invalid params",
	ERROR:          "server error",

	POLYPAY_SUCCESS: "success",
	POLYPAY_ERROR:   "error",
}

// GetMessage get error information based on Code
func GetMessage(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
