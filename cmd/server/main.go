package main

import (
	"github.com/gin-gonic/gin"

	"job-worker/internal/controllers"
	"job-worker/internal/storage"
)

func createLogsDir() {
	storage.CreateLogsDir()
}

func createDB() {
	storage.CreateDB()
}

func startAPI() {
	router := gin.Default()
	router.POST("/jobs", controllers.CreateJob)
	router.POST("/jobs/:id/stop", controllers.StopJob)
	router.GET("/jobs", controllers.GetJobs)
	router.GET("/jobs/:id/status", controllers.GetJobStatus)
	router.GET("/jobs/:id/logs", controllers.GetJobLogs)
	router.Run(":8080")
}

func main() {
	createLogsDir()
	createDB()
	startAPI()
}
