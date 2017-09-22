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


	names:=map[int]string{
		1:"shubham",
		2: "mukul",
		3: "aman",
	}

	err:=tpl.Execute(os.Stdout,names)
	if(err!=nil){
		log.Fatal(err)
	}

}
