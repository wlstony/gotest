package heartbeat

import (
	"fmt"
	"github.com/rabbitmq/rabbit"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbit.New(os.Getenv("RABBITMQ_SERVER"))
	fmt.Println("data StartHeartbeat")
	defer func() {
		fmt.Println("close data StartHeartbeat")
		q.Close()
	}()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
