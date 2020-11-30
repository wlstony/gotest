package heartbeat

import (
	"github.com/rabbitmq/rabbit"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbit.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
