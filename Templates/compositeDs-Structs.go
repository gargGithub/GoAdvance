package main

import (
	"text/template"
	"os"
	"log"
)


var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("Templates/sliceParsing.gohtml"))
}

type names struct {
	Fname string
	Lname string
}
func main() {


	name1:=names{
		"shubham",
		"garg",
	}
	name2:=names{
		"aman",
		"patel",
	}
     names:=[]names{name1,name2}
	err:=tpl.Execute(os.Stdout,names)
	if(err!=nil){
		log.Fatal(err)
	}

}
