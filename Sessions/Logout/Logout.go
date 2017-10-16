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
	bs,_ := bcrypt.GenerateFromPassword([]byte("password"),bcrypt.MinCost)
	dbUsers["test@test.com"] = User{"test@test.com",bs,"shubham","garg"}
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

func login(w http.ResponseWriter,req *http.Request){

	if alreadyLoggedIn(req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}
	//process form submission
	if req.Method==http.MethodPost {
		un := req.FormValue("username")
		pwd := req.FormValue("password")

		//is there a username

		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or Password do not match!", http.StatusForbidden)
			return
		}

		//does entered password match the stored passsword

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pwd))
		if err != nil {
			http.Error(w, "Username and/or Password do not match!", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		dbSessions[c.Value] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"login.html",nil)
}

func logout(w http.ResponseWriter, req *http.Request){
	if !alreadyLoggedIn(req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	c,_:= req.Cookie("session")
	//delete session
	delete(dbSessions,c.Value)
	//removing cookie
	c = &http.Cookie{
		Name: "session",
		Value: " ",
		MaxAge: -1,
	}

	http.SetCookie(w,c)

	http.Redirect(w,req,"/login",http.StatusSeeOther)

}

func main(){

	http.HandleFunc("/",index)
	http.HandleFunc("/bar",bar)
	http.HandleFunc("/signup",signup)
	http.HandleFunc("/login",login)
	http.HandleFunc("/logout",logout)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}
