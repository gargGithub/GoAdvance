package main

import (
	"net/http"
	"context"
	"fmt"
	"time"
)

func main() {
  // to time out something

  http.HandleFunc("/",foo)
  http.HandleFunc("/bar",bar)
  http.Handle("/favicon.ico",http.NotFoundHandler())
  http.ListenAndServe(":8080",nil)

}

func foo(w http.ResponseWriter, req *http.Request){
	ctx:=req.Context()

	ctx = context.WithValue(ctx,"userID",777)
	ctx = context.WithValue(ctx,"fname", "shubham garg")


	results,err:= dbAccess(ctx)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusRequestTimeout)
	}

	fmt.Fprint(w,results)

}

func dbAccess(ctx context.Context) (int ,error){

	ctx,cancel:=context.WithTimeout(ctx,1*time.Second)

	defer cancel()

	ch:=make(chan int)

	go func(){
		//long running task

		uid:=ctx.Value("userID").(int)
		time.Sleep(4*time.Second)

		//check to make sure we are not running in vain
		//if ctx.Done() has

		if ctx.Err() !=nil{
			return
		}
		ch<-uid
	}()

	select {
	  case <-ctx.Done():
	  	return 0, ctx.Err()

	case i:=<-ch:
		return i,nil
		
	}
}

func bar(w http.ResponseWriter, req *http.Request){

	ctx:=req.Context()
	fmt.Println(ctx)
	fmt.Fprint(w,ctx)

}


//output: context deadline exceeded
//		  0