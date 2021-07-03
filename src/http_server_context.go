package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	context := request.Context()
	fmt.Println("Serve hello handler started")
	defer fmt.Println("Serve hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(writer, "hello\n")
	case <-context.Done():
		err := context.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(writer, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
