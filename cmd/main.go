package main

import (
	"context"
	"github.com/ksean42/test_task/pkg"
	"github.com/ksean42/test_task/pkg/handler"
	"github.com/ksean42/test_task/pkg/service"
	"github.com/ksean42/test_task/pkg/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	repo := storage.NewStorage()
	UserGradeService := service.NewUserGradeService(repo)
	h := handler.NewHandler(UserGradeService)

	getServer := &pkg.Server{}
	setServer := &pkg.Server{}

	go gracefulShutdown(ctx, cancel, getServer, setServer)

	go func() {

		err := getServer.Start("8070", h.NewGetRouter())
		if err != nil {
			log.Println(err)
		}
	}()

	func() {
		err := setServer.Start("8071", h.NewSetRouter())
		if err != nil {
			log.Println(err)
		}
	}()
	log.Println("Exited.")
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, getServer *pkg.Server, setServer *pkg.Server) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit
	cancel()

	log.Println("Shutting down....")
	if err := getServer.Stop(ctx); err != nil {
		log.Println(err)
	}
	if err := setServer.Stop(ctx); err != nil {
		log.Println(err)
	}
}
