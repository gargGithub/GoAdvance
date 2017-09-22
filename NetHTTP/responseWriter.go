package main

import (
	"net/http"
	"fmt"
)

type val int

func (m val) ServeHTTP(w http.ResponseWriter, req *http.Request){

	w.Header().Set("Garg-Key","this is from shubham")
	w.Header().Set("Content-Type","text/html	; charset=utf-8")
	fmt.Fprintln(w,"<h1>any code you want here</h1>")
}

func main() {
	var d val
	http.ListenAndServe(":8080",d)

}
