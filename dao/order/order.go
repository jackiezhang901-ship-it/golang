package order

type Order struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	OrderNo    string `json:"order_no"`
	UserId     int64  `json:"user_id"`
	TotalPrice int64  `json:"total_price"`
	PayStatus  int    `json:"pay_status"`
	PayType    int    `json:"pay_type"`
}

func (o *Order) TableName() string {
	return "order"
}
