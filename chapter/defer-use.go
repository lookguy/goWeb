package main

import(
	"fmt"
)

func main(){
	deferCall()
	fmt.Println("C")
}

func deferCall(){
	defer func1()
	defer func2()
	fmt.Println("D")
}

func func1(){
	fmt.Println("A")
}

func func2(){
	fmt.Println("B")
}