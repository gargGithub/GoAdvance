package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

func main() {
	http.HandleFunc("/session",sessions)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func sessions(w http.ResponseWriter, req *http.Request){
	cookie,err:=req.Cookie("session-cookie")
	if err!=nil{
		id:=uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-cookie",
			Value: id.String(),
			HttpOnly:true,
		}
		http.SetCookie(w,cookie)
	}
	fmt.Println(cookie)
}