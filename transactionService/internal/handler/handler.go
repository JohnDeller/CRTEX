package handler

import (
	"github.com/gin-gonic/gin"
	"transactionService/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/add", h.createUser)
		user.GET("/:id", h.getUserById)
	}

	transaction := router.Group("/transaction")
	{
		transaction.POST("/create", h.createTransaction)
		transaction.GET("/:id", h.getTransactionById)
	}

	return router
}
