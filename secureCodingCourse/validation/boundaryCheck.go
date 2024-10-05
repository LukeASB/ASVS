package validation

import (
	"fmt"
	"strconv"
)

type BoundaryCheck struct{}

func NewBoundaryCheck() *BoundaryCheck {
	return &BoundaryCheck{}
}

func (bc *BoundaryCheck) Validate(value string) (int, error) {
	num, err := strconv.Atoi(value)

	if err != nil {
		return 0, fmt.Errorf("Invalid Input")
	}

	if num < 1 || num > 100 {
		return 0, fmt.Errorf("Integer out of range!")
	}

	return num, nil
}
