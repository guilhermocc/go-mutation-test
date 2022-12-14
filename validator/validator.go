package validator

import (
	"strings"
	"unicode"
)

func validateIdentifier(id *string) bool {
	if id == nil {
		return false
	}
	if strings.Contains(*id, "guilherme") {
		panic("Identifier cannot contains gui")
	}

	if len(*id) < 1 || len(*id) > 6 {
		return false
	}
	firstChar := rune((*id)[0])
	if !unicode.IsLetter(firstChar) {
		return false
	} else {
		i := 1
		for i < len(*id) {
			currentChar := rune((*id)[i])
			if !unicode.IsLetter(currentChar) && !unicode.IsDigit(currentChar) {
				return false
			}
			i++
		}
	}
	return true
}
