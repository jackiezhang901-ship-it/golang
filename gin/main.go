package main

import (
	"fmt"
	"net/http"
	"web/controller"
	"web/db"
	. "web/log"
	. "web/middleware"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	fmt.Printf("start to run service")
	Logger.Info("start to launch application")
	Logger.Info("start to initiate db")
	db.InitDB()
	r := gin.Default()
	r.Use(Cors())
	api := r.Group("api/user")
	{
		api.GET("", controller.SelectUserById)
		api.POST("", controller.InserUser)
		api.PUT(":id", controller.UpdateUserInfo)
		api.GET("/list", controller.SelectUserList)
	}

	r.Group("api/order")
	{
		r.GET("", controller.GetOrders)
		r.POST("", controller.AddOrder)
		r.PUT(":id", controller.UpdateOrderInfo)
	}
	r.Run(":8080")
	Logger.Info("application launch successuflly")

}
