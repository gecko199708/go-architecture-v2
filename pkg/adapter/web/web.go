package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Context = gin.Context
)

type Server struct {
	*gin.Engine

	Host string
	Port uint16

	server *http.Server
}

func NewServer() *Server {
	server := &Server{
		Engine: gin.New(),
	}
	return server
}

func (s *Server) Run() error {
	s.server = &http.Server{Addr: s.address(), Handler: s.Engine}
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Close(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
