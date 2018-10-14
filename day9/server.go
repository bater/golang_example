package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<H1>Hello, World!</H2>")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
