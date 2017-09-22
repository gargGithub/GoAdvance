package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template


func init(){
	tpl = template.Must(template.ParseFiles("Templates/Struct-Slice-Struct.gohtml"))
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


	err:= tpl.Execute(os.Stdout,fullnames)
    if(err!=nil){
    	log.Fatal(err)
	}
}
