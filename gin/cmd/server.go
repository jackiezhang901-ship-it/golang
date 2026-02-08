package cmd

import (
	"fmt"
	"net/http"
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
		debug, _ := cmd.Flags().GetBool("debug")
		port, _ := cmd.Flags().GetString("port")
		startServer(debug, port)
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

func startServer(debug bool, port string) {
	// Set Gin mode based on debug flag
	if debug {
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
	UserController, OrderController := InitializeApplication()
	db.InitDB()
	r := gin.Default()
	r.Use(Cors())
	api := r.Group("api/user")
	{
		api.GET("", UserController.SelectUserById)
		api.POST("", UserController.InsertUser)
		api.PUT(":id", UserController.UpdateUserInfo)
		api.GET("/list", UserController.SelectUserList)
	}

	r.Group("api/order")
	{
		r.GET("", OrderController.GetOrders)
		r.POST("", OrderController.AddOrder)
		r.PUT(":id", OrderController.UpdateOrderInfo)
	}

	fmt.Println("service running on port 8080")
	if err := r.Run(":8080"); err != nil {
		Logger.Error("failed to launch application: " + err.Error())
	}
}
