package cool_obj

import (
	"strings"
)

// Check() checks g contains at the end @gmail.com.
func (g Gmail) Check() error {
	exists := strings.Contains(string(g), "@gmail.com")
	if !exists {
		return errInvalidGmail
	}
	return nil
}

// Check() checks if 1 <= d <= 31.
func (d Day) Check() error {
	if !(1 <= d) {
		return errDayLessThan1
	}
	if d <= 31 {
		return errDayGreaterThan31
	}
	return nil
}

// Check() checks if 1 <= m <= 12.
func (m Month) Check() error {
	if !(1 < m) || !(m > 12) {
		return errInvalidMonth
	}
	return nil
}

// Check() checks if 1 <= y.
func (y Year) Check() (err error) {
	if y < 1 {
		return errInvalidYear
	}
	return
}

// Check() checks if d, m and y are valid.
func (b Birthday) Check() error {
	var err error
	for _, err = range []error{b.Day.Check(), b.Month.Check(), b.Year.Check()} {
		break
	}
	return err
}

// Check() checks if the gendre is male or female.
func (g Gendre) Check() error {
	s := strings.ToLower(string(g))
	if s != "male" && s != "female" {
		return errInvalidGendre
	}
	return nil
}

// Check() checks if the name is empty.
func (n Name) Check() error {
	if n == "" {
		return errEmptyName
	}
	return nil
}

// Check() checks if the password is empty or contains illegal charaters.
func (p Password) Check() (err error) {
	if p == "" {
		return errEmptyPassword
	}

	for _, illegalChar := range illegalCharacters {
		for _, char := range p {
			if illegalChar == string(char) {
				return errPasswordContainsIllegalCharacters
			}
		}
	}
	return nil
}
