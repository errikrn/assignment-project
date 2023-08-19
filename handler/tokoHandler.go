package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	ID           int       `json:"id"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items"`
}

var orders []Order
var nextID = 1

func CreateOrder(ctx *gin.Context) {
	var newOrder Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder.ID = nextID
	nextID++
	orders = append(orders, newOrder)

	ctx.JSON(http.StatusCreated, newOrder)
}

func GetAllOrders(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, orders)
}

func UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")
	var updatedOrder Order
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var found bool
	for i, order := range orders {
		if order.ID == id {
			orders[i].OrderedAt = updatedOrder.OrderedAt
			orders[i].CustomerName = updatedOrder.CustomerName
			orders[i].Items = updatedOrder.Items
			found = true
			break
		}
	}

	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}

func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var found bool
	for i, order := range orders {
		if order.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
