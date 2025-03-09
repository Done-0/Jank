package utils

import "github.com/go-playground/validator/v10"

type (
	ValidErrRes struct {
		Error bool
		Field string
		Tag   string
		Value interface{}
	}
)

var NewValidator = validator.New()

// Validator 参数验证器
func Validator(data interface{}) []ValidErrRes {
	var Errors []ValidErrRes
	errs := NewValidator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var el ValidErrRes
			el.Error = true
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Value()

			Errors = append(Errors, el)
		}
	}
	return Errors
}
