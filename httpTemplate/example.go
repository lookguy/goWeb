package main

import(
	"fmt"
	"html/template"
	"net/http"
)

type ObjInfo struct{
	Name string
	do string
	adj string
	date int
}

func renderExample(w http.ResponseWriter, r *http.Request){
	// 解析
	// 直接用编辑器的终端会出现，使用系统的powershell可以避免
	t,err := template.ParseFiles("example.tmpl")
	if err!= nil{
		fmt.Println("template parsefile fail,err:",err)
		return
	}
	//  渲染模板
	obj := ObjInfo {
		Name: "招雄彬",
		do: "学GO",
		adj: "很累，刚跑完",
		date: 3,
	}

	t.Execute(w,obj)
}

func main(){
	http.HandleFunc("/",renderExample)
	http.ListenAndServe(":8080",nil)
}