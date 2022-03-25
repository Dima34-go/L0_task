package handler

import (
	"WB_GO_L0/pkg/httpServer/appError"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	idParam = "id"
)

func (h *Handler) GetUserById(c *gin.Context) error {
	orderId, err := strconv.Atoi(c.Param(idParam))
	if err != nil {
		return appError.ErrOrderNotFound
	}
	order, err := h.services.GetOrderById(orderId)
	if err != nil {
		return err
	}
	c.HTML(http.StatusOK,"index.html",order)
	return nil
}
