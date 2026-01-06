package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr: ":8000",
	}
	println("Server started...listening at port 8000")
	fmt.Println(s.ListenAndServe())

}
