package main

import (
	"fmt"
	"sort"
)

// Finding min value in an array
// 1st Method - comparing digits in list
func array_ds(min_arr []int) {
	for i := 0; i < len(min_arr)-1; i++ {
		if min_arr[0] < min_arr[i+1] {
			fmt.Println("Shortest Digit", min_arr[i])
		}
	}
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

	array_ds([]int{2, 2, 3, 4, 5, 6})
	ascending_arr(([]int{5, 3, 1, 6, 7, 8}))
	isPallindrome("racecar")
}
