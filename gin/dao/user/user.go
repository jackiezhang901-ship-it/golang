package user

import (
	"time"
)

type User struct {
	Id      int       `gorm:"primaryKey" json:"id"`
	Name    string    `json:"name" binding:"required"`
	Age     int       `json:"age" binding:"required"`
	Email   string    `json:"email" binding:"required"`
	AddTime time.Time `gorm:"autoCreateTime"`
}

func (User) TableName() (name string) {
	return "user"
}
