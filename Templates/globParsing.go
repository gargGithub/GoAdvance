package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tpl,err:= template.ParseGlob("Templates/*.gohtml")
	if(err!=nil){
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "sampleParsing.gohtml","sampleParsing")

	if(err!=nil){
		log.Fatal(err)
	}




}
