package main

import (
	"net/http"

	"github.com/B0go/itask-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=itask_app dbname=itask sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Task{})

	r := gin.Default()

	r.PUT("/tasks/:uuid", func(c *gin.Context) {
		u, err := uuid.Parse(c.Param("uuid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error parsing uuid": err.Error()})
			return
		}
		var task model.Task
		task.Uuid = u.String()
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&task)

		c.JSON(http.StatusCreated, task)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var tasks []model.Task
		if err := c.ShouldBindJSON(&tasks); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&tasks)

		c.JSON(http.StatusCreated, tasks)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
