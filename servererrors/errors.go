package servererrors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func NewFieldError(fieldErr validator.FieldError) string {
	var sb strings.Builder
	sb.WriteString("validation failed on field '" + fieldErr.Field() + "'")
	sb.WriteString(", condition: " + fieldErr.ActualTag())

	if fieldErr.Param() != "" {
		sb.WriteString("; param: '" + fieldErr.Param() + "'")
	}
	if fieldErr.Value() != nil && fieldErr.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", fieldErr.Value()))
	}
	return sb.String()
}
