package controllers

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	"server/internal/dto"
	userEntity "server/internal/models/user"
)

func (s *Server) CreateJob(context *gin.Context) {
	user, exists := context.Get("user")
	if !exists {
		context.AbortWithStatus(http.StatusNotFound)
	}

	var createJobRequest dto.CreateJobRequest
	err := context.ShouldBindJSON(&createJobRequest)
	if err != nil {
		log.Printf("failed to create job request object: %s\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createJobResponse, err := s.Interactor.CreateJob(createJobRequest, user.(*userEntity.User))
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
	err := s.Interactor.StopJob(jobID)
	if err != nil {
		log.Printf("failed to stop job: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, "")
}

func (s *Server) GetJobs(context *gin.Context) {
	user, exists := context.Get("user")
	if !exists {
		context.AbortWithStatus(http.StatusNotFound)
	}

	getJobsResponse, err := s.Interactor.GetJobs(user.(*userEntity.User))
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
	getJobStatusResponse, err := s.Interactor.GetJobStatus(jobID)
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
	getJobLogsResponse, err := s.Interactor.GetJobLogs(jobID)
	if err != nil {
		log.Printf("failed to get job logs: %s\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusOK, *getJobLogsResponse)
}
