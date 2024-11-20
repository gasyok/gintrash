package handler

import (
	"net/http"

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

type contract interface{}

func Router(h *Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/healthz")
	r.GET("/hell", h.BaseGet)

	return r
}

func R(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
