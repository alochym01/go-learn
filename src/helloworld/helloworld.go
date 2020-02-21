package main

import "fmt"

// Hello function return "Hello, World" string
func Hello(name string) string {
	return "Hello, " + name
	// Output: "Hello, World"
}
func main() {
	fmt.Println(Hello("Ha"))
}
