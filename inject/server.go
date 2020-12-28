package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/dig"
	"inject/entity"
	"inject/repository"
	"inject/service"
	"net/http"
)

type Server struct {
	config        *entity.Config
	personService *service.PersonService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", s.people)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *entity.Config, service *service.PersonService) *Server {
	fmt.Println("NewServer")
	return &Server{
		config:        config,
		personService: service,
	}
}
func ConnectDatabase(config *entity.Config) (*sql.DB, error) {
	fmt.Println("ConnectDatabase")
	db, err := sql.Open("sqlite3", config.DatabasePath)
	if err != nil {
		panic(err)
	}
	return  db, err
}
//func main() {
//	config := entity.NewConfig()
//
//	db, err := ConnectDatabase(config)
//
//	if err != nil {
//		panic(err)
//	}
//
//	personRepository := repository.NewPersonRepository(db)
//
//	personService := service.NewPersonService(config, personRepository)
//
//	server := NewServer(config, personService)
//
//	server.Run()
//}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(entity.NewConfig)
	container.Provide(ConnectDatabase)
	container.Provide(repository.NewPersonRepository)
	container.Provide(service.NewPersonService)
	container.Provide(NewServer)

	return container
}

func main() {
	container := BuildContainer()

	err := container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}