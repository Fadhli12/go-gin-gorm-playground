package common

type RequestError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (r *RequestError) Error() string {
	return r.Message
}

func ErrorRequest(message string, code int) error {
	return &RequestError{
		StatusCode: code,
		Message:    message,
	}
}
