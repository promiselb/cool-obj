package cool_obj

import (
	"errors"
)

var (
	errDayLessThan1                      = errors.New("Day < 1")
	errDayGreaterThan31                  = errors.New("Day > 31")
	errInvalidMonth                      = errors.New("invalid month of birth? Must be 1 <= m <= 12! ")
	errInvalidYear                       = errors.New("invalid year of birth? Must be 1 < y <= 2024! ")
	errInvalidGendre                     = errors.New("invalid Gendre?")
	errEmptyName                         = errors.New("empty name?")
	errEmptyPassword                     = errors.New("empty password?")
	errPasswordContainsIllegalCharacters = errors.New("password contains illegal characters?")
	errInvalidGmail                      = errors.New("invalid gmail? Must ends with '@gmail.com'! ")
	errInvalidMonthName                  = errors.New("invalid month name?")
	errInvalidMonthOrdre                 = errors.New("invalid month order")
)
