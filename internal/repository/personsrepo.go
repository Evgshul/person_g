package repository

import (
	"log"

	"gorm.io/gorm"

	"github.com/evgshul/person_g/internal/entity"
)

type PersonRepository interface {
	Create(person *entity.Person) (*entity.Person, error)
	GetPersons() ([]entity.Person, error)
	GetById(id int) (*entity.Person, error)
	DeletePerson(id int) error
	UpdatePerson(person *entity.Person) (*entity.Person, error)

	InitTable() error
}

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) PersonRepository {
	return &personRepository{db: db}
}

func (r *personRepository) Create(person *entity.Person) (*entity.Person, error) {
	err := r.db.Create(person).Error
	return person, err
}

func (r *personRepository) GetPersons() ([]entity.Person, error) {
	var persons []entity.Person
	err := r.db.Find(&persons).Error
	return persons, err
}

func (r *personRepository) GetById(id int) (*entity.Person, error) {
	var person entity.Person
	err := r.db.First(&person, id).Error
	return &person, err
}

func (r *personRepository) DeletePerson(id int) error {
	return r.db.Delete(&entity.Person{}, id).Error
}

func (r *personRepository) UpdatePerson(person *entity.Person) (*entity.Person, error) {
	err := r.db.Save(person).Error
	return person, err
}

func (r *personRepository) InitTable() error {
	log.Println("Initializing persons table...")
	err := r.db.AutoMigrate((&entity.Person{}))
	if err != nil {
		log.Printf("Failed to initialize persons table: %v\n", err)
		return err
	}
	log.Println("Persons table initialized successfully.")
	return nil
}
