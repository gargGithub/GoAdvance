package main

import (
	"html/template"
	"net/http"
	"log"
)

type hotdog int

func(num hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err:= req.ParseForm()
	if err!=nil{
		log.Fatal(err)
	}

	err1:= tpl.ExecuteTemplate(w,"index.gohtml.gohtml",req.Form)
	if(err1!=nil){
		log.Fatal(err1)
	}

}
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml.gohtml"))
}


func main() {
	var d hotdog
	http.ListenAndServe(":8080",d)

}
