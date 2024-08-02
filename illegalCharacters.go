package cool_obj

// These are the illegal characters by default:
// () {} [] #
var illegalCharacters = []string{"(", ")", "{", "}", "[", "]", "#"}

func SetIllegalChars(chars ...string) {
	illegalCharacters = chars
}
