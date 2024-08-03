package cool_obj

// These are the illegal characters by default:
// () {} [] #
var illegalCharacters = []string{"(", ")", "{", "}", "[", "]", "#"}

// These are the illegal characters by default:
// () {} [] #
func SetIllegalChars(chars ...string) {
	illegalCharacters = chars
}
