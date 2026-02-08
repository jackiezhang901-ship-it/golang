package controller

import (
	"log"
	"net/http"
	"strconv"
	. "web/dao/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userDao *UserDao
}

func (u *UserController) SelectUserById(c *gin.Context) {
	id := c.DefaultQuery("id", "4")
	num, _ := strconv.Atoi(id)
	user := u.userDao.SelectById(num)
	c.JSON(http.StatusOK, user)

}

func (u *UserController) InsertUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	u.userDao.AddUser(&user)
	c.JSON(http.StatusOK, user)
}

func (u *UserController) UpdateUserInfo(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Panic("number conversion error:", err)
	}
	user.Id = num
	u.userDao.UpdateUser(&user)
	c.JSON(http.StatusOK, user)
}

func (u *UserController) SelectUserList(c *gin.Context) {
	var users []User
	u.userDao.SelectAll(&users)
	c.JSON(http.StatusOK, users)
}

func NewUserController(userDao *UserDao) *UserController {
	return &UserController{userDao: userDao}
}
