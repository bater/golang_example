package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/bool/", isHappy)
	http.HandleFunc("/num/", add)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func add(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.URL.Path, "/")
	_, _, s := a[0], a[1], a[2]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(w, "n + 100 = %d\n", i+100)
	}
}

func isHappy(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.URL.Path, "=")
	_, b := a[0], a[1] == "true"
	if b {
		fmt.Fprint(w, "I am a Happy man")
	} else {
		fmt.Fprint(w, "I am not a Happy man")
	}
}
