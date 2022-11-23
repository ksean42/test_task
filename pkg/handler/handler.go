package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ksean42/test_task/pkg/middleware"
	"github.com/ksean42/test_task/pkg/service"
)

type Handler struct {
	service service.UserGrade
}

func NewHandler(service service.UserGrade) *Handler {
	return &Handler{service: service}
}
func (h *Handler) NewGetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/get", h.Get)
	return r
}

func (h *Handler) NewSetRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/", middleware.Auth)
	{
		auth.GET("/set", h.Set)
	}
	return r
}
