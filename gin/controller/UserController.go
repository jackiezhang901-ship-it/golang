package controller

import (
	"log"
	"net/http"
	"strconv"
	. "web/dao/user"

	"github.com/gin-gonic/gin"
)

func SelectUserById(c *gin.Context) {
	id := c.DefaultQuery("id", "4")
	num, _ := strconv.Atoi(id)
	user := SelectById(num)
	c.JSON(http.StatusOK, user)

}

func InserUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	AddUser(&user)
	c.JSON(http.StatusOK, user)
}

func UpdateUserInfo(c *gin.Context) {
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
	UpdateUser(&user)
	c.JSON(http.StatusOK, user)
}

func SelectUserList(c *gin.Context) {
	var users []User
	SelectAll(&users)
	c.JSON(http.StatusOK, users)
}
