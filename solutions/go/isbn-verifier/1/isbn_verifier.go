package isbnverifier

import (
	"fmt"
	"strings"
)

// IsValidISBN checks whether an ISBN number is valid.
// It must support X instead of 10.
func IsValidISBN(isbn string) bool {
	// too short or too long
	if len(isbn) < 10 || len(isbn) > 13 {
		return false
	}
	if strings.Count(isbn, "-") != 3 && len(isbn) > 10 {
		return false
	}
	// Parse the ISBN
	if strings.Index(strings.ToUpper(isbn), "X") != len(isbn)-1 && strings.Contains(strings.ToUpper(isbn), "X") {
		return false
	}
	ParsedISBN := strings.ReplaceAll(isbn, " ", "")
	ParsedISBN = strings.ReplaceAll(ParsedISBN, "-", "")
	checkSum := 0
	for k, v := range ParsedISBN {
		// Digit is equal to 40 if isbn trails with an X.
		digit := int(v - '0')
		// Absolute cinema - I can get fired for this right?
		if digit == 40 {

			digit = 10
		}
		// Helps debugging.
		fmt.Printf("Processing k: %d, v: %d, digit: %d\n", k, v, digit)
		checkSum += digit * (10 - k)
	}
	fmt.Printf("Checksum: %d\n", checkSum)
	return checkSum%11 == 0
}