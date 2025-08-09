package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not adult")
	}

	num := -10
	if num > 0 {
		fmt.Println("Positive")
	} else if num == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative")
	}

	var role string = "admin"
	var hasPermission bool = false
	if role == "admin" && hasPermission == true {
		fmt.Println("Access given!")
	} else {
		fmt.Println("Access denied!")
	}

	// We can declare variable inside if construct
	if gender := "Male"; gender == "Male" {
		fmt.Println("Male")
	}

	// fmt.Println(gender) // can't access it outside if block. (Block scoped)

	// go doesn't have ternary operator

}
