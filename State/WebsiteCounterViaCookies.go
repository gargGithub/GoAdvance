package main

import (
	"net/http"
	"strconv"
	"io"
)

func main() {
	http.HandleFunc("/info",Counter)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}


func Counter(w http.ResponseWriter, req *http.Request){
	cookie,err:= req.Cookie("counter-cookie")
	if err == http.ErrNoCookie{
		cookie = &http.Cookie{
			Name: "counter-cookie",
			Value: "0",
		}
	}

	count,err:=strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w,cookie)
	io.WriteString(w,cookie.Value)

}