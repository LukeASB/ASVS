package validation

import (
	"strconv"
)

type NumericValidation struct{}

func NewNumericValidation() *NumericValidation {
	return &NumericValidation{}
}

func (ne *NumericValidation) IsNumber(input string) bool {
	_, err := strconv.Atoi(input)

	return err == nil
}
