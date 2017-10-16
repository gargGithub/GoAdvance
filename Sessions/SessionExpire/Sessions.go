package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) User{

	c,err:= req.Cookie("session")
	if err!=nil{
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value:sID.String(),
		}

	}
	c.MaxAge = sessionLength

	http.SetCookie(w,c)
	//if the user already exists
	var u User
	if s,ok:=dbSessions[c.Value];ok{
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.Un]
	}

   return u
}

func alreadyLoggedIn(w http.ResponseWriter,req *http.Request) bool{
 	c,err:=req.Cookie("session")
 	if err!=nil{
 		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.Un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
 }