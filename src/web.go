package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>Welcome to home page</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	fmt.Println("Server Starting -------")
	http.ListenAndServe(":3000", nil)
}
