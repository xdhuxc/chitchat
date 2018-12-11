package main

import "net/http"

func main() {
	// 创建一个多路复用器
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
