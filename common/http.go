package common

import (
	"log"
	"net/http"
	"time"
)

func StartHttpServer(listenPort string) {
	mux := http.NewServeMux()
	mux.Handle("/*", &myHandler{})
	server := &http.Server{
		Addr:         ":" + listenPort,
		WriteTimeout: time.Second * 3, //设置3秒的写超时
		Handler:      mux,
	}
	log.Fatal(server.ListenAndServe())
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO 分析报文
	w.WriteHeader(404)
}
