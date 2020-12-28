package service

import (
	"fmt"
	"inject/entity"
	"inject/repository"
)

type PersonService struct {
	config     *entity.Config
	repository *repository.PersonRepository
}

func (service *PersonService) FindAll() []*entity.Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*entity.Person{}
}

func NewPersonService(config *entity.Config, repository *repository.PersonRepository) *PersonService {
	fmt.Println("NewPersonService")
	return &PersonService{config: config, repository: repository}
}
