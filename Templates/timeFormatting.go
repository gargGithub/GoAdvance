package main

import (
	"text/template"
	"time"
	//	"fmt"
	"os"
	"log"
	//"fmt"
)

var tpl *template.Template

func init(){
	tpl =template.Must(template.New("").Funcs(fm).ParseFiles("Templates/timeParsing.gohtml"))
}

var fm =template.FuncMap{

	"datePrint":datePrint,

}

func datePrint(t time.Time) string{
	return t.Format("02 Jan 2006")

}

func main() {


	err := tpl.ExecuteTemplate(os.Stdout,"timeParsing.gohtml",time.Now())
	if err!=nil{
		log.Fatal(err)
	}
}
