package main

import (
	"net/http"
	"fmt"
)

type value int

func (m value) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"code")
}
func main() {
 var d value

 http.ListenAndServe(":8080",d)
}
