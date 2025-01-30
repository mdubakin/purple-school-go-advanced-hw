package request

import "github.com/go-playground/validator/v10"

func validate[T any](payload T) error {
	v := validator.New()
	return v.Struct(payload)
}
