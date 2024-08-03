package cool_obj

const (
	Male           = "male"
	Female         = "female"
	PreferNotToSay = "prefer not to say"
)

// Thank you chatgpt

// MonthMap holds the mapping between month names and their corresponding order.
var MonthMap = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

// OrderToMonthMap holds the reverse mapping between order and month names.
var OrderToMonthMap = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

// MonthToOrder converts a month name to its corresponding order.
func MonthToOrder(month string) (int, error) {
	if order, exists := MonthMap[month]; exists {
		return order, nil
	}
	return 0, errInvalidMonthName
}

// OrderToMonth converts an order to its corresponding month name.
func OrderToMonth(order int) (string, error) {
	if month, exists := OrderToMonthMap[order]; exists {
		return month, nil
	}
	return "", errInvalidMonthOrdre
}

// IsLeapYear checks if a year is a leap year.
func IsLeapYear(y Year) bool {
	if y%4 == 0 {
		if y%100 == 0 {
			return y%400 == 0
		}
		return true
	}
	return false
}
