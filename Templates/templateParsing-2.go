package main

import (
	"html/template"

	"os"
	"log"
	"fmt"
)

func main() {
	tpl,err:=template.ParseFiles("sampleParsing.gohtml")
	if(err!=nil){
		log.Fatal(err)
	}


	nf,err:= os.Create("index4.html")
	if(err!=nil){
		fmt.Println("Error Creating file",err)
	}
	defer nf.Close()


	err = tpl.Execute(nf,nil)
	if(err!=nil){
		log.Fatal(err)
	}
}
