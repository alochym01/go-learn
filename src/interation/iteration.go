package interation

// Repeat function takes a string as a parameter and return a string
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
