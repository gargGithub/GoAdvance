package main

import (
	"net/http"
	"log"
	"io"

)

func main() {

	http.HandleFunc("/foo",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080",nil))

}

func foo(w http.ResponseWriter, req *http.Request){
	v:=req.FormValue("q")
	io.WriteString(w,"the query parameter is: "+v)
}
//VISIT THIS PAGE
//localhost:8080/foo/?q=dog
//output: the query parameter is:dog