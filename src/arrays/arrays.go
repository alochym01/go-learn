package arrays

// For function takes an array as parameter and return total sum of array's elements
func For(numbers [5]int) int {
	var sum int = 0

	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}

// Range function takes an array as parameter and return total sum of array's elements
func Range(numbers [5]int) int {
	var sum int = 0
	// range lets you iterate over an array.
	// Every time it is called it returns two values, the index and the value
	for _, number := range numbers {
		sum += number
	}

	return sum
}
