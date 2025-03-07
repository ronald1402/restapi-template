package main

import (
	"context"
	"github.com/apex/log"
	"os"
	"os/signal"
	"restapi/internal/handler"
	"restapi/internal/repository"
	"restapi/internal/service"
	"restapi/server"
	"syscall"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	exitc := make(chan os.Signal, 1)
	signal.Notify(exitc, syscall.SIGINT, syscall.SIGTERM)

	e := server.Start()
	defer server.Stop(e)
	log.Infof("server started")
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	handler.NewHandler(e, svc)
	handler.NewHealthCheckHandler(e)

	sig := <-exitc
	log.Infof("received %s signal, exiting now", sig)
}
