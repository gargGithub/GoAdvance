package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"io"
)

func main() {


	c:=getcode("test@example.com")
	fmt.Println(c)
	c=getcode("test@exampl.com")
	fmt.Println(c)

}

func getcode(s string) string{
	h:=hmac.New(sha256.New,[]byte("our-key"))
	io.WriteString(h,s)
	return fmt.Sprintf("%x",h.Sum(nil))

}
