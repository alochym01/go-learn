package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("want '%d' but got '%d'", want, got)
		}
	}
	got := Add(2, 2)
	want := 4
	assertCorrectMessage(t, got, want)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
