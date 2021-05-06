package controllers

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"server/internal/dto"
	"server/internal/interactors"
)

func CreateJob(context *gin.Context) {
	var createJobRequest dto.CreateJobRequest
	err := context.ShouldBindJSON(&createJobRequest)
	if err != nil {
		log.Printf("failed to create job request object: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createJobResponse, err := interactors.CreateJob(createJobRequest)
	if err != nil {
		log.Printf("failed to create job: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, createJobResponse)
}

// TODO return 404 code when job doesn't exist
func StopJob(context *gin.Context) {
	jobID := context.Param("id")
	err := interactors.StopJob(jobID)
	if err != nil {
		log.Printf("failed to stop job: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, "")
}

func GetJobs(context *gin.Context) {
	getJobsResponse, err := interactors.GetJobs()
	if err != nil {
		log.Printf("failed to get jobs: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, getJobsResponse)
}

// TODO return 404 code when job doesn't exist
func GetJobStatus(context *gin.Context) {
	jobID := context.Param("id")
	getJobStatusResponse, err := interactors.GetJobStatus(jobID)
	if err != nil {
		log.Printf("failed to get job status: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, getJobStatusResponse)
}

// TODO return 404 code when job doesn't exist
func GetJobLogs(context *gin.Context) {
	jobID := context.Param("id")
	getJobLogsResponse, err := interactors.GetJobLogs(jobID)
	if err != nil {
		log.Printf("failed to get job logs: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, *getJobLogsResponse)
}
