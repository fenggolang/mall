package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r:=mux.NewRouter()
	// 静态文件服务器：http://ip:port/doc/查看项目doc目录下的文件
	r.Handle("/doc/",http.StripPrefix("/doc/",http.FileServer(http.Dir("./doc"))))
	http.Handle("/",r)
	//http.ListenAndServe(":8080",r)
	http.ListenAndServe(":8080",nil)
}

