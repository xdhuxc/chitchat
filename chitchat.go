package main

import (
	"fmt"
	"net/http"
)

/**
处理器：
在 Go 语言中，一个处理器就是一个拥有 ServeHTTP 方法的接口，这个 ServeHTTP 方法需要接收两个参数：
第一个参数是一个 ResponseWriter 接口，而第二个参数则是一个指向 Request 结构体的指针。
*/

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

/**
处理器函数：
处理器函数实际上就是与处理器拥有相同行为的函数，也就是与 ServeHTTP 方法拥有相同的签名，都接受 ResponseWriter 和 指向 Request 结构体的指针作为参数
*/
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func main() {
	Hello := HelloHandler{}

	World := WorldHandler{}

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	/**
	使用多个处理器对请求进行处理
	*/
	http.Handle("/Hello", &Hello)
	http.Handle("/World", &World)

	/**
	使用处理器函数对请求进行处理
	*/
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()

}
