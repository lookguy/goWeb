package main

import(
	"fmt"
	"time"
)

func HelloWorld()  {
	fmt.Println("this is goroutine msg")
}

func main(){
	go HelloWorld()
	fmt.Println("start")
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

