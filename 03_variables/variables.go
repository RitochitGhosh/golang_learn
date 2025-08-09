package main

import "fmt"

func main() {
	// var name string = "golang"\
	// infer
	var name = "ritochit ghosh"
	var isAdult bool = true
	var age int = 20
	fmt.Println(name, ", Age - ", age, ", IsAdult - ", isAdult)

	// shorthand syntax
	word := "golang"
	fmt.Println(word)

	// use scenarios
	var price float32 // declaring first
	price = 49.5      // initialyze later

	fmt.Println(price)
}
