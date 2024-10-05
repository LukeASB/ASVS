package validation

import "strings"

type NewLineCharacterEscape struct{}

func NewNewLineCharacterEscape() *NewLineCharacterEscape {
	return &NewLineCharacterEscape{}
}

func (n *NewLineCharacterEscape) CheckForNewLineCharacter(input string) bool {
	index := strings.IndexByte(input, '\n')
	return index >= 0
}
