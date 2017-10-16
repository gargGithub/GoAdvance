package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password []byte
	FirstName string
	LastName string
}

var tpl *template.Template
var dbUsers = map[string]User{}
var dbSessions = map[string]string{}

func init(){
	tpl = template.Must(template.ParseGlob("Templates/*.html"))
}


func index(w http.ResponseWriter, req *http.Request){
	u:=getUser(w,req)
	tpl.ExecuteTemplate(w,"index.html",u)
}

func bar(w http.ResponseWriter, req *http.Request){
	u:=getUser(w,req)
	if !alreadyLoggedIn(req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w,"bar.html",u)
}

func signup(w http.ResponseWriter, req *http.Request){
	if alreadyLoggedIn(req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	//process form submission
	if req.Method == http.MethodPost{
		un:=req.FormValue("username")
		pwd:=req.FormValue("password")
		fn:=req.FormValue("firstname")
		ln:=req.FormValue("lastname")


		//username already taken

		if _,ok:= dbUsers[un];ok{
			http.Error(w,"Username already exists",http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:"session",
			Value:sID.String(),
		}
		http.SetCookie(w,c)

		dbSessions[c.Value] = un

		//store user in dbUsers
		bs,err:=bcrypt.GenerateFromPassword([]byte(pwd),bcrypt.MinCost)
		if err!=nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return
		}

		u := User{un,bs,fn,ln}
		dbUsers[un] = u

		//redirect
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"signup.html",nil)

}


func main(){

	http.HandleFunc("/",index)
	http.HandleFunc("/bar",bar)
	http.HandleFunc("/signup",signup)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}
