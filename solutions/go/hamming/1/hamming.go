package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("a and b string length must be the same!")
	}
	differences := 0
	for key, _ := range []byte(a) {
		if a[key] != b[key] {
			differences++
		}
	}
	return differences, nil
}
