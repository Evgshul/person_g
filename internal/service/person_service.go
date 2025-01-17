package service

import (
	"fmt"

	"github.com/evgshul/person_g/internal/dto"
	"github.com/evgshul/person_g/internal/entity"
	_ "github.com/evgshul/person_g/internal/entity/errors"
	"github.com/evgshul/person_g/internal/repository"
)

type PersonService interface {
	CreatePerson(dto *dto.PersonDto) (*dto.ResponsePersonDto, error)
	GetPersonsList() ([]dto.ResponsePersonDto, error)
	GetPersonById(id int) (*dto.ResponsePersonDto, error)
	DeletePerson(id int) error
	UpdatePerson(id int, dto dto.PersonDto) (*dto.ResponsePersonDto, error)
}

type personService struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) PersonService {
	return &personService{repo: repo}
}

func (s *personService) CreatePerson(dto *dto.PersonDto) (*dto.ResponsePersonDto, error) {
	person := mapPersonDtoToPerson(dto)
	createdPerson, err := s.repo.Create(person)
	if err != nil {
		return nil, fmt.Errorf(" Error occured during new Person creation with name %v error: %v", dto.FullName, err)
	}
	return mapToResponsePersonDto(createdPerson), nil
}

func (s *personService) GetPersonsList() ([]dto.ResponsePersonDto, error) {
	persons, err := s.repo.GetPersons()
	if err != nil {
		return nil, err
	}
	var response []dto.ResponsePersonDto
	for _, person := range persons {
		response = append(response, *mapToResponsePersonDto(&person))
	}
	return response, nil
}

func (s *personService) GetPersonById(id int) (*dto.ResponsePersonDto, error) {
	person, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return mapToResponsePersonDto(person), nil
}

func (s *personService) DeletePerson(id int) error {
	return s.repo.DeletePerson(id)
}

func (s *personService) UpdatePerson(id int, personToUpdate dto.PersonDto) (*dto.ResponsePersonDto, error) {
	person, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	person.Fullname = personToUpdate.FullName
	person.Gender = personToUpdate.Gender
	person.PhoneNumber = personToUpdate.PhoneNumber
	person.Email = personToUpdate.Email

	updatedPerson, err := s.repo.UpdatePerson(person)
	if err != nil {
		return nil, err
	}
	return mapToResponsePersonDto(updatedPerson), nil
}

func mapToResponsePersonDto(p *entity.Person) *dto.ResponsePersonDto {
	return &dto.ResponsePersonDto{
		ID:          p.ID,
		FullName:    p.Fullname,
		Gender:      p.Gender,
		PhoneNumber: p.PhoneNumber,
		Email:       p.Email,
	}
}

func mapPersonDtoToPerson(personDto *dto.PersonDto) *entity.Person {
	return &entity.Person{
		Fullname:    personDto.FullName,
		Gender:      personDto.Gender,
		PhoneNumber: personDto.PhoneNumber,
		Email:       personDto.Email,
	}
}
