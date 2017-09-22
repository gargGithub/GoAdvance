package main

import (
	"net/http"
	"io"
)

type number int

func(num number) ServeHTTP(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"this is dog url")
}

type number2 int
func(num2 number2) ServeHTTP(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"this is cat url")
}

func main() {

	var d number
	var c number2

	mux:=http.NewServeMux()
	mux.Handle("/dog/",d)
	mux.Handle("/cat",c)
	http.ListenAndServe(":8080",mux)

}
