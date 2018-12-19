package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	r:=mux.NewRouter()
	r.HandleFunc("/products/{key}",ProductHandle)
	r.HandleFunc("/articles/{category}/",ArticlesCategoryHandle)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}",ArticleHandle)
	// 常用的是：在一条路线上组合多个匹配器.
	r.HandleFunc("/products",ProductsHandle).Methods("GET").Schemes("http")
	//r.HandleFunc("/products",ProductsHandle)

	http.ListenAndServe(":8080",r)
}

func ProductHandle(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("请求路径是："+req.URL.Path))
}

func ProductsHandle(w http.ResponseWriter,req *http.Request) {
	w.Write([]byte("请求路径是："+req.URL.Path))
}
func ArticlesCategoryHandle(w http.ResponseWriter,req *http.Request) {
	vars:=mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Category: %v\n",vars["category"])
}

func ArticleHandle(w http.ResponseWriter,req *http.Request) {
	w.Write([]byte("请求路径是："+req.URL.Path))
}