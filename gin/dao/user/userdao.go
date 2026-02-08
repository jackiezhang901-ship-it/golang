package user

import (
	"fmt"
	"web/db"
)

type UserDao struct {
}

func (u *UserDao) SelectById(id int) (user *User) {
	db.DB.First(&user, id)
	fmt.Printf("id is %d", id)
	fmt.Println("user is", user)
	return user
}

func (u *UserDao) AddUser(user *User) (id int) {
	fmt.Printf("user info is %s:", user.Name)
	db.DB.Create(&user)
	return int(user.Id)
}

func (u *UserDao) UpdateUser(user *User) (id int) {
	fmt.Printf("user info is %s:", user.Name)
	db.DB.Model(&user).Where(&user.Id).Updates(&user)
	return int(user.Id)
}

func (u *UserDao) SelectAll(users *[]User) {
	db.DB.Find(&users)
	fmt.Println(users)
}

func NewUserDao() *UserDao {
	return &UserDao{}
}
