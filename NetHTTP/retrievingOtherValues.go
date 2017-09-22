package main

import (
	"net/http"
	"log"
	"html/template"
	"net/url"
)

type hotdog int

func(num hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){

	err:= req.ParseForm()
	if err!=nil{
		log.Fatal(err)
	}

	data:= struct {
		Method string
		Submissions url.Values
		URL  *url.URL
		Header http.Header
		Host string
		ContentLength int64
	}{
		req.Method,
		req.Form,
		req.URL,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w,"index1.gohtml",data)

}
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index1.gohtml"))
}

func main() {
 var d hotdog
 http.ListenAndServe(":8080",d)

}
