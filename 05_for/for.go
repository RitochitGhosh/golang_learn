package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop
	fmt.Println("While loop:")
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i += 1
	}

	// infinite loop
	// for {

	// }

	fmt.Println("For loop:")
	// classic for loop
	for j := 0; j < 10; j++ {
		if j == 5 {
			break
		}
		if j == 3 {
			continue
		}
		fmt.Println(j)
	}

	fmt.Println("Range loop:")
	// range
	for k := range 3 { // 0 - (n - 1)
		fmt.Println(k)
	}
}
