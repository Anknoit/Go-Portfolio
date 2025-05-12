package main

import "fmt"

// link - https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp_arr := []int{}
	// Appending m elements into array
	for i := 0; i < m; i++ {
		fmt.Println(i)
		temp_arr = append(temp_arr, nums1[i])

	}
	fmt.Println(temp_arr)
	// Appending n elements into array

	for k := 0; k < n; k++ {
		fmt.Println(k)
		temp_arr = append(temp_arr, nums2[k])

	}
	fmt.Println(temp_arr)
	fmt.Println(len(temp_arr))
	nums1 = temp_arr
	// Ascending order conversion
	for j := 0; j < len(nums1)-1; j++ {
		for h := 0; h < len(nums1)-j-1; h++ {
			if nums1[h] > nums1[h+1] {
				nums1[h], nums1[h+1] = nums1[h+1], nums1[h]
			}
		}
	}
	fmt.Println(nums1)

}

func main() {
	merge([]int{0}, 0, []int{1}, 1)
}
