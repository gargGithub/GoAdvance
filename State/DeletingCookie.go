package main

import (
	"net/http"
	"fmt"

)

func main() {

	http.HandleFunc("/",index)
	http.HandleFunc("/set",setCookie)
	http.HandleFunc("/read",readCookie)
	http.HandleFunc("/expire",expireCookie)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w, `<a href = "/set">Set Cookie</a>`)
}

func setCookie(w http.ResponseWriter, req *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name: "session",
		Value: "Any value",
	})
	fmt.Fprintln(w,`<br><a href ="/read">Read Cookie</a>`)
}

func readCookie(w http.ResponseWriter, req *http.Request){
	cookie,err:=req.Cookie("session")
	if err!=nil{
		http.Redirect(w,req,"/set",http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprintf(w,`<h2>Your Cookie: %v </h2><br><a href = "/expire">Expire Cookie</a>`,cookie)
}

func expireCookie(w http.ResponseWriter, req *http.Request){
	cookie,err:=req.Cookie("session")
	if err!=nil{
		http.Redirect(w,req,"/set",http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w,cookie)
	http.Redirect(w,req,"/",http.StatusSeeOther)

}