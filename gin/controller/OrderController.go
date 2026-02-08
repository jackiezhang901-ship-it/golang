package controller

import (
	"strconv"
	. "web/dao/order"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderDao *OrderDao
}

func (o *OrderController) AddOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindBodyWithJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	o.orderDao.InsertOrder(&order)
	c.JSON(200, gin.H{"message": "Order added successfully"})
}

func (o *OrderController) GetOrders(c *gin.Context) {
	id := c.Param("id")
	newid, _ := strconv.Atoi(id)
	order := o.orderDao.SelectById(newid)
	c.JSON(200, order)
}

func (o *OrderController) UpdateOrderInfo(c *gin.Context) {
	var order Order
	if err := c.ShouldBindBodyWithJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	o.orderDao.UpdateOrder(&order)
	c.JSON(200, gin.H{"message": "Order added successfully"})
}

func NewOrderController(repo *OrderDao) *OrderController {
	return &OrderController{}
}
