package validation

type whiteList struct {
	allowedValues map[string]bool
}

func NewWhiteList() *whiteList {
	return &whiteList{
		allowedValues: map[string]bool{
			"Luke":   true,
			"Trevor": true,
		},
	}
}

func (wl *whiteList) Check(input string) bool {
	_, ok := wl.allowedValues[input]

	return ok
}
