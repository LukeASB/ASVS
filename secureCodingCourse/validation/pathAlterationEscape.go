package validation

import "path"

type PathAlterationEscape struct{}

func NewPathAlterationEscape() *PathAlterationEscape {
	return &PathAlterationEscape{}
}

func (p *PathAlterationEscape) IsValidPath(input string) bool {
	return path.Clean(input) == input
}
