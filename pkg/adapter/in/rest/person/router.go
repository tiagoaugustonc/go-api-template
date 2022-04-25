package person

import "github.com/gin-gonic/gin"

func (c *personController) SetupRouter(group *gin.RouterGroup) {

	router := group.Group("/person")

	router.POST("/", c.Create)
	router.GET("/:id", c.Retrieve)
}
