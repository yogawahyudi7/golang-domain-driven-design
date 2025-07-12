package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator wraps the go-playground validator
type Validator struct {
	validator *validator.Validate
}

// New creates a new validator instance
func New() *Validator {
	v := validator.New()

	// Register custom tag name function to use json tags
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		validator: v,
	}
}

// Validate validates a struct
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// ValidateField validates a single field
func (v *Validator) ValidateField(field interface{}, tag string) error {
	return v.validator.Var(field, tag)
}

// GetValidationErrors returns formatted validation errors
func (v *Validator) GetValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			tag := fieldError.Tag()

			switch tag {
			case "required":
				errors[field] = field + " is required"
			case "email":
				errors[field] = field + " must be a valid email address"
			case "min":
				errors[field] = field + " must be at least " + fieldError.Param() + " characters long"
			case "max":
				errors[field] = field + " must be at most " + fieldError.Param() + " characters long"
			default:
				errors[field] = field + " is invalid"
			}
		}
	}

	return errors
}
