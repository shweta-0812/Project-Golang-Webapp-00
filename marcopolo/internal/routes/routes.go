package routes

import (
	"github.com/gin-gonic/gin"
	"marcopolo/internal/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	// Group routes under a common prefix if needed
	api := router.Group("/api")
	{
		// User routes
		api.GET("/users", handlers.GetUsers)
		//api.GET("/users/:id", handlers.GetUserByID)
		//api.POST("/users", handlers.CreateUser)
		//api.PUT("/users/:id", handlers.UpdateUser)
		//api.DELETE("/users/:id", handlers.DeleteUser)
		//
		//// Payment routes
		//api.POST("/payments", handlers.CreatePayment)
		//api.GET("/payments/:id", handlers.GetPaymentByID)

		// Additional routes can be added here
	}
}
