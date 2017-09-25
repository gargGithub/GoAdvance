package main

import (
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/forms",forms)
	http.Handle("favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)


}

func forms(w http.ResponseWriter, req *http.Request){
	v:=req.FormValue("name")
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`
	<form method = "post">
	<input type = "text" name="name">
	<input type = "submit" value="click here">
	</form>
	<br>
	`+"The name is:"+v)
}
