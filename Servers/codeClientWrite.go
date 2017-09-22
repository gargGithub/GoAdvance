package main

import (
	"net"
	"log"
	"fmt"
)

func main() {
	conn,err:=net.Dial("tcp","localhost:8080")
	if(err!=nil){
		log.Fatal(err)
	}

	defer conn.Close()

fmt.Fprintln(conn,"I dialed you")


}
