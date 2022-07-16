package controller

import (
	"errors"
	"net/http"

	"github.com/B0go/itask-backend/main/db"
	"github.com/B0go/itask-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskController struct {
	Db *db.Database
}

func (tc *TaskController) PutTask(c *gin.Context) {
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

	if err := tc.Db.Db.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) PostTasks(c *gin.Context) {
	var tasks []model.Task
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.Db.Db.Create(&tasks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, tasks)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	var task model.Task
	task.Uuid = c.Param("uuid")

	if err := tc.Db.Db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Status(http.StatusOK)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	var task model.Task
	if err := tc.Db.Db.First(&task, "uuid = ?", c.Param("uuid")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		return
	}
	c.JSON(http.StatusOK, &task)
}
