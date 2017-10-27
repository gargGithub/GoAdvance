package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	s:="In word processing and desktop publishing, a hard return or paragraph break indicates a new paragraph," +
		" to be distinguished from the soft return at the end of a line internal to a paragraph. This distinction " +
		"allows word wrap to automatically re-flow text as it is edited, without losing paragraph breaks. The " +
		"software may apply vertical whitespace or indenting at paragraph breaks, depending on the selected style."

	//encodeStd:="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/" //to specify our own encoding standard

//	s64:=base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

//to use std. encoding

  s64:=base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)

	bs,err:=base64.StdEncoding.DecodeString(s64)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(string(bs))


}
