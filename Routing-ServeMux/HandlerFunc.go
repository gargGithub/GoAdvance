package main

import (
	"net/http"
	"io"

)



func d(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"this is dog url")
}


func c(w http.ResponseWriter, req *http.Request){
	io.WriteString(w,"this is cat url")
}

func main() {




	http.Handle("/dog/",http.HandlerFunc(d))
	http.Handle("/cat",http.HandlerFunc(c))
	http.ListenAndServe(":8080",nil)

}
