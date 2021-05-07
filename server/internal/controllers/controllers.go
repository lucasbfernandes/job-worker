package controllers

import (
	"github.com/gin-gonic/gin"

	"server/internal/interactors"
	"server/internal/repository"

	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	// This will create files inside pwd/cert
	defaultCertFilePath = "cert/server.crt"

	// This will create files inside pwd/cert
	defaultKeyFilePath = "cert/server.key"
)

type Server struct {
	interactor *interactors.ServerInteractor
}

func NewServer() (*Server, error) {
	serverInteractor, err := interactors.NewServerInteractor()
	if err != nil {
		return nil, err
	}

	return &Server{
		interactor: serverInteractor,
	}, nil
}

func (s *Server) Start(port int) error {
	err := s.createLogsDir()
	if err != nil {
		return err
	}

	s.handleTerminationSignals()

	err = s.startAPI(port)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) handleTerminationSignals() {
	terminationSignal := make(chan os.Signal)
	signal.Notify(terminationSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-terminationSignal
		err := repository.DeleteLogsDir()
		if err != nil {
			log.Printf("failed to cleanup state before exiting %s\n", err)
		}
		close(terminationSignal)
		os.Exit(1)
	}()
}

func (s *Server) createLogsDir() error {
	err := repository.CreateLogsDir()
	if err != nil {
		return fmt.Errorf("failed to create logs dir %s", err)
	}
	return nil
}

func (s *Server) startAPI(appPort int) error {
	router := gin.Default()

	router.POST("/jobs", s.CreateJob)
	router.POST("/jobs/:id/stop", s.StopJob)
	router.GET("/jobs", s.GetJobs)
	router.GET("/jobs/:id/status", s.GetJobStatus)
	router.GET("/jobs/:id/logs", s.GetJobLogs)

	err := router.RunTLS(":"+strconv.Itoa(appPort), getCertFilePath(), getKeyFilePath())
	if err != nil {
		return fmt.Errorf("failed to start api: %s", err)
	}
	return nil
}

func getCertFilePath() string {
	certFilePath, envExists := os.LookupEnv("TLS_CERT_FILE_PATH")
	if envExists && certFilePath != "" {
		return certFilePath
	}
	return defaultCertFilePath
}

func getKeyFilePath() string {
	keyFilePath, envExists := os.LookupEnv("TLS_KEY_FILE_PATH")
	if envExists && keyFilePath != "" {
		return keyFilePath
	}
	return defaultKeyFilePath
}
