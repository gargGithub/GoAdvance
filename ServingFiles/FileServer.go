package main

import (
	"net/http"
	"io"
)

func d(w http.ResponseWriter, req* http.Request){

	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`<img src="batman.jpg">`)
}


func main() {
	http.HandleFunc("/d",d)
	http.Handle("/",http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080",nil)
}
