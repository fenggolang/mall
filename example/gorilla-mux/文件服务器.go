package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r:=mux.NewRouter()

	// 静态文件服务器：http://ip:port/doc/查看项目doc目录下的文件
	//r.Handle("/doc/",http.StripPrefix("/doc/",http.FileServer(http.Dir("./doc")))) // 和下面一行看上去等价(访问http://localhost:8080/doc/都可以列出目录下的文件，但是这行点击文件查看的时候却返回404 page not found)
	r.PathPrefix("/doc/").Handler(http.StripPrefix("/doc/",http.FileServer(http.Dir("./doc"))))

	http.ListenAndServe(":8080",r)

	//http.Handle("/",r) // 如果把r注册到go原生mux里面，那么http.ListenAndServe的第二个参数handler的值就传nil(内部默认也是defaultServeMux)
	//http.ListenAndServe(":8080",nil)
}