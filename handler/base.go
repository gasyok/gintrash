package handler

import (
	"errors"
	"fmt"
	"net/http"
	"rest/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BaseGet(c *gin.Context) {
	var resp struct {
		Count int    `json:"count"`
		Ans   string `json:"answer"`
	}
	flow := func() error {
		// c.JSON
		// cnt, err := h.

		info, err := h.baseService.Info()
		if err != nil {
			return fmt.Errorf("baseService.Info: %w", err)
		}
		resp.Count = info.Count
		resp.Ans = "alright"

		c.JSON(http.StatusOK, resp)

		return nil
	}

	var notFound *domain.ErrNotFound
	switch err := flow(); {
	case err == nil:
	case errors.Is(err, domain.ErrInvalidArgument):
		c.JSON(http.StatusBadRequest, "invalid argument")
		// http.Error(w, "invalid argument", http.StatusBadRequest)
		return
	case errors.As(err, &notFound):
		c.JSON(http.StatusNotFound, notFound.Error())
		// http.Error(w, notFound.Error(), http.StatusNotFound)
		return
	default:
		// logger.Error("Internal error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, "internal error")
		// http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
