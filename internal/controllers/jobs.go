package controllers

import (
	"github.com/gin-gonic/gin"

	"job-worker/internal/dto"
	"job-worker/internal/interactors"
	"log"
)

func CreateJob(context *gin.Context) {
	var createJobRequest dto.CreateJobRequest
	err := context.Bind(&createJobRequest)
	if err != nil {
		log.Printf("failed to create job request object: %s\n", err)
		context.JSON(400, gin.H{"error": err})
	}

	jobID, err := interactors.CreateJob(createJobRequest)
	if err != nil {
		log.Printf("failed to create job: %s\n", err)
		context.JSON(500, gin.H{"error": err})
	}
	context.JSON(201, dto.CreateJobResponse{ID: jobID})
}

func StopJob(context *gin.Context) {
}

func GetJobs(context *gin.Context) {
}

func GetJobStatus(context *gin.Context) {
}

func GetJobLogs(context *gin.Context) {
}
