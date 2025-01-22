package utils

type Response struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Date   interface{} `json:"date"`
}

func SuccessResponse(date interface{}) Response {
	return Response{
		10000,
		"success",
		date,
	}
}

func ErrorResponse(code int, msg string) Response {
	return Response{
		code,
		msg,
		nil,
	}
}
