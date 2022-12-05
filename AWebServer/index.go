package main

import (
	"net/http"
	"log"
	"fmt"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Go Web!")
}

func main(){
	http.HandleFunc("/hello",helloWorld)
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}

