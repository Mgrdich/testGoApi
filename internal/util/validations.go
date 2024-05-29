package util

import "unicode"

func ValidatePassword(s string) (isValid, sevenOrMore, number, upper, special bool) {
	letters := 0

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}

	sevenOrMore = letters >= 7
	isValid = sevenOrMore && number && upper && special

	return
}
