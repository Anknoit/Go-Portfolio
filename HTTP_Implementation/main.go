package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting HTTP Server on port 8000...")
	s := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server Started, Waiting for client requests..")
	fmt.Println(s.ListenAndServe())

}
