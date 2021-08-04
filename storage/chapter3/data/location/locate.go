package location

import (
	"fmt"
	"github.com/rabbitmq/rabbit"
	"os"
	"strconv"
	"strings"
)

func Locate(name string) bool {
	fmt.Println("data locate name:", name)
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
func StartLocate() {
	fmt.Println("data StartLocate")
	q := rabbit.New(os.Getenv("RABBITMQ_SERVER"))
	defer func() {
		fmt.Println("close data StartLocate")
		q.Close()
	}()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))
		fmt.Println("data locate ", object)
		if e != nil {
			panic(e)
		}
		//if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
		//	q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		//}
		if Locate(os.Getenv("STORAGE_ROOT") + "/" + strings.Trim(object, " ")) {
			fmt.Println("exist")
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
