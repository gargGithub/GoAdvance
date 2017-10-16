package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
	"fmt"
)

type User struct {
	Username string
	Password []byte
	FirstName string
	LastName string
	Role string
}


type Session struct {
	Un string
	LastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]User{}
var dbSessions = map[string]Session{}
var dbSessionsCleaned time.Time
const sessionLength int = 30

func init(){
	tpl = template.Must(template.ParseGlob("Templates/*.html"))
	dbSessionsCleaned = time.Now()
}



func index(w http.ResponseWriter, req *http.Request){
	u:=getUser(w,req)
	tpl.ExecuteTemplate(w,"index.html",u)
}



func bar(w http.ResponseWriter, req *http.Request){
	u:=getUser(w,req)
	if !alreadyLoggedIn(w,req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	if u.Role!="007"{
		http.Error(w,"You must be 007 to enter the room",http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(w,"bar.html",u)
}



func signup(w http.ResponseWriter, req *http.Request){
	if alreadyLoggedIn(w,req){
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	//process form submission
	if req.Method == http.MethodPost{
		un:=req.FormValue("username")
		pwd:=req.FormValue("password")
		fn:=req.FormValue("firstname")
		ln:=req.FormValue("lastname")
		rl:=req.FormValue("role")


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

		dbSessions[c.Value] = Session{un,time.Now()}

		//store user in dbUsers
		bs,err:=bcrypt.GenerateFromPassword([]byte(pwd),bcrypt.MinCost)
		if err!=nil{
			http.Error(w,"Internal Server Error",http.StatusInternalServerError)
			return
		}

		u := User{un,bs,fn,ln,rl}
		dbUsers[un] = u

		//redirect
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"signup.html",nil)

}

func login(w http.ResponseWriter,req *http.Request){

	if alreadyLoggedIn(w,req){
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

		dbSessions[c.Value] = Session{un,time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"login.html",nil)
}

func logout(w http.ResponseWriter, req *http.Request){
	if !alreadyLoggedIn(w,req){
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

	//clean dbSessions
	if time.Now().Sub(dbSessionsCleaned)>(time.Second*30){
		go cleanSessions()
	}

	http.Redirect(w,req,"/login",http.StatusSeeOther)

}

func cleanSessions(){
	fmt.Println("BEFORE CLEAN")//for demonstration purpose
	showSessions()//for demonstration purpose
	for k,v:=range dbSessions{
		if time.Now().Sub(v.LastActivity)>(time.Second*30){
			delete(dbSessions,k)
		}
	}

	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")////for demonstration purpose
	showSessions()//for demonstration purpose
}

//for demonstration purpose
func showSessions()  {
	fmt.Println("**************")
	for k,v:=range dbSessions{
		fmt.Println(k,v)
	}
	fmt.Println("")
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