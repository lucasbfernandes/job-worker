package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"fmt"
	"job-worker/internal/controllers"
	"job-worker/internal/storage"
	"log"
	"os"
)

// TODO this configuration method should be refactored for production releases
func setupEnv() {
	env := os.Getenv("GO_ENV")
	if "" == env {
		err := godotenv.Load(".env.dev")
		if err != nil {
			log.Fatalf(fmt.Sprintf("failed to load .env: %s\n", err))
		}
	}
}

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
	setupEnv()
	createLogsDir()
	createDB()
	startAPI()
}
