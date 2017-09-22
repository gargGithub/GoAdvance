package main

import (
	"net/http"
	"io"
	"os"
)

func batman(w http.ResponseWriter, req *http.Request){

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w,`

	<img src = "/batman.jpg">

	`)
}


func batmanHandle(w http.ResponseWriter, req *http.Request){

	f,err:=os.Open("batman.jpg")
	if err!=nil{
		http.Error(w,"File Not Found",404)
		return
	}
	defer f.Close()

	fi,err:=f.Stat()
	if err!=nil{
		http.Error(w,"File Not Found",404)
	}

	http.ServeContent(w,req,fi.Name(),fi.ModTime(),f)
	//http.ServeFile(w,req,fi.Name())


}

func main() {

	http.HandleFunc("/",batman)
	http.HandleFunc("/batman.jpg",batmanHandle)
	http.ListenAndServe(":8080",nil)
}
