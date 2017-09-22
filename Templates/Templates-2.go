package main

import (
	"fmt"

	"log"
	"io"
	"strings"

	// "path/filepath"
	"os"
)

func main() {
	name:="shubham"

	hcode:=fmt.Sprint(`
	<!DOCTYPE html>
	<html lang = "en">
	<head><meta charset = "UTF-8">
	this will print the name of the user!</head>
	<body>
	<h1>`+name+`</h1>
	</body>
	</html>
	`)
	fmt.Println(hcode)

	//nf,_:=filepath.Abs("C:/Users/CUB/GoglandProjects/GoAdvance/Templates/index2.html")
      fp,err:=os.Create("C:/Users/CUB/GoglandProjects/GoAdvance/Templates/index2.html")
	if(err!=nil){
		log.Fatal("Error creating file.",err)

	}else{

		io.Copy(fp,strings.NewReader(hcode))
	}

}
