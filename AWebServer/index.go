package main

import (
	"net/http"
	"fmt"
)
// 定义多个处理器
type handle1 struct{}

func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hi,handle 1")
}

type handle2 struct{}

func (h2 *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hi,handle 2")
}

func main(){
	handle1 := handle1{}
	handle2 := handle2{}
	mux := http.NewServeMux()
	
	server := http.Server{
		Addr: "0.0.0.0:8085",
		Handler: mux,
	}
	mux.Handle("/handle1", &handle1)
	mux.Handle("/handle2", &handle2)
	server.ListenAndServe()
}

