package controller

import (
	"net/http"
	"strconv"

	"github.com/evgshul/person_g/internal/dto"
	"github.com/evgshul/person_g/internal/service"
	"github.com/gin-gonic/gin"
)

type PersonController struct {
	service service.PersonService
}

func NewPersonController(service service.PersonService) *PersonController {
	return &PersonController{service: service}
}

func (c *PersonController) CreatePerson(ctx *gin.Context) {
	var personDto dto.PersonDto
	if err := ctx.ShouldBindJSON(&personDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := c.service.CreatePerson(&personDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, person)
}

func (c *PersonController) GetPersonsList(ctx *gin.Context) {
	persons, err := c.service.GetPersonsList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, persons)
}

func (c *PersonController) GetPersonById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	person, err := c.service.GetPersonById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, person)
}

func (c *PersonController) UpdatePerson(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var personDto dto.PersonDto

	if err := ctx.ShouldBindJSON(&personDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, err := c.service.UpdatePerson(id, personDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, person)
}

func (c *PersonController) DeletePerson(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.DeletePerson(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": " Person deleted successfully"})
}
