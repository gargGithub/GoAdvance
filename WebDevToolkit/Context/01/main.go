package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}


func foo(w http.ResponseWriter, req *http.Request){

	ctx:=req.Context()
	fmt.Println(ctx)
	fmt.Fprintln(w,ctx)

}