package routes

import (
	"go-task-api/controllers"
	"go-task-api/middlewares"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/tasks")
	taskGroup.Use(middlewares.AuthMiddleware())
	{
		taskGroup.GET("/", controllers.GetTasks)
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
}
