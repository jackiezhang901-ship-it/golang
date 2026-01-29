package controller

import (
	"strconv"
	. "web/dao/order"

	"github.com/gin-gonic/gin"
)

func AddOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindBodyWithJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	InsertOrder(&order)
	c.JSON(200, gin.H{"message": "Order added successfully"})
}

func GetOrders(c *gin.Context) {
	id := c.Param("id")
	newid, _ := strconv.Atoi(id)
	order := SelectById(newid)
	c.JSON(200, order)
}

func UpdateOrderInfo(c *gin.Context) {
	var order Order
	if err := c.ShouldBindBodyWithJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	UpdateOrder(&order)
	c.JSON(200, gin.H{"message": "Order added successfully"})
}
