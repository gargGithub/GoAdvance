package main

import (
 "html/template"

 "os"
 "log"
)

func main() {
 tpl,err:=template.ParseFiles("C:/Users/CUB/GoglandProjects/GoAdvance/Templates/sampleParsing.gohtml")
 if(err!=nil){
   log.Fatal(err)
 }

 err = tpl.Execute(os.Stdout,nil)
 if(err!=nil){
  log.Fatal(err)
 }
}
