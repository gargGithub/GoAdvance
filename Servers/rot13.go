package main

import (
	"net"
	"log"
	"bufio"
	"strings"
	"fmt"
)

func handleRequest(conn net.Conn){
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan(){
		line:=strings.ToLower(scanner.Text())
		bs:=[]byte(line)
		r:=rot13(bs)
		fmt.Fprintf(conn,"%s -- %s",line,r)
	}
	defer conn.Close()
}

func rot13(bs []byte)[]byte{
	var r13 = make([]byte,len(bs))
	for i,v:=range bs{
		if(v<=109){
			r13[i] = v +13
		}else{
			r13[i]=v -13
		}
	}
	return r13
}

func main() {

	li,err:=net.Listen("tcp",":8080")
	if(err!=nil){
		log.Fatal(err)
	}
	defer li.Close()

	for{
		conn,err:=li.Accept()
		if(err!=nil){
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}
