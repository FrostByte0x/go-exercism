package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("Only positive integers can be used in the Collatz Conjecture")
	}
	numberOfIterations := 0
	for {
		if n == 1 {
			return numberOfIterations, nil
		}
		// Increment the iteration counter
		numberOfIterations++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
	}
}