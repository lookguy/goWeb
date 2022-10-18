package main

import(
	"fmt"
)

func main()  {
	// example 1
	var score int = 100
	var name string = "Barry"
	// %p 可输出用 & 连接变量的地址指针值
	fmt.Printf("%p %p", &score, &name)

}