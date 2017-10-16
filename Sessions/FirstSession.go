package main

import  (
	"net/http"
	"html/template"
	"github.com/satori/go.uuid"
)
type User struct {
   UserName string
   FirstName string
   LastName string
}

var tpl *template.Template
var DbUsers =map[string]User{}  //UserID and User
var DbSessions =map[string]string{}//Session ID and UserID

func init(){
	tpl = template.Must(template.ParseGlob("Templates/*"))
}


func main() {
	http.HandleFunc("/panel",index)
	http.HandleFunc("/portal",portal)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request){
   c,err:=req.Cookie("session")
   if err!=nil{
   	  sID := uuid.NewV4()
   	  c = &http.Cookie{
   	  	Name: "session",
   	  	Value: sID.String(),
	  }

	  http.SetCookie(w,c)
   }

   //If the user exists already get the user
	var users User
	if userid,ok:= DbSessions[c.Value];ok{
		users = DbUsers[userid]
	}
	//process form submission

	if req.Method == http.MethodPost{
		userid:= req.FormValue("username")
		fname:=req.FormValue("firstname")
		lname:=req.FormValue("lastname")
		users = User{userid,fname,lname}
		DbSessions[c.Value] = userid
		DbUsers[userid] = users

	}

	tpl.ExecuteTemplate(w,"index.gohtml",users)
}

func portal(w http.ResponseWriter, req *http.Request){
	c,err:=req.Cookie("session")
	if err!=nil{
		http.Redirect(w,req,"/panel",http.StatusSeeOther)
		return
	}
	userid,ok:=DbSessions[c.Value]
	if !ok{
		http.Redirect(w,req,"/panel",http.StatusSeeOther)
	}
	users:=DbUsers[userid]
    tpl.ExecuteTemplate(w,"panel.gohtml",users)
}