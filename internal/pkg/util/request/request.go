package request

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FieldErrorCheck(err error) ([]ErrorMsg, error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))

		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}

		return out, nil
	}

	return nil, err
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return "Maximum value should be " + fe.Param()
	}
	return "Unknown error"
}
