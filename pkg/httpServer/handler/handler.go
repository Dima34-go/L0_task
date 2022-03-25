package handler

import (
	"WB_GO_L0/pkg/httpServer/appError"
	"WB_GO_L0/pkg/httpServer/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//Get information about Order
	order := router.Group("/order")
	{
		order.GET("/:id", appError.Middleware(h.GetUserById))
	}

	//Path to external files: html,css
	router.LoadHTMLGlob("frontend/html/*")
	router.Static("/css/", "./frontend/css")

	return router
}
