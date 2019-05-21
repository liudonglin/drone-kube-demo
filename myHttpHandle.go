package main

import (
	"net/http"
)

type myHttpHandle struct {
}

func (h *myHttpHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	val := []byte(`<h1>Administration Guide</h1> 
	If you are installing Drone we recommend starting with a basic installation. Once you are familiar with Drone, this section will help you setup a distributed, secure, production-ready instance.`)
	w.Write(val)
	return
}
