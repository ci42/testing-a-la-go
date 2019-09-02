// +build exampletests

package main

func ExampleHello() {
	hello()
	// Output:
	// Hello Go Usergroup Rhein-Ruhr
}

func ExampleAlternatingHello() {
	alternatingHello()
	// Unordered output:
	// Hello World
	// Hello Go Usergroup Rhein-Ruhr
}
