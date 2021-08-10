package location

import (
	"encoding/json"
	"fmt"
	"github.com/api/constant"
	"github.com/rabbitmq/rabbit"
	"net/http"
	"os"
	"strings"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("api locate info:", info)
	b, _ := json.Marshal(info)
	w.Write(b)
}

type locateMessage struct {
	Addr string
	Id   int
}

func Locate(name string) (locateInfo map[int]string) {
	q := rabbit.New(os.Getenv("RABBITMQ_SERVER"))
	fmt.Println("api Locate:", name)
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		//1 s 超时，无论拿到多少条消息
		time.Sleep(time.Second)
		q.Close()
	}()
	locateInfo = make(map[int]string)
	fmt.Println("api locate locateInfo:", locateInfo)
	for i:=0; i< constant.AllShards; i++  {
		msg := <-c
		fmt.Println("msg:", string( msg.Body))
		if len(msg.Body) == 0 {
			return
		}
		var info locateMessage
		err := json.Unmarshal(msg.Body, &info)
		if err != nil {
			fmt.Println("Locate:", err)
		}
		locateInfo[info.Id] = info.Addr
	}
	fmt.Println("api locate locateInfo:", locateInfo)

	return
}
func Exist(name string) bool {
	return len(Locate(name)) >= constant.DataShard
}

