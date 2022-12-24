package initializer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"newsSearchEngine/engine"
	"newsSearchEngine/engine/tokenizer"
	"newsSearchEngine/global"
	"newsSearchEngine/web/router"
	"newsSearchEngine/web/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// init
func Initialize() {
	global.CONFIG = Parser()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %s\n", r)
		}
	}()

	tokenizer := NewTokenizer(global.CONFIG.Dictionary)
	engine.Dir = global.CONFIG.Data
	engine.Debug = global.CONFIG.Debug
	engine.Tokenizer = tokenizer
	engine.Shard = global.CONFIG.Shard
	engine.Timeout = global.CONFIG.Timeout
	engine.BufferNum = global.CONFIG.BufferNum

	engine.Init()

	// 初始化业务逻辑
	service.NewServices()

	// 注册路由
	r := router.SetupRouter()
	// 启动服务
	srv := &http.Server{
		Addr:    global.CONFIG.Addr,
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("listen:", err)
		}
	}()

	// close
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

func NewTokenizer(dictionaryPath string) *tokenizer.Tokenizer {
	return tokenizer.NewTokenizer(dictionaryPath)
}
