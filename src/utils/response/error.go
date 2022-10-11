package response

type ErrorResponse struct {
	Message string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

func SendErrRes(message string) *ErrorResponse {
	return &ErrorResponse{message}
}
