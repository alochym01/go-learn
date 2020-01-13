package calc

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	t.Run("Running test add function", func(t *testing.T) {
		got := Add(2, 3)
		want := 5
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Running test substract function", func(t *testing.T) {
		got := Substract(3, 2)
		want := 1
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Running test multiply function", func(t *testing.T) {
		got := Multiply(2, 3)
		want := 6
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Running test divide function", func(t *testing.T) {
		got := Divide(4, 2)
		want := 2
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func ExampleSubstract() {
	sum := Substract(1, 5)
	fmt.Println(sum)
	// Output: -4
}

func ExampleMultiply() {
	sum := Multiply(1, 5)
	fmt.Println(sum)
	// Output: 5
}

func ExampleDivide() {
	sum := Divide(1, 5)
	fmt.Println(sum)
	// Output: 0
}
