package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// DateValidation date yyyy-MM-dd format
func DateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return true
}

// DatetimeValidation date yyyy-MM-dd H:i:s format
func DatetimeValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	if err != nil {
		return false
	}
	return true
}

// DateRangeValidation date range validator
func DateRangeValidation(fl validator.FieldLevel) bool {
	var date = fl.Field().String()
	var minDate = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	var maxDate = time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)

	datetime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	if datetime.Before(minDate) || datetime.After(maxDate) {
		return false
	}
	return true
}
