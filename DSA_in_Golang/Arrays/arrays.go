package main

import (
	"fmt"
	"sort"
)

// Finding min value in an array
// 1st Method - comparing digits in list, two pointer method
func array_ds(min_arr []int) {
	if len(min_arr) < 2 {
		fmt.Println("Array must have atleast 2 Elements!")
	}
	m, n := 0, 1
	for n < len(min_arr) {
		if min_arr[n] < min_arr[m] {
			m = n
		}
		n++
	}
	fmt.Println(min_arr)
	if m != 0 {
		min_arr[0], min_arr[m] = min_arr[m], min_arr[0]
		fmt.Println(min_arr)
	}
	fmt.Println("Shortest element", min_arr[0])
}

// 2nd Method - Arranging in ascending order
func ascending_arr(min_arr []int) {
	sort.Ints(min_arr) //Will directly sort the list in ascending order
	fmt.Println("Shortest value", min_arr[0])
}

// PALINDROME
func isPallindrome(s string) bool {
	//two pointer method
	m, n := 0, len(s)-1
	if s[m] != s[n] {
		fmt.Println("Not a Pallindrome")
		return false
	} else {
		m++
		n--
	}
	fmt.Println("Is Pallindrome")
	return true
}

func main() {

	array_ds([]int{4, 0, 1})
	ascending_arr(([]int{5, 3, 1, 6, 7, 8}))
	isPallindrome("racecar")
}
