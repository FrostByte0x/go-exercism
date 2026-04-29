package phonenumber

import (
	"fmt"
	"log/slog"
	"strings"
)

// Number should probably return just the numbers in the string. Tests are not clear.
func Number(phoneNumber string) (string, error) {
	var result strings.Builder
	// Go over each rune in the input string, and only keep numbers
	for _, v := range phoneNumber {
		if v >= '0' && v <= '9' {
			result.WriteRune(v)
		}
	}
	cleanedUpString := result.String()
	slog.Info(fmt.Sprintf("Number found:%s", cleanedUpString))
	// Validation time!
	// Rules for NANP
	// Max length is 11, min is 10 after cleanup
	if len(cleanedUpString) < 10 || len(cleanedUpString) > 11 {
		return "", fmt.Errorf("Invalid number: rule length")
	}
	// If length is 11, the first byte can only be 1 (US Country code)
	if len(cleanedUpString) == 11 && !(cleanedUpString[0] == '1') {
		return "", fmt.Errorf("Invalid number: rule 1")
	}
	// If Length is 10, the first byte CANNOT be 1, it must be 2-9
	if len(cleanedUpString) == 10 && cleanedUpString[0] == '1' {
		return "", fmt.Errorf("Invalid number: rule 2")
	}
	// Also the country code is not part of the NANP, it must be stripped.
	if len(cleanedUpString) == 11 {
		cleanedUpString = cleanedUpString[1:]
	}
	// Also the 1st and 4th number cannot be 1 or 0
	if cleanedUpString[0] < '2' || cleanedUpString[3] < '2' {
		return "", fmt.Errorf("Invalid number: rule 3")
	}
	return cleanedUpString, nil
}

// AreaCode returns the area code .
// It is the first 3 numbers, unless 1 is the first phone number (only for NANP, USA Canade etc)
func AreaCode(phoneNumber string) (string, error) {
	// Clean up the input by only keeping the unicode caracters between 0 and 9
	// But wait! There can be a country code, which is 1.
	phoneString, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	if len(phoneString) == 10 {
		return phoneString[:3], nil
	} else if len(phoneString) == 11 {
		return phoneString[1:4], nil
	} else {
		return "", fmt.Errorf("Invalid phone number")
	}
}

// Format output a formatted NANP phone number, such as (613) 995-0253
func Format(phoneNumber string) (string, error) {
	phoneString, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	starterIndex := len(phoneString)
	areaCode, err := AreaCode(phoneString)
	if err != nil {
		return "", err
	}
	// We'll format:
	// The Area code
	// 3 numbers (exchange code)
	// 4 numbers (subscriber code)
	// Try and work backward from the formatted number
	return fmt.Sprintf("(%s) %s-%s", areaCode, phoneString[starterIndex-7:starterIndex-4], phoneString[starterIndex-4:starterIndex]), nil
}
