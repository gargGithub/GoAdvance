package main

import (
	"net/http"
	"fmt"
)

func dog(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.URL)
	fmt.Fprintln(w,"look at your terminal")
}

func main() {


	http.HandleFunc("/",dog)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}
