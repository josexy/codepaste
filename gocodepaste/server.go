package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	app     *gin.Engine
}

func NewServer(address string, app *gin.Engine) *Server {
	return &Server{
		address: address,
		app:     app,
	}
}

func (svr *Server) Run() {
	server := &http.Server{
		Addr:    svr.address,
		Handler: svr.app,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	// Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
