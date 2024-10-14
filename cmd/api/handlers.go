package main

import (
	"fmt"
	"net/http"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello, world")
}
