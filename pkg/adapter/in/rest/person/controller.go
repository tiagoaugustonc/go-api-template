package person

import (
	"net/http"

	"example.com/cadastro-power/pkg/application/person"
	"github.com/gin-gonic/gin"
)

type personController struct {
	service person.Service
}

func NewController(service person.Service) *personController {

	return &personController{
		service: service,
	}
}

func (c *personController) Create(ctx *gin.Context) {

	var dto PersonDTO

	// create and validate the payload
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// mapping to the model
	person := dto.toModel()

	// call to the service
	if personCreated, err := c.service.Create(*person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": personCreated})
	}

}

func (c *personController) Retrieve(ctx *gin.Context) {

	id := ctx.Param("id")

	// call to the service
	if person, err := c.service.Retrieve(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if person == nil {
		ctx.Status(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, person)
	}

}
