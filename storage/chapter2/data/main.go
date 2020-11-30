package main

import (
	"github.com/data/heartbeat"
	"github.com/data/location"
	"github.com/data/objects"
	"log"
	"net/http"
	"os"
)
func main() {
	go heartbeat.StartHeartbeat()
	go location.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}