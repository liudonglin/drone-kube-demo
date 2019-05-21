package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("server start up")
	handle := &myHttpHandle{}
	http.Handle("/ci/", handle)
	http.ListenAndServe(":8000", nil)
}

func hello() string {
	return "hello world"
}
