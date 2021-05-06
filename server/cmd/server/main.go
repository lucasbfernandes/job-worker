package main

import (
	"github.com/gin-gonic/gin"

	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"server/internal/controllers"
	"server/internal/storage"
	"strconv"
	"syscall"
)

func handleTerminationSignals() {
	terminationSignal := make(chan os.Signal)
	signal.Notify(terminationSignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-terminationSignal
		err := storage.DeleteLogsDir()
		if err != nil {
			fmt.Printf("failed to cleanup state before exiting %s\n", err)
		}
		os.Exit(1)
	}()
}

func createLogsDir() {
	err := storage.CreateLogsDir()
	if err != nil {
		log.Fatalf("failed to create logs dir %s\n", err)
	}
}

func createDB() {
	err := storage.CreateDB()
	if err != nil {
		log.Fatalf("failed to create db %s\n", err)
	}
}

func startAPI(appPort int) {
	router := gin.Default()

	router.POST("/jobs", controllers.CreateJob)
	router.POST("/jobs/:id/stop", controllers.StopJob)
	router.GET("/jobs", controllers.GetJobs)
	router.GET("/jobs/:id/status", controllers.GetJobStatus)
	router.GET("/jobs/:id/logs", controllers.GetJobLogs)

	err := router.Run(":" + strconv.Itoa(appPort))
	if err != nil {
		log.Fatalf("failed to start api: %s\n", err)
	}
}

func main() {
	appPort := flag.Int("port", 8080, "application port")
	flag.Parse()

	handleTerminationSignals()
	createLogsDir()
	createDB()
	startAPI(*appPort)
}
