package main

import (
	"net/http"
	"io"
)

func d(w http.ResponseWriter, req *http.Request){

	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`<img src = "/resources/batman.jpg">`)
}

func main() {
	http.HandleFunc("/",d)
	http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("./StripPrefix"))))
	http.ListenAndServe(":8080",nil)
}
