package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"

)

func main() {

	rp, err := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")
	if err != nil {
		fmt.Println(err)
	}
	defer rp.Body.Close()




	a, err1 := ioutil.ReadAll(rp.Body)
	if err != nil {
		fmt.Println(err1)
	}





	words := string(a)

	list:= strings.Split(words,"\n")



    var new string
	var str string

	  fmt.Println("Enter the word you want to search for: ")
	  fmt.Scanln(&str)
 fmt.Println()
 fmt.Println()




	  for{
		  fmt.Println("Search Suggestions: ")
	  for i, v := range list {

		  if (strings.HasPrefix(v, str)) {
			  fmt.Println(i," : ",v)
		  }
	  }
	  fmt.Println()
	  fmt.Println("Want to search more (Y/N) :")
	  var ch string
	  fmt.Scanln(&ch)

	  if (ch == "y"|| ch == "Y") {
	  	fmt.Print(str)
	  	fmt.Scanf("%v",&new)
	  	str = str+new
	  } else{
          break
	  }
  }

  fmt.Println("Choose index.gohtml of your required word: ")
  var index int
  fmt.Scanln(&index)
  fmt.Print("You have choosen: ")
  fmt.Println(list[index])




}