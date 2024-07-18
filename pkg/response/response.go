package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusError = "Error"
	StatusOK    = "OK"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(vr validator.ValidationErrors) Response {
	var errMsgs []string
	for _, err := range vr {
		switch err.ActualTag() {
		case "min":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s must be at least %s characters", err.Field(), err.Param()))
		case "max":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s must be at most %s characters", err.Field(), err.Param()))
		case "email":
			errMsgs = append(errMsgs, fmt.Sprintf("invalid email format"))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("this field is not valid"))
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
