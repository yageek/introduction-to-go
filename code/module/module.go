//Package module
func Square(n int) int {
	return n * n
}

//Private to package because of 's' (lowercase)
func pow3(n int) int {
	return Square(n) * n
}
