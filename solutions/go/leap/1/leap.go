package leap

// IsLeapYear returns whether a year is a leap year or not:
// It must be divisible by 4, and if it's the start of the century, it must also be divisible by 400
func IsLeapYear(year int) bool {
	if year%100 == 0 && year%400 == 0 {
		return true
	} else if year%100 == 0 && year%400 != 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
