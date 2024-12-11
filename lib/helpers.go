package lib

// Returns true if c is an uppercase letter, a lowercase letter, or a digit.
func IsAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// Returns the absolute value of a given integer
func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
