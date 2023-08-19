package router

import (
	"assignment-project/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/order", handler.CreateOrder)

	router.GET("/orders", handler.GetAllOrders)

	router.PUT("/order/:id", handler.UpdateOrder)

	router.DELETE("/order/:id", handler.DeleteOrder)

	return router
}
