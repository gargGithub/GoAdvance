package main

import (
	"net/http"
	"io"
)

func d(w http.ResponseWriter, req *http.Request){

	w.Header().Set("Content-Type","text/html; charset=utf-8")

	io.WriteString(w,`

	<img src = "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c5/Airplane_silhouette.svg/1024px-Airplane_silhouette.svg.png">

	`)
}

func main() {

	http.HandleFunc("/",d)
	http.ListenAndServe(":8080",nil)

}
