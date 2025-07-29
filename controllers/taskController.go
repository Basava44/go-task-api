package controllers

import (
	"go-task-api/database"
	"go-task-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task

	userID := c.MustGet("userID").(uint)

	if err := database.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// "binding" from gin for input validation
	// This is to validate the string and to add requird fields
	var input struct {
		Title  string `json:"title" binding:"required"`
		Status string `json:"status" binding:"required,oneof=pending completed"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:  input.Title,
		Status: input.Status,
		UserID: userID,
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)

}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	userId := c.MustGet("userID").(uint)
	var task models.Task

	if err := database.DB.First(&task, taskId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if task.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	// "binding" from gin for input validation
	// This is to validate the string and to add requird fields
	var input struct {
		Title  string `json:"title" binding:"required"`
		Status string `json:"status" binding:"required,oneof=pending completed"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Only update fields allowed to change
	task.Title = input.Title
	task.Status = input.Status

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	userId := c.MustGet("userID").(uint)
	var task models.Task

	if err := database.DB.First(&task, taskId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if task.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
