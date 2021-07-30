package location

import (
	"fmt"
	"github.com/rabbitmq/rabbit"
	"os"
	"strconv"
)

func Locate(name string) bool {
	fmt.Println("data locate name:", name)
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
func StartLocate() {
	q := rabbit.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		//if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
		//	q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		//}
		if Locate(os.Getenv("STORAGE_ROOT") + "/" + object) {
			fmt.Println("exist")
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
