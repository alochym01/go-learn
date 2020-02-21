package interation

const repeatCount = 5

// RepeatFactor function takes a string as a parameter and return a string
func RepeatFactor(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
