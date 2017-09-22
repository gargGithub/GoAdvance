package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func handleRequestResponse(conn net.Conn){
	defer conn.Close()

	request(conn)

	respond(conn)
}


func request(conn net.Conn){
	i:=0
	scanner:= bufio.NewScanner(conn)

	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
		if i==0{
			m:=strings.Fields(line)[0]
			u:=strings.Fields(line)[1]
			fmt.Println("METHOD********",m)
			fmt.Println("URI********",u)
		}
		if line == ""{
			break
		}
		i++

	}
}

func respond(conn net.Conn){
	body:=`<html><head><title>hello world</title></head><body><h1>hello world</body></html>`
	fmt.Fprint(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}

func main() {


	li,err:=net.Listen("tcp",":8080")
	if(err!=nil){
		log.Fatal(err)
	}

	defer li.Close()
	for{
		conn,err:= li.Accept()
		if(err!=nil){
			log.Fatal(err)
		}

		go handleRequestResponse(conn)
	}
}
