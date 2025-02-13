package main

import (
	"fmt"
)

func merge_arr(nums1 []int, m int, nums2 []int, n int) {
	for m := 0; m < len(nums1)-1; m++ {
		nums1 = append(nums1, nums1[m])
	}
	fmt.Println(nums1)
}

func main() {
	merge_arr([]int{1, 23, 5, 67, 87}, 3, []int{1, 5, 7, 98, 3}, 0)
}
