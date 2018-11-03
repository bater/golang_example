package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "username",
		Value:    "bater",
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "Cannot get cookie")
	}
	fmt.Fprintln(w, c)

	cs := r.Cookies()
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}
	http.HandleFunc("/cookie/set", setCookie)
	http.HandleFunc("/cookie/get", getCookie)
	server.ListenAndServe()

}
