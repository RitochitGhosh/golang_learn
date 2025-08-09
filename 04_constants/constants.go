package main

import "fmt"

const age = 20

// word := "new_string" // expected declaration, found word => shorthand syntax not allowed outside function

func main() {
	const name string = "golang"
	// name = "javascript" // cannot assign to name (neither addressable nor a map index expression)compiler

	fmt.Println(age)

	// constant grouping
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(host, port)
}
