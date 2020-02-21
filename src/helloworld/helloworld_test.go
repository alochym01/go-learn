package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// In Go you can declare functions inside other functions and assign them to variables
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// using Helper() function to refactor assertion function
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		// This will help other developers track down problems easier.
		t.Helper() // is needed to tell the test suite that this method is a helper
		if got != want {
			// print out a message and fail the test
			t.Errorf("got %q but want %q", got, want)
		}
	}
	// Run subtest with multiple values
	t.Run("say hello to Ha", func(t *testing.T) {
		got := Hello("Ha")
		want := "Hello, Ha"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// make Example for go doc
// func ExampleHello() {
// 	fmt.Println(Hello("Ha"))
// 	// Output: Hello, Ha
// }
