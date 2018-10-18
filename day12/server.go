package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Fprint(w, "Url Param 'key' is missing")
		return
	}

	key := keys[0]
	fmt.Fprint(w, "Url Param 'key' is: "+string(key))
}
