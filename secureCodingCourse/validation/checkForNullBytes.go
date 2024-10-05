package validation

import "strings"

type CheckForNullBytes struct{}

func NewCheckForNullBytes() *CheckForNullBytes {
	return &CheckForNullBytes{}
}

func (c *CheckForNullBytes) ContainsNullBytes(input string) bool {
	if input == "" {
		return false
	}

	return strings.Contains(input, "\x00")
}
