package slices

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	}

	// Slices can have any capacity
	slices := []int{1, 2, 3, 4, 5}

	// Running subsequences testing function of slices.go file
	t.Run("Slice testing using for loop", func(t *testing.T) {
		got := For(slices)
		want := 15
		assertCorrectMessage(t, got, want)
	})

	// Running subsequences testing function of slices.go file
	t.Run("Array test using range loop", func(t *testing.T) {
		got := Range(slices)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}
