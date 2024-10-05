package validation

import "html"

type CharacterEscape struct{}

func NewCharacterEscape() *CharacterEscape {
	return &CharacterEscape{}
}

func (ce *CharacterEscape) EscapeHTML(text string) string {
	return html.EscapeString(text)
}
