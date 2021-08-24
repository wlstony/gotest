package main

import (
	"github.com/api/heartbeat"
	"github.com/api/location"
	"github.com/api/objects"
	"github.com/api/temp"
	"github.com/api/version"
	"log"
	"net/http"
	"os"
)

/*
example
export
export RABBITMQ_SERVER=amqp://guest:guest@127.0.0.1:32771 &&  export LISTEN_ADDRESS=127.0.0.1:7001 && go run main.go
curl -X PUT 127.0.0.1:7001/objects/test2 -d 'this is test2'
curl -X GET 127.0.0.1:7001/objects/test2

curl -vvv http://127.0.0.1:7001/versions/test1
curl -vvv http://127.0.0.1:7001/objects/test2 -X PUT -d 'this is test2' -H "digest:SHA-256abcd1234"
curl -vvv -XDELETE  http://127.0.0.1:7001/objects/test1
*/

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", location.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/versions/", version.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
