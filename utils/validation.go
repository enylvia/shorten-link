package utils

import "github.com/go-playground/validator/v10"

func ValidateRequest(request interface{}) Response {
	var errors []ValidationResponseError
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var val ValidationResponseError
			val.Field = err.StructField()
			val.Tag = err.Tag()
			val.Value = err.Param()
			errors = append(errors, val)
		}
	}
	return BadRequestResponse(errors, "Validation error")
}
