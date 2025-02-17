// ABOUT GO
/*-Cross platform
-Faster than C++
-Supports Mutithreadin (Goroutine, connsumes 2KB compared to normal 1MB used by other languages)\
-Supports Concurrent Programming applications
-Great Memory Management
-Statically Typed (i.e type checking before execution => variables are defined befoore execution duromg complie time, unlike python where variables are dtermined during the runtime which makes it slower and susceptible to errors.)
*/
// TEST PUSH

//APPLICATIONS
/*
- Web Dev
- Cloud and NEtworking
- Devops
- BlockChain
- Networking Infrastructure Application
*/

// Hello World
package main

import (
	"fmt"
	// if we import a package witha underscore in front
	// The underscore import is a way to include a package without directly using its exported identifiers.
	// It is commonly used for packages that perform initialization or registration tasks.
	// It helps keep the code clean and avoids unused import errors while still allowing the necessary side effects to occur.
)

func main() {
	fmt.Println("Hello Fellow Warriors!")
	// To run >>go run file_name.go
	// To build an executable file >>go build filename.go

	// 1. VARIABLE TYPES AND USES
	// var var_name var_type = value (variable type given if known) (can be used inside and outside functions, value can be separately assigned later once var is declared)
	// var_name := value (Python style declaration, compiler decides the type by lookinng at the value, HENCE THE VALUE NEEDS TO BE GIVEN, variable is INFERRED, pnly inside function, value must be given)
	var empName string = "Ankit"  //Since we know it for a fact that it will be a string
	empNameBritish := "Ankit 111" //If its a name with integer then well inferr the variable

	fmt.Println(empName)
	fmt.Println(empNameBritish)

	//If no value is provided to the variable it self assigns the default values
	// int = 0
	// bool = false
	// string = ""

	//Multi Variable declaration
	var multi_1, multi_2, multi_3 int = 4, 5, 6
	var a, b = 1, "Hello"
	c, d := 6, "World"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//Block Declaration
	var (
		e int
		f int    = 1
		g string = "Some String"
	)
	fmt.Println("Default Integer", e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(multi_1)
	fmt.Println(multi_2)
	fmt.Println(multi_3)

	// Variables are CASE SENSITIVE - age, Age and AGE are three different variables

	//CONSTANTS
	// variables with fixed values
	// Can be declared inside or outside the fucntion
	// Typed Constant - type is known
	const AUTHENTICATION int = 1812
	// Untyped Constant - type not known
	const PASSWORD = "A@#@MKANSLD"
	// Multiple Constant Declaration
	const (
		A             int = 1812
		AUTHORIZATION int = 1813
		COA               = "Change of Authorization"
	)

	// OUTPUT - Printing
	/*
		-Print() - Prints with no spaces or lines, put space iff printing integers
		-Println() - Prints with new line and puts space between printing variables Println(i, j) = "i_value j_value"
		-Printf() - consists of formatting verbs for special case outputs = https://www.w3schools.com/go/go_formatting_verbs.php

	*/

	// Data Types
	/*
		-Boolean
		-String
		-Float
			-float32 & float64
		-int
			-Signed int = +ve and -ve values
				- int = 32(in 32 bit os systems) 64 bits(64 bits os systems)
				- int8 = -128 to 127 , var _num int8 = 500 -----> ERROR! FUCK YOU!
				- int16, 32, 64
			-Unsigned Integers = only +ve
				- uint
				- will have same ranges as above STARTING FROM 0
	*/

	// START FROM ARRAYS

	// ARRAYS
	var len_arr = [4]int{1, 2, 3, 4}            // Fininte length array
	var no_len_arr = [...]int{1, 2, 3, 4, 5, 6} // No length array
	no_var := [3]string{"BMW", "Audi", "Merc"}

	fmt.Print(no_var) //Using only Print() will not make new line for print
	fmt.Println((len_arr))
	fmt.Println(no_len_arr)
	fmt.Println("Acess Array Element:", len_arr[2])

	// Array Initialization
	not_initialized := [5]int{}
	partial := [4]int{1, 2}
	full := [3]int{1, 2, 3}
	specific := [5]int{1: 20, 3: 5} //NOT SLICING but assigning values to specific index numbers
	fmt.Println("ARRAY NOT INITIALIZED", not_initialized)
	fmt.Println("ARRAY PARTIALLY INITIALIZED", partial)
	fmt.Println("ARRAY FULLY INITIALIZED", full)
	fmt.Println("ARRAY SPECIFICALLY INITIALIZED", specific)
	fmt.Println("LENGTH OF ARRAY", len(specific))

	// START FROM SLICES

	// SLICES
	// Similar to a	rray but size is flexible
	myslice := []int{} //Initialized an empty slice of length=capacity=0
	myslice2 := []int{1, 2, 3, 4}
	fmt.Println(myslice)
	fmt.Println(myslice2)
	fmt.Println("length of slice1:", len(myslice))
	fmt.Println("capacity of slice1:", cap(myslice))
	fmt.Println("length of slice2:", len(myslice2))
	fmt.Println("capacity of slice2:", cap(myslice2))
	// Creating slice from an Array
	var arr_temp = [6]int{1, 2, 3, 4, 5, 6}
	slice_arr := arr_temp[1:4] // [start-include : end-index-1]
	fmt.Println("slice made from array:", slice_arr)
	fmt.Println("Length of slice made from array:", len(slice_arr))
	fmt.Println("Capacity of slice made from array:", cap(slice_arr)) //Capacity will be the count from first index in slice to last index in that array

	// Creating Slice from Make
	// make([]datatype, length, capacity)
	slice_make1 := make([]string, 4, 6)
	//Appending elements in slice
	values_in := append(slice_make1, "Apple", "Mango", "Grapes")

	fmt.Println("Slice made from make", slice_make1)
	fmt.Println("Values appended into Slice made from make", values_in)

	fmt.Println("Length of Slice made from make", len(slice_make1))
	fmt.Println("Capacity of Slice made from make", cap(slice_make1))
	//Why do we use slice when we have an option to define flexible length in arrays using [...], my guess is that it offers unlimited length for the array rather than flexible length according to the inputs

	//Combining two ormore slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3_combined := append(slice1, slice2...) // ... at the end is important
	fmt.Println("Combined SLice", slice3_combined)

	// Copying of part of elements in an array or slice and copying arrays into new one for memory efficiency

	//OPERATORS

}
