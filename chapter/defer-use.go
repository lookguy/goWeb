package main

import(
	"fmt"
	"github.com/lookguy/goWeb/blob/b9fde909d32d4690fc849188a1a0810078204d46/chapter/oob-use.go "
)

func main(){
	deferCall()
	fmt.Println("C")

	s := new (chapter.Student)
	s.SetName("Zxb")
	fmt.Println(s.GetName())
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