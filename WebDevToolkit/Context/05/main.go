package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()  //make sure all paths cancel the context to avoid the content leak

	for n:=range gen(ctx){
		fmt.Println(n)
		if n==5{
			cancel()
			break
		}
	}
	time.Sleep(20*time.Second)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int

		for {
			select {
			case <-ctx.Done():
				return // avoid leaking of this goroutine when the ctx is done
			case ch<-n:
				n++
			}
		}
	}()
return ch
}