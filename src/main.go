package main

import "net/http"


func index(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// 创建一个多路复用器
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	// 使用 StripPrefix 函数去除请求 URL 中的指定前缀
	mux.Handle("/static", http.StripPrefix("/static", files))

	/**
		因为所有处理器都接受一个 ResponseWriter 和一个指向 Request 结构的指针作为参数，
		并且所有请求参数都可以通过访问 Request 结构体得到，所以程序并不需要向处理器显示地传入任何请求参数。
	 */
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
