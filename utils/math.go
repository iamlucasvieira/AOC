package utils

// GCD returns the Greatest Common Divisor of a and b
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns the Least Common Multiple of a and b and integers
func LCM(integers ...int) int {

	if len(integers) == 0 {
		return 0
	}

	a := integers[0]
	b := integers[1]

	result := a * b / GCD(a, b)

	if len(integers) == 2 {
		return result
	}

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
