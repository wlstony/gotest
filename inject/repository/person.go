package repository

import (
	"database/sql"
	"fmt"
	"inject/entity"
)

type PersonRepository struct {
	database *sql.DB
}

func (repository *PersonRepository) FindAll() []*entity.Person {
	rows, _ := repository.database.Query(
		`SELECT id, name, age FROM people;`,
	)
	defer rows.Close()

	people := []*entity.Person{}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		rows.Scan(&id, &name, &age)

		people = append(people, &entity.Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}

	return people
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	fmt.Println("NewPersonRepository")
	return &PersonRepository{database: database}
}
