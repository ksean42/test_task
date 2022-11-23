package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ksean42/test_task/pkg/model"
	"log"
	"net/http"
)

func (h *Handler) Get(c *gin.Context) {
	id, ok := c.GetQuery("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad request"})
		return
	}
	resp, err := h.service.Get(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, model.ResultResponse{Result: resp})
}

func (h *Handler) Set(c *gin.Context) {
	req := &model.UserGrade{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad request"})
		return
	}
	err = h.service.Set(req)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad request"})
		return
	}
	err = h.service.Set(req)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Result": "OK"})
}
