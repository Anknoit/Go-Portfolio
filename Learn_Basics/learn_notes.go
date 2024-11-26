// ABOUT GO
/*-Cross platform
-Faster than C++
-Supports Mutithreadin (Goroutine, connsumes 2KB compared to normal 1MB used by other languages)\
-Supports Concurrent Programming applications
-Great Memory Management
-Statically Typed (i.e type checking before execution => variables are defined befoore execution duromg complie time, unlike python where variables are dtermined during the runtime which makes it slower and susceptible to errors.)
*/

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

}
