package controllers

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"server/internal/dto"
)

func (s *Server) CreateJob(context *gin.Context) {
	apiToken, exists := context.Get("apiToken")
	if !exists {
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	var createJobRequest dto.CreateJobRequest
	err := context.ShouldBindJSON(&createJobRequest)
	if err != nil {
		log.Printf("failed to create job request object: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createJobResponse, err := s.interactor.CreateJob(createJobRequest, apiToken.(string))
	if err != nil {
		log.Printf("failed to create job: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, createJobResponse)
}

// TODO return 404 code when job doesn't exist
func (s *Server) StopJob(context *gin.Context) {
	jobID := context.Param("id")
	err := s.interactor.StopJob(jobID)
	if err != nil {
		log.Printf("failed to stop job: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, "")
}

func (s *Server) GetJobs(context *gin.Context) {
	getJobsResponse, err := s.interactor.GetJobs()
	if err != nil {
		log.Printf("failed to get jobs: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, getJobsResponse)
}

// TODO return 404 code when job doesn't exist
func (s *Server) GetJobStatus(context *gin.Context) {
	jobID := context.Param("id")
	getJobStatusResponse, err := s.interactor.GetJobStatus(jobID)
	if err != nil {
		log.Printf("failed to get job status: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, getJobStatusResponse)
}

// TODO return 404 code when job doesn't exist
func (s *Server) GetJobLogs(context *gin.Context) {
	jobID := context.Param("id")
	getJobLogsResponse, err := s.interactor.GetJobLogs(jobID)
	if err != nil {
		log.Printf("failed to get job logs: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, *getJobLogsResponse)
}
