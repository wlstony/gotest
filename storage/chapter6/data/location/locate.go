package location

import (
	"fmt"
	"github.com/rabbitmq/rabbit"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)
var objects = make(map[string]int)
var mutex sync.Mutex

func Locate(hash string) int {
	fmt.Println("data locate objects:", objects)
	mutex.Lock()
	id, ok := objects[hash]
	mutex.Unlock()
	if !ok {
		return -1
	}
	return id
}
func Add(hash string, id int)  {
	mutex.Lock()
	objects[hash] = id
	mutex.Unlock()
}
func Del(hash string)  {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}
type locateMessage struct {
	Addr string
	Id   int
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
		id := Locate(hash)
		if id != -1 {
			//q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
			q.Send(msg.ReplyTo, locateMessage{
				Addr: os.Getenv("LISTEN_ADDRESS"),
				Id:   id,
			})
		}
	}
}

func CollectObjects()  {
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	fmt.Println("files:", os.Getenv("STORAGE_ROOT") + "/objects/*")
	for i := range files {
		file := strings.Split(filepath.Base(files[i]), ".")
		if len(file) != 3 {
			panic(files[i] + " file length not 3")
		}
		hash := file[0]
		id, e := strconv.Atoi(file[1])
		if e != nil {
			panic(e)
		}
		objects[hash] = id
	}
	fmt.Println("objects:",objects)
}