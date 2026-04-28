package luhn

import (
	"strings"
)

func Valid(id string) bool {
	// TrimmedId will be all numbers, no spaces, and sufficient length.
	TrimmedId := strings.ReplaceAll(id, " ", "")
	if len(TrimmedId) < 2 {
		return false
	}
	SumOfDigits := 0
	for k, v := range TrimmedId {
		// this will check that each rune is actually between 0 and 9 in the unicode digit.
		if v < '0' || v > '9' {
			return false
		}
		positionFromRight := len(TrimmedId) - 1 - k
		if positionFromRight%2 == 1 {
			// currentNumber is the number, a rune converted to int.
			currentNumber := int(TrimmedId[k] - '0')
			if currentNumber*2 > 9 {
				SumOfDigits += ((currentNumber * 2) - 9)
			} else {
				SumOfDigits += currentNumber * 2
			}
		} else {
			// else we just add the number to the sum of the digits
			// currentNumber is the number, a rune converted to int.
			currentNumber := int(TrimmedId[k] - '0')
			SumOfDigits += currentNumber
		}
	}
	return SumOfDigits%10 == 0
}