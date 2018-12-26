package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
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

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	Hello := HelloHandler{}

	World := WorldHandler{}

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	/**
	如果被绑定的 URL 不是以 / 结尾，那么它只会与完全相同的 URL 匹配；
	但如果被绑定的 URL 以 / 结尾，那么即使请求的 URL 只有前缀部分与被绑定 URL 相同，ServeMux 也会认定这两个 URL 是匹配的。
	*/

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

	/**
	函数链
	*/
	http.HandleFunc("/chain", log(hello))

	server.ListenAndServe()

}
