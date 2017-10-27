package main

import (
	"fmt"
	"time"
)

func main() {

	for n:= range gen(){

		fmt.Println(n)
		if n==5{      //although for n=5 loop will break but the goroutine is still running, so this will be resolved using context
			break
		}
	}

	time.Sleep(10*time.Second)

}

//gen is broken generator that will leak goroutine
func gen()<-chan int{
	ch := make(chan int)

	go func() {
		var n int
		for{
			ch<-n
			n++
		}
	}()

	return ch
}