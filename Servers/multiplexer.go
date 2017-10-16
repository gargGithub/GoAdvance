package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	li,err:=net.Listen("tcp",":8080")
	if err!=nil {
		log.Fatal(err)
	}

	defer li.Close()
	for{
		conn,err:=li.Accept()
		if err!=nil {
			log.Fatal(err)
		}

		go handleMux(conn)
	}
}


func handleMux(conn net.Conn){
	defer conn.Close()
	request(conn)
}


func request(conn net.Conn){
	i:=0
	scanner:=bufio.NewScanner(conn)
	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)

		if i==0 {
			mux(conn, line)
		}
		if line ==""{
			break
		}
		i++
	}
}


func mux(conn net.Conn, line string){
	m:=strings.Fields(line)[0]
	u:=strings.Fields(line)[1]
	fmt.Println("******METHOD",m)
	fmt.Println("*******URI",u)

	if m=="GET" && u=="/"{
		index(conn)
	}
	if m=="GET" && u=="/about"{
		about(conn)
	}
	if m=="GET" && u=="/contact"{
		contact(conn)
	}
	if m=="GET" && u=="/apply"{
		apply(conn)
	}
	if m=="POST" && u=="/apply"{
		applyProcess(conn)
	}
}

func index(conn net.Conn){
	body:=`<html><head>INDEX<br><title>index.gohtml page</title></head><body>
	<a href = "/">Index</a><br>
	<a href = "/about">About</a><br>
	<a href = "/contact">Contact</a><br>
	<a href = "/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}

func about(conn net.Conn){
	body:=`<html><head>ABOUT<br><title>about page</title></head><body>
	<a href = "/">Index</a><br>
	<a href = "/about">About</a><br>
	<a href = "/contact">Contact</a><br>
	<a href = "/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}

func contact(conn net.Conn){
	body:=`<html><head>CONTACT<br><title>contact page</title></head><body>
	<a href = "/">Index</a><br>
	<a href = "/about">About</a><br>
	<a href = "/contact">Contact</a><br>
	<a href = "/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}

func apply(conn net.Conn){
	body:=`<html><head>APPLY<br><title>apply page</title></head><body>
	<a href = "/">Index</a><br>
	<a href = "/about">About</a><br>
	<a href = "/contact">Contact</a><br>
	<a href = "/apply">Apply</a><br>
	<form action = "/apply" method = "post">
	<input type="submit" value="apply"></form>
	</body>
	</html>`
	fmt.Fprint(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}

func applyProcess(conn net.Conn){
	body:=`<html><head>APPLY PROCESS<br><title>apply page</title></head><body>
	<a href = "/">Index</a><br>
	<a href = "/about">About</a><br>
	<a href = "/contact">Contact</a><br>
	<a href = "/apply">Apply</a><br>
	</body>
	</html>`
	fmt.Fprint(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content length: %d\r\n",len(body))
	fmt.Fprint(conn,"Content-Type: text/html\r\n")
	fmt.Fprint(conn,"\r\n")
	fmt.Fprint(conn,body)
}