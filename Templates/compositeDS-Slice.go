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

func main() {


	names:=[]string{"Shubham","Mukul","Aman","Kiran"}

	err:=tpl.Execute(os.Stdout,names)
	if(err!=nil){
		log.Fatal(err)
	}

}
