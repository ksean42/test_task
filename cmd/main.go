package main

import (
	"github.com/ksean42/test_task/pkg"
	"github.com/ksean42/test_task/pkg/handler"
	"github.com/ksean42/test_task/pkg/service"
	"github.com/ksean42/test_task/pkg/storage"
)

func main() {
	repo := storage.NewStorage()
	UserGradeService := service.NewUserGradeService(repo)
	h := handler.NewHandler(UserGradeService)

	s1 := &pkg.Server{}
	s2 := &pkg.Server{}

	go s1.Start("8070", h.NewGetRouter())
	s2.Start("8071", h.NewSetRouter())

}
