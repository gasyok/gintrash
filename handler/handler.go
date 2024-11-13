package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	baseService baseService
	// logger
	// ...
}

func New(baseService baseService) *Handler {
	return &Handler{baseService: baseService}
}

func Router(h *Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/hell", h.BaseGet)

	return r
}
