package main

import (
	"net/http"
	"io"
)

type number int

func(d number) ServeHTTP(w http.ResponseWriter, req *http.Request){

	switch req.URL.Path{

	case "/dog":
		io.WriteString(w,"this is dog url")

	case "/cat":
		io.WriteString(w,"this is cat url")

	}
}

func main() {

	var d number
	http.ListenAndServe(":8080",d)
}
