package main

import (
	"text/template"
	"os"
	"log"
	"strings"
)


var fm = template.FuncMap{

	"uc": strings.ToUpper,
	"firstThree": firstThree,
}


func firstThree(s string) string{
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
	}



var tpl *template.Template


func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("Templates/functionsParsing.gohtml"))
}

type firstname struct {
	Fname string
}

type lastname struct {
	Lname string
}

type fullname struct{
	Fnames []firstname
	Lnames []lastname
}
func main() {
	fname1:=firstname{
		"Shubham",
	}
	fname2:=firstname{
		"Aman",
	}
	fname3:=firstname{
		"Kiran",
	}
	lname1:=lastname{
		"Garg",
	}
	lname2:=lastname{
		"Patel",
	}
	lname3:=lastname{
		"Kannade",
	}


	Fnames:=[]firstname{fname1,fname2,fname3}
	Lnames:=[]lastname{lname1,lname2,lname3}


	fullnames:=fullname{
		Fnames:Fnames,
		Lnames:Lnames,
	}


	err:= tpl.ExecuteTemplate(os.Stdout,"functionsParsing.gohtml",fullnames)
	if(err!=nil){
		log.Fatal(err)
	}
}
