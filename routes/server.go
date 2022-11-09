package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var server *http.Server

// 启动服务
func HttpServerRun() {
	engine := InitRouter()
	addr := fmt.Sprintf(":%d", 9000)
	server = &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun%s\n", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun err:%v\n", err)
		}
	}()
}

// 停止服务
func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
