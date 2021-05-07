package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/interactors"
	"server/internal/repository"
	"server/internal/sec"
	"strconv"
	"strings"
	"syscall"
)

type Server struct {
	interactor *interactors.ServerInteractor
	secService *sec.SecurityService
}

func NewServer() (*Server, error) {
	serverInteractor, err := interactors.NewServerInteractor()
	if err != nil {
		return nil, err
	}

	secService, err := sec.NewSecService()
	if err != nil {
		return nil, err
	}

	return &Server{
		interactor: serverInteractor,
		secService: secService,
	}, nil
}

func (s *Server) Start(port int) error {
	err := s.createLogsDir()
	if err != nil {
		return err
	}

	// TODO remove this after creating users CRUD
	err = s.interactor.Database.SeedUsers()
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

	router.Use(s.AuthorizeJWT())

	router.POST("/jobs", s.CreateJob)
	router.POST("/jobs/:id/stop", s.StopJob)
	router.GET("/jobs", s.GetJobs)
	router.GET("/jobs/:id/status", s.GetJobStatus)
	router.GET("/jobs/:id/logs", s.GetJobLogs)

	err := router.RunTLS(":"+strconv.Itoa(appPort), sec.GetTLSCertFilePath(), sec.GetTLSKeyFilePath())
	if err != nil {
		return fmt.Errorf("failed to start api: %s", err)
	}
	return nil
}

func (s *Server) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" && !strings.HasPrefix(authHeader, bearerSchema) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := s.secService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if apiToken, ok := claims["apiToken"]; !ok || apiToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("apiToken", claims["apiToken"])
	}
}
