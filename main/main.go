package main

import (
	"net/http"

	"github.com/B0go/itask-backend/config"
	"github.com/B0go/itask-backend/controller"
	"github.com/B0go/itask-backend/main/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	cfg, _ := config.LoadConfig()

	db := db.Database{Cfg: cfg}
	db.MustConnect()
	db.MustMigrate()

	tc := controller.TaskController{Db: &db}

	r := gin.Default()
	r.PUT("/tasks/:uuid", tc.PutTask)
	r.POST("/tasks", tc.PostTasks)
	r.DELETE("/tasks/:uuid", tc.DeleteTask)
	r.GET("/tasks/:uuid", tc.GetTask)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
