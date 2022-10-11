package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
		// return fiber.ErrBadRequest
	}

	return nil
}

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return Validate(body)
}

// NewValidator func for create a new validator for model fields.
func Validate(payload interface{}) *fiber.Error {
	err := validate.Struct(payload)

	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(
				errors,
				fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", err.Field(), err.Value(), err.Tag()),
			)
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errors, ","),
		}
	}

	return nil
}

// // ValidatorErrors func for show validation errors for each invalid fields.
// func ValidatorErrors(err error) map[string]string {
// 	// Define fields map.
// 	fields := map[string]string{}

// 	// Make error message for each invalid field.
// 	for _, err := range err.(vldtr.ValidationErrors) {
// 		errMsg := fmt.Sprintf("validation failed on '%s' tag", err.Tag())
// 		param := err.Param()
// 		if param != "" {
// 			errMsg = fmt.Sprintf("%s. allow: %s", errMsg, param)
// 		}
// 		fields[err.Field()] = errMsg
// 	}

// 	return fields
// }
