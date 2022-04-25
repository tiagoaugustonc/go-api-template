package main

import (
	"fmt"
	"os"

	"example.com/cadastro-power/pkg/infra/di"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

}

func run() error {

	container := di.BuildContainer()
	engine := gin.Default()

	server := NewServer(container, engine)
	server.SetupRouter()

	return server.Start()
}
