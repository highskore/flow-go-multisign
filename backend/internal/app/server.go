package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lukaracki/flow-go-multisign/backend/api/cosign"
)

// Server is the main struct for the server application
type Server struct {
	Router *gin.Engine
}

// NewServer creates a new instance of the Server struct
func NewServer() *Server {
	r := gin.Default()
	r.Use(cors.Default()) // TODO: Change for non POC version
	s := &Server{
		Router: r,
	}
	return s
}

// Start starts the server
func (s *Server) Start() error {
	// Define the auth route and its corresponding handler
	s.Router.POST("/cosign", cosign.AuthHandler)

	// Start the server
	return s.Router.Run()
}
