package main

import (
	"fmt"
	"net/http"

	"github.com/B0go/itask-backend/model"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("something")

	r := gin.Default()

	r.PUT("/tasks/:uuid", func(c *gin.Context) {
		var task model.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task.Uuid = c.Param("uuid")

		c.JSON(http.StatusOK, task)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
