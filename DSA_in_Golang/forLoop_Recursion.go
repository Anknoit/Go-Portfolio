package main

import (
	"fmt"
)

// FIBONNACI ALGORITHM
// 1. Two variables to hold the first two fibonacci no.s
// 2. add two variables to produce new fibonnaci no.
// 3. now add var2 + var_new_no. and so on
func fibo_recur(p1 int, p2 int) {
	p1 = 0
	p2 = 1
	itr := 18
	fibo_recur_arr := []int{p1, p2}
	for i := 0; i <= itr; i++ {
		new_val := fibo_recur_arr[i] + fibo_recur_arr[i+1]
		fibo_recur_arr = append(fibo_recur_arr, new_val)
	}
	fmt.Println("Recursion Fibonacci", fibo_recur_arr)
}
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
	fibo_recur(0, 1)

}
