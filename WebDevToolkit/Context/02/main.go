package main

import (
	"net/http"
	"context"
	"fmt"
)

func main() {

	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.Handle("favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, req *http.Request){
	ctx:=req.Context()
	ctx = context.WithValue(ctx,"userID",777)
	ctx = context.WithValue(ctx,"fname","James Bond")

	results:= dbAccess(ctx)

	fmt.Fprint(w,results)
}

func dbAccess(ctx context.Context) int{
	uid:=ctx.Value("userID").(int)
	return uid
}

func bar(w http.ResponseWriter, req *http.Request){
	ctx:=req.Context()
	fmt.Println(ctx)
	fmt.Fprintln(w,ctx)
}
