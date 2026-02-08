package order

import (
	"fmt"
	. "web/conf/db"
)

type OrderDao struct {
}

func (o *OrderDao) SelectById(id int) *Order {
	fmt.Println("select order by id")
	var oder Order
	DB.Raw("select * from order where id = ?", id).Scan(&oder)
	return &oder
}

func (o *OrderDao) InsertOrder(order *Order) {
	fmt.Println("insert order")
	DB.Create(order)
}

func (o *OrderDao) UpdateOrder(order *Order) {
	fmt.Println("update order")
	DB.Save(order)
}

func NewOrderDao() *OrderDao {
	return &OrderDao{}
}
