package main

import (
	"fmt"
	"slices"
)

// much like arraylist - dynamic array
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice in nil
	var nums []int // (array -> var nums [n]int)
	fmt.Println(nums == nil)
	fmt.Println(nums)
	fmt.Println(len(nums))

	// make([]data-type, initial-size, capacity)
	// capacity ->  maximum numbers of elements can fit
	var nums1 = make([]int, 0, 5)
	// fmt.Println(cap(nums1))
	// fmt.Println(nums1 == nil)

	// In Go, a slice’s capacity is not a hard limit — it’s just the current underlying array size allocated for that slice.
	// When you append beyond its capacity, Go automatically allocates a new, bigger underlying array, copies the old elements into it, and then appends the new ones. capacity -> 2x

	nums1 = append(nums1, 1, 3, 5, 7, 9, 11, 13)
	nums1[0] = 10

	fmt.Println(nums1)
	fmt.Println(cap(nums1))

	// var nums2 = make([]int, 0, 5)
	// nums2 = append(nums2, 1, 2)
	// var nums3 = make([]int, len(nums2))

	// fmt.Println(nums2)
	// fmt.Println(nums3)

	// copy function
	// fmt.Println(copy(nums3, nums2)) // Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	// fmt.Println(nums3)

	// var arr = []int{1, 2, 3, 4, 5}
	// fmt.Println(arr[0:1])
	// fmt.Println(arr[:2])
	// fmt.Println(arr[1:])

	// slices
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 4}

	fmt.Println(slices.Equal(slice1, slice2))

	var arr = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr)
}
