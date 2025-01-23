package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/evgshul/person_g/internal/dto"
	"github.com/evgshul/person_g/internal/entity"
	"github.com/evgshul/person_g/internal/repository"
)

type PersonService interface {
	CreatePerson(dto *dto.PersonDto) (*dto.ResponsePersonDto, error)
	GetPersonsList() ([]dto.ResponsePersonDto, error)
	GetPersonById(id int) (*dto.ResponsePersonDto, error)
	SearchPersons(keyword string) ([]dto.ResponsePersonDto, error)
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

	isNameExist, err := s.isFullnameExist(dto.FullName)
	if err != nil {
		return nil, err
	}
	if isNameExist {
		return nil, fmt.Errorf("full name %s taken", dto.FullName)
	}

	isEmailExist, err := s.isEmailExist(dto.Email)
	if err != nil {
		return nil, err
	}
	if isEmailExist {
		return nil, fmt.Errorf("email %s taken please provide another", dto.Email)
	}
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

func (s *personService) SearchPersons(key string) ([]dto.ResponsePersonDto, error) {
	persons, err := s.repo.SearchPersons(key)
	if err != nil {
		return nil, err
	}

	var responseDtos []dto.ResponsePersonDto
	for _, person := range persons {
		responseDtos = append(responseDtos, dto.ResponsePersonDto{
			ID:          person.ID,
			FullName:    person.Fullname,
			Gender:      person.Gender,
			PhoneNumber: person.PhoneNumber,
			Email:       person.Email,
		})
	}
	return responseDtos, nil
}

func (s *personService) DeletePerson(id int) error {
	return s.repo.DeletePerson(id)
}

func (s *personService) UpdatePerson(id int, personToUpdate dto.PersonDto) (*dto.ResponsePersonDto, error) {
	person, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.New("person not found")
	}

	personNameToUpdate := personToUpdate.FullName
	if personNameToUpdate != "" && personNameToUpdate != person.Fullname {
		isFullNameNotUnique, err := s.isFullnameExist(personNameToUpdate)
		if err != nil {
			return nil, err
		}
		if isFullNameNotUnique {
			return nil, fmt.Errorf("person_service :: UpdatePerson :: fullname %s not unique,"+
				" available only one unique entry for a person", personNameToUpdate)

		}
		person.Fullname = personNameToUpdate
	}

	if strings.TrimSpace(personToUpdate.Gender) != "" && personToUpdate.Gender != person.Gender {
		person.Gender = personToUpdate.Gender
	}

	if strings.TrimSpace(personToUpdate.PhoneNumber) != "" && personToUpdate.PhoneNumber != person.PhoneNumber {
		person.PhoneNumber = personToUpdate.PhoneNumber
	}

	emailToUpdate := personToUpdate.Email
	if strings.TrimSpace(emailToUpdate) != "" && emailToUpdate != person.Email {
		isPhoneNumberExist, err := s.isEmailExist(emailToUpdate)
		if err != nil {
			return nil, err
		}
		if isPhoneNumberExist {
			return nil, fmt.Errorf("person_service :: UpdatePerson :: email %s is taken, please input another", emailToUpdate)
		}
		person.Email = emailToUpdate
	}

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

func (s *personService) isFullnameExist(fullname string) (bool, error) {
	person, err := s.repo.GetPersonByFullname(fullname)
	if err != nil {
		return false, err
	}
	if person != nil {
		return true, nil
	}
	return false, nil
}

func (s *personService) isEmailExist(email string) (bool, error) {
	person, err := s.repo.GetPersonByEmail(email)
	if err != nil {
		return false, err
	}
	if person != nil {
		return true, nil
	}
	return false, nil
}
