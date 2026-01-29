package order

import (
	"fmt"
	. "web/db"
)

func SelectById(id int) *Order {
	fmt.Println("select order by id")
	var oder Order
	DB.Raw("select * from order where id = ?", id).Scan(&oder)
	return &oder
}

func InsertOrder(order *Order) {
	fmt.Println("insert order")
	DB.Create(order)
}

func UpdateOrder(order *Order) {
	fmt.Println("update order")
	DB.Save(order)
}
