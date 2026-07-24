package validator

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/YogaRP/finansial/transaction-service/internal/model"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

// Get returns a singleton validator instance.
func Get() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		RegisterCustomValidations(validate)
	})
	return validate
}

// Validate runs struct validation and returns a readable error message.
func Validate(s any) error {
	if err := Get().Struct(s); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return fmt.Errorf("%s", formatErrors(validationErrors))
		}
		return err
	}
	return nil
}

func formatErrors(errs validator.ValidationErrors) string {
	messages := make([]string, 0, len(errs))
	for _, e := range errs {
		messages = append(messages, formatFieldError(e))
	}
	return strings.Join(messages, "; ")
}

func formatFieldError(e validator.FieldError) string {
	field := strings.ToLower(e.Field())

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters/value", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters/value", field, e.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", field, e.Param())
	default:
		return fmt.Sprintf("%s is invalid (%s)", field, e.Tag())
	}
}

func RegisterCustomValidations(v *validator.Validate) {
	v.RegisterValidation("trxcategory", func(fl validator.FieldLevel) bool {
		switch value := fl.Field().Interface().(type) {
		case model.TrxCategory:
			return value.IsValid()
		case *model.TrxCategory:
			if value == nil {
				return true
			}
			return value.IsValid()
		case string:
			return model.TrxCategory(value).IsValid()
		case *string:
			if value == nil {
				return true
			}
			return model.TrxCategory(*value).IsValid()
		default:
			return false
		}
	})
}
