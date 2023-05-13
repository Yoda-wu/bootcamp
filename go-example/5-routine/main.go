package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(j int) {
	fmt.Println("hello ", j)
}

func main() {
	// 快速打印hello world
	for i:=0 ; i < 5; i++ {
		go hello(i);
	}
	time.Sleep(1*time.Second)
}

func CalSquare(){
	src := make(chan int)
	dest := make(chan int ,3)
	go func ()  {
		defer close(src);
		for i:= 0; i < 10 ;i++ {
			src <- i;
		}
	}()

	go func() {
		defer close(dest)
		for i := range src {
			dest <- i*i
		}
	}()

	for i := range dest {
		fmt.Println(i);
	}
}

var (
	x int64
	lock sync.Mutex
)

func addWithLock(){
	for i := 0; i < 100 ; i++ {
		lock.Lock()
		x+=1
		lock.Unlock()
	}
}

func addWithoutLock(){
	for i := 0; i < 100 ; i++ {
		x+=1
	}
}