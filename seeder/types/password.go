package types

//Password has the options of Password type
type Password struct {
	min                    int
	max                    int
	allowUpper             bool
	allowNumeric           bool
	allowSpecialCharacters bool
}

//NewPassword is a constructor for Password
func NewPassword(min int, max int, allowUpper bool, allowNumeric bool, allowSpecialCharacters bool) Password {

	return Password{
		min:                    min,
		max:                    max,
		allowUpper:             allowUpper,
		allowNumeric:           allowNumeric,
		allowSpecialCharacters: allowSpecialCharacters,
	}
}

// Min is a helper function that implements Format() from Password interface
func (p Password) Min() int {
	return p.min
}

// Max is a helper function that implements Format() from Password interface
func (p Password) Max() int {
	return p.max
}

// AllowUpper is a helper function that implements Format() from Password interface
func (p Password) AllowUpper() bool {
	return p.allowUpper
}

// AllowNumeric is a helper function that implements Format() from Password interface
func (p Password) AllowNumeric() bool {
	return p.allowNumeric
}

// AllowSpecialCharacters is a helper function that implements Format() from Password interface
func (p Password) AllowSpecialCharacters() bool {
	return p.allowSpecialCharacters
}
