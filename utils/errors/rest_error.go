package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewNotFoundError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  404,
		Error:   "Not found",
	}
}

func NewInternalServerError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  500,
		Error:   "Internal server error",
	}
}
