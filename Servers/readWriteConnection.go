package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func handleWriting(conn net.Conn){
	scanner:= bufio.NewScanner(conn)

	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn,"I heard you say: %s\n", line)
	}
	defer conn.Close()
	fmt.Println("End point reached")


}


func main() {

li,err:= net.Listen("tcp",":8080")
if(err!=nil){
	log.Fatal(err)
}
defer li.Close()

for{
	conn,err:=li.Accept()
	if(err!=nil){
		log.Fatal(err)
	}
	go handleWriting(conn)
}


}
