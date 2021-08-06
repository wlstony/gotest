package location

import (
	"fmt"
	"github.com/rabbitmq/rabbit"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)
var objects = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) bool {
	fmt.Println("data locate objects:", objects)
	mutex.Lock()
	_, ok := objects[hash]
	mutex.Unlock()
	return ok
}
func Add(hash string)  {
	mutex.Lock()
	objects[hash] = 1
	mutex.Unlock()
}
func Del(hash string)  {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}

func StartLocate() {
	fmt.Println("data StartLocate")
	q := rabbit.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		hash, e := strconv.Unquote(string(msg.Body))
		fmt.Println("data locate ", hash)
		if e != nil {
			panic(e)
		}
		//if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
		//	q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		//}
		//if Locate(os.Getenv("STORAGE_ROOT") + "/" + strings.Trim(object, " ")) {
		//	fmt.Println("exist")
		//	q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		//}
		exist := Locate(hash)
		if exist {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}

func CollectObjects()  {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	for i := range files {
		hash := filepath.Base(files[i])
		objects[hash] = 1
	}
}