package arrays

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	}

	// Arrays have a fixed capacity which you define when you declare the variable
	numbers := [5]int{1, 2, 3, 4, 5}

	// Running subsequences testing function of array.go file
	t.Run("Array test using for loop", func(t *testing.T) {

		got := For(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})

	// Running subsequences testing function of array.go file
	t.Run("Array test using range loop", func(t *testing.T) {
		got := Range(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}
