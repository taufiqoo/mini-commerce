package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Failure struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type Delete struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func APIResponse(code int, message string, data interface{}) Response {
	var response = Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return response
}

func APIFailure(code int, message string, err interface{}) Failure {
	var failure = Failure{
		Code:    code,
		Message: message,
		Error:   err,
	}
	return failure
}

func FormatDelete(code int, message interface{}) Delete {
	var formatDelete = Delete{
		Code:    code,
		Message: message,
	}
	return formatDelete
}

func SplitErrorInformation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
