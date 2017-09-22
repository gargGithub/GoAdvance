package main

import (
	"text/template"
	"os"
	"log"
)

type course struct{

	Number int
	Name string
}

type semester struct {
	Term string
	Courses []course
}

type year struct {
	Fall, Spring ,Summer semester
}

var tpl *template.Template

func init(){
	tpl =template.Must(template.New("").ParseFiles("Templates/compostion.gohtml"))
}




func main() {


	y:=year{
		Fall:semester{
			Term:"Fall",
			Courses:[]course{
				{1, "introduction to go"},
				{2,"Go web"},

			},

		},
		Spring:semester{

			Term:"Spring",
			Courses:[]course{

				{3, "Mobile apps using Go"},
				{4,"go advance"},
				},
			},

		}


		err:=tpl.ExecuteTemplate(os.Stdout,"compostion.gohtml",y)
		if err!=nil{
			log.Fatal(err)
	}

}
