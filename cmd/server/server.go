package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"example.com/cadastro-power/pkg/infra/config"
)

type server struct {
	container *dig.Container
	engine    *gin.Engine
}

func NewServer(cont *dig.Container, engine *gin.Engine) *server {
	return &server{container: cont, engine: engine}
}

func (s *server) Start() error {

	var cfg *config.Config
	if err := s.container.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}

	return s.engine.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
