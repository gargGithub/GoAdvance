package main

import (
	"encoding/json"
	"log"
	"fmt"

)
 type JSON struct {
	 Msg string `json:"msg,omitempty"`
	 Id int `json:"id,omitempty"`
 }
func stringInterfaceMap(){
	raw_data:= []byte(`{"msg":"Hello Go!","id":12345}`)
	decoded:=JSON{}
	err:=json.Unmarshal(raw_data,&decoded)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(decoded.Msg,decoded.Id)

	enc,_:=json.Marshal(&decoded)
	fmt.Println(string(enc))
}

func main() {

	stringInterfaceMap()
}
