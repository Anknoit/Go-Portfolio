package main

import (
	"fmt"
)

// FIBONNACI ALGORITHM
// 1. Two variables to hold the first two fibonacci no.s
// 2. add two variables to produce new fibonnaci no.
// 3. now add var2 + var_new_no. and so on

func main() {

	fmt.Println("Understanding For Loops and Recursion")
	fmt.Println("--For Loop Method")

	// fibonacci serires using FOR Loops
	var a int = 0
	var b int = 1
	itr := 18 // No. of iteration
	fibo_arr := []int{a, b}
	for i := 0; i <= itr; i++ {
		c := fibo_arr[i] + fibo_arr[i+1]
		fibo_arr = append(fibo_arr, c)
	}
	fmt.Println(fibo_arr)

	// fibonacci serires using Recursion
}
