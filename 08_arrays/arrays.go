package main

import "fmt"

// numbered sequence of specific length
func main() {
	// zeroed values
	// int -> 0, string -> "", bool -> false
	var nums [4]int // declare => [0, 0, 0, 0]

	// array length
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)

	var vals [5]bool
	vals[2] = true
	fmt.Println(vals)

	var names [5]string
	names[0] = "golang"
	fmt.Println(names)

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// 2D array
	arr1 := [2][2]int{{10, 20}, {30, 40}}
	fmt.Println(arr1)

	// - fixed size
	// - Memory optimazation
	// - Contant time access
}
