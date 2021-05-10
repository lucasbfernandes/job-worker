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
	userEntity "server/internal/models/user"
	"server/internal/repository"
	"server/internal/security"
	"strconv"
	"strings"
	"syscall"
)

type Server struct {
	Interactor *interactors.ServerInteractor
	secService *security.SecService
}

func NewServer() (*Server, error) {
	serverInteractor, err := interactors.NewServerInteractor()
	if err != nil {
		return nil, err
	}

	secService, err := security.NewSecService()
	if err != nil {
		return nil, err
	}

	return &Server{
		Interactor: serverInteractor,
		secService: secService,
	}, nil
}

func (s *Server) Start(port int) error {
	err := s.SetupState()
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

func (s *Server) SetupState() error {
	err := s.createLogsDir()
	if err != nil {
		return err
	}

	// TODO remove this after creating users CRUD
	err = s.Interactor.Database.SeedUsers()
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
	ginEngine := s.GetGinEngine()
	err := ginEngine.RunTLS(":"+strconv.Itoa(appPort), security.GetTLSCertFilePath(), security.GetTLSKeyFilePath())
	if err != nil {
		return fmt.Errorf("failed to start api: %s", err)
	}
	return nil
}

func (s *Server) GetGinEngine() *gin.Engine {
	router := gin.Default()

	router.Use(s.JWTGuard())
	router.Use(s.Authenticate())

	router.POST("/jobs", s.CreateJob)
	router.GET("/jobs", s.GetJobs)

	authzRoutes := router.Group("/")
	authzRoutes.Use(s.Authorize())

	authzRoutes.POST("/jobs/:id/stop", s.StopJob)
	authzRoutes.GET("/jobs/:id/status", s.GetJobStatus)
	authzRoutes.GET("/jobs/:id/logs", s.GetJobLogs)

	return router
}

func (s *Server) JWTGuard() gin.HandlerFunc {
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

func (s *Server) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiToken, exists := c.Get("apiToken")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := s.Interactor.Database.GetUserOrFailByAPIToken(apiToken.(string))
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Set("user", user)
	}
}

func (s *Server) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		userOBJ := user.(*userEntity.User)

		jobID := c.Param("id")
		job, err := s.Interactor.Database.GetJobOrFail(jobID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}

		// If user is ADMIN, then we skip this verification
		if userOBJ.Role == userEntity.UserRole && userOBJ.ID != job.UserID {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
