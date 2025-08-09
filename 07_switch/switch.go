package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	day := 6

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work day")
	}

	// type switch
	whoAmI := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a bool")
		case float32, float64:
			fmt.Println("It's a float")
		default:
			fmt.Println("Some other type ")
		}
	}

	whoAmI("ritochit")
}
