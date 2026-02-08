package cmd

import (
	"web/controller"
	"web/dao/order"
	"web/dao/user"

	"github.com/google/wire"
)

func InitializeApplication() (userController *controller.UserController, orderController *controller.OrderController) {

	wire.Build(
		controller.NewOrderController,
		controller.NewUserController,
		user.NewUserDao(),
		order.NewOrderDao,
	)
	return &controller.UserController{}, &controller.OrderController{}
}
