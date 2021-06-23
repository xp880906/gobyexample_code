package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}

func headers(writer http.ResponseWriter, request *http.Request) {
	for name, headers := range request.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(writer, "%v: %v\n", name, h)
			if err != nil {
				return
			}
		}
	}
}

func hello(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "hello\n")
	if err != nil {
		return
	}
}
