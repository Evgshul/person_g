package router

import (
	"net/http"

	"github.com/evgshul/person_g/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(personController controller.PersonController) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Persons API!"})
	})

	//API versioning
	v1 := router.Group("/api/v1")
	{
		personRoutes := v1.Group("/persons")
		{
			personRoutes.POST("", personController.CreatePerson)
			personRoutes.GET("", personController.GetPersonsList)
			personRoutes.GET("/:id", personController.GetPersonById)
			personRoutes.GET("/search", personController.SearchPersons)
			personRoutes.PUT("/:id", personController.UpdatePerson)
			personRoutes.DELETE("/:id", personController.DeletePerson)
		}
	}
	return router
}
