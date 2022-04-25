package main

import (
	"github.com/gin-gonic/gin"

	personCtl "example.com/cadastro-power/pkg/adapter/in/rest/person"
	"example.com/cadastro-power/pkg/application/person"
)

func (s *server) SetupRouter() {

	// Config the server
	apiGroup := s.engine.Group("api/v1")

	// Person route
	s.personRouter(apiGroup)

	// Address route
	//s.addressRouter(apiGroup)
}

func (s *server) personRouter(group *gin.RouterGroup) {

	if err := s.container.Invoke(func(personSrv person.Service) {

		personCtl := personCtl.NewController(personSrv)
		personCtl.SetupRouter(group)

	}); err != nil {
		panic(err)
	}

}
