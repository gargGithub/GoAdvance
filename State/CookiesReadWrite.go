package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}


func set(w http.ResponseWriter, req *http.Request){

	http.SetCookie(w,&http.Cookie{
		Name: "my-cookie",
		Value:"12345",
	})
fmt.Fprintln(w,"COOKIE WRITTEN - CHECK YOUR BROWSWER")
fmt.Fprintln(w,"in chrome goto dev tools->application->cookies")
}

func read(w http.ResponseWriter, req *http.Request){

	c,err:=req.Cookie("my-cookie")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusNotFound)
	}

	fmt.Fprintln(w,"YOUR COOOKIE: ",c)
}