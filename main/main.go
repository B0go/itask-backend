package main

import (
	"errors"
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

		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusCreated, task)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var tasks []model.Task
		if err := c.ShouldBindJSON(&tasks); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&tasks).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusCreated, tasks)
	})

	r.DELETE("/tasks/:uuid", func(c *gin.Context) {
		var task model.Task
		task.Uuid = c.Param("uuid")

		if err := db.Delete(&task).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Status(http.StatusOK)
	})

	r.GET("/tasks/:uuid", func(c *gin.Context) {
		var task model.Task
		if err := db.First(&task, "uuid = ?", c.Param("uuid")).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusNotFound)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			}
			return
		}
		c.JSON(http.StatusOK, &task)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
