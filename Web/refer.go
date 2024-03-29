package main

import(
	"net/http"
	"fmt"
	"io/ioutil"
)

type Refer struct{
	handler http.Handler
	refer string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Referer() == this.refer{
		this.handler.ServeHTTP(w,r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r  *http.Request){
	w.Write([]byte("this is handler"))
}

func hellow(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hellow"))
}

func main(){
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer: "www.zxb.com",
	}
	http.HandleFunc("/hello",hellow)
	http.ListenAndServe(":8080",referer)
}