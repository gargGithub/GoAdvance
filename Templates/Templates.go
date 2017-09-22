package main

import "fmt"

func main() {
   name:="shubham"

	hcode:=`
	<!DOCTYPE html>
	<html lane = "en">
	<head><meta chatset = "UTF-8">
	this will print the name of the user!</head>
	<body>
	<h1>`+name+`</h1>
	</body>
	</html>
	`
fmt.Println(hcode)

}
