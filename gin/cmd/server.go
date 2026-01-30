package cmd

import (
	"fmt"
	"net/http"
	"web/controller"
	"web/db"

	. "web/log"
	. "web/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	_ "net/http/pprof"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the web server",
	Long:  `Start the web server that provides user and order management APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Server port flag
	serverCmd.Flags().StringP("port", "p", "8080", "Server port")

	// Debug mode flag
	serverCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")
}

func startServer() {
	// Set Gin mode based on debug flag
	if debugMode, _ := serverCmd.Flags().GetBool("debug"); debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

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
