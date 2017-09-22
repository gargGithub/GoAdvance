package main

import
(
	"bufio"
	"fmt"
	"net"
	"log"

)

func handle(conn net.Conn){
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
	}

	defer conn.Close()
	fmt.Println("this will never print")
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


		go handle(conn)
	}


}


