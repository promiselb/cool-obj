package cool_obj

import (
	"errors"
	"strings"
)

type Checker interface {
	Check() error
}

// Check() checks g ends with @gmail.com.
func (g Gmail) Check() error {
	if !strings.HasSuffix(string(g), "@gmail.com") {
		return errInvalidGmail
	}
	return nil
}

// Check() checks if 1 <= d <= 31.
func (d Day) Check() error {
	if d < 1 {
		return errDayLessThan1
	}
	if d > 31 {
		return errDayGreaterThan31
	}
	return nil
}

// Check() checks if 1 <= m <= 12.
func (m Month) Check() error {
	if !(1 <= m) || !(m <= 12) {
		return errInvalidMonth
	}
	return nil
}

// Check() checks if 1 < y <= 2024.
func (y Year) Check() error {
	if !(1 < y && y <= 2024) {
		return errInvalidYear
	}
	return nil
}

// Check() checks if d, m and y are valid.
func (b Birthday) Check() error {
	var err error
	for _, err = range []error{b.Day.Check(), b.Month.Check(), b.Year.Check()} {
		if err != nil {
			return err
		}
	}

	switch b.Month {
	case 1, 3, 5, 7, 8, 10, 12:
		if b.Day > 31 {
			return errors.New("invalid day: must be between 1 and 31 for the given month")
		}
	case 4, 6, 9, 11:
		if b.Day > 30 {
			return errors.New("invalid day: must be between 1 and 30 for the given month")
		}
	case 2:
		if IsLeapYear(b.Year) {
			if b.Day > 29 {
				return errors.New("invalid day: must be between 1 and 29 for February in a leap year")
			}
		} else {
			if b.Day > 28 {
				return errors.New("invalid day: must be between 1 and 28 for February in a non-leap year")
			}
		}
	default:
		return errors.New("invalid month")
	}

	return nil
}

// Check() checks if the gendre is male or female.
func (g Gendre) Check() error {
	s := strings.ToLower(string(g))
	if s != Male && s != Female && s != PreferNotToSay {
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
