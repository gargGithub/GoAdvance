package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"path/filepath"
)

func main() {

	http.HandleFunc("/files",UploadFile)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func UploadFile(w http.ResponseWriter, req *http.Request){
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost{
		f,h,err:=req.FormFile("q")
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}

		defer f.Close()
		fmt.Println("\nFile Name: ",f,"\nHeader: ",h,"\nError: ",err)

		bs,err:=ioutil.ReadAll(f)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		s = string(bs)

		dst,err:=os.Create(filepath.Join("./Uploaded Files/",	h.Filename))
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}


		_, err =dst.Write(bs)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}


	}
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`
		<form method = "POST" enctype = "multipart/form-data">
		Choose File: <input type = "file" name = "q">
		<input type = "submit" value = "Submit">
		</form>
		<br>`+s)


}