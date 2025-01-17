package main

import (
	"log"

	"github.com/evgshul/person_g/internal/config"
	"github.com/evgshul/person_g/internal/controller"
	"github.com/evgshul/person_g/internal/repository"
	"github.com/evgshul/person_g/internal/router"
	"github.com/evgshul/person_g/internal/service"
)

func main() {

	db := config.InitDb()
	//config.CreatTablePerson(*db)
	defer config.CloseDatabaseConnection()

	personRepo := repository.NewPersonRepository(db)
	personService := service.NewPersonService(personRepo)
	personController := controller.NewPersonController(personService)

	r := router.SetupRouter(*personController)

	log.Println("Application run on port: 8080")
	r.Run(":8080")
}
