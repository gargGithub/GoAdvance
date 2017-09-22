package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

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
		io.WriteString(conn,"Hello from tcp server")
		fmt.Fprintln(conn,"hi tcp server")
		fmt.Fprintf(conn,"%v","bye bye")

     conn.Close()

	}


}
