package main

import(
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Write([]byte("Index"))
}

func main(){
	router := httprouter.New()

	router.GET("/",Index)

	router.GET("/default",func(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
		w.Write([]byte("default get"))
	})

	router.POST("/default", func(w http.ResponseWriter,r *http.Request, _ httprouter.Params){
		w.Write([]byte("default POST"))
	})

	// 精确匹配
	// router.GET("/user/name",func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	// 	w.Write([]byte("user name:" + p.ByName("name")))
	// })

	// 匹配所有
	router.GET("/user/*name",func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		w.Write([]byte("user name:" + p.ByName("name")))
	})
	
	http.ListenAndServe(":8083",router)
}