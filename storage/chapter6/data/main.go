package main

import (
	"github.com/data/heartbeat"
	"github.com/data/location"
	"github.com/data/objects"
	"github.com/data/temp"
	"log"
	"net/http"
	"os"
)

/*
example
export RABBITMQ_SERVER=amqp://guest:guest@127.0.0.1:32771 && export STORAGE_ROOT=/tmp/storage/8001 && LISTEN_ADDRESS=127.0.0.1:8001 && go run data/main.go
export RABBITMQ_SERVER=amqp://guest:guest@127.0.0.1:32771 && export STORAGE_ROOT=/tmp/storage/8002 && LISTEN_ADDRESS=127.0.0.1:8002 && go run data/main.go
export RABBITMQ_SERVER=amqp://guest:guest@127.0.0.1:32771 && export STORAGE_ROOT=/tmp/storage/8003 && LISTEN_ADDRESS=127.0.0.1:8003 && go run data/main.go
RABBITMQ_SERVER=amqp://guest:guest@127.0.0.1:32771;LISTEN_ADDRESS=127.0.0.1:8001;STORAGE_ROOT=/tmp/8001
*/

func main() {
	location.CollectObjects()
	go heartbeat.StartHeartbeat()
	go location.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}