package validation

import "unicode/utf8"

type CheckForExtendedUTF8Encoding struct{}

func NewCheckForExtendedUTF8Encoding() *CheckForExtendedUTF8Encoding {
	return &CheckForExtendedUTF8Encoding{}
}

func (u *CheckForExtendedUTF8Encoding) IsValidUTF8(input string) bool {
	return utf8.ValidString(input)
}
