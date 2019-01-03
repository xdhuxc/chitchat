package main

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/xdhuxc/chitchat/src/models"
	"github.com/xdhuxc/chitchat/src/routes"
	"github.com/xdhuxc/chitchat/src/utils"
)

var conf models.Configuration

func init() {
	conf = utils.LoadConfig("../../conf.test.yaml")

	/**
	设置日志输出格式，自带的只有两种格式：
	logrus.JSONFormatter{}
	logrus.JSONFormatter{}
	*/
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//
	logrus.SetOutput(os.Stdout)
	// 设置最低级别日志
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {

	utils.P("ChitChat", utils.Version(), "started at")

	// 创建一个多路复用器
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(conf.Static))
	// 使用 StripPrefix 函数去除请求 URL 中的指定前缀
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	/**
	因为所有处理器都接受一个 ResponseWriter 和一个指向 Request 结构的指针作为参数，
	并且所有请求参数都可以通过访问 Request 结构体得到，所以程序并不需要向处理器显示地传入任何请求参数。
	*/
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/error", routes.Error)

	//
	mux.HandleFunc("/login", routes.Login)
	mux.HandleFunc("/logout", routes.Logout)
	mux.HandleFunc("/signup", routes.SignUp)
	mux.HandleFunc("/signup_account", routes.SignUpAccount)
	mux.HandleFunc("/authenticate", routes.Authenticate)

	//
	mux.HandleFunc("/thread/new", routes.NewThread)
	mux.HandleFunc("/thread/create", routes.CreateThread)
	mux.HandleFunc("/thread/post", routes.PostThread)
	mux.HandleFunc("/thread/read", routes.ReadThread)

	server := &http.Server{
		Addr:           conf.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(conf.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(conf.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
