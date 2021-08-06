package location

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/rabbit"
	"net/http"
	"os"
	"strconv"
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
func Locate(name string) string {
	q := rabbit.New(os.Getenv("RABBITMQ_SERVER"))
	fmt.Println("api Locate:", name)
	q.Publish("dataServers", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	fmt.Println("api locate body s:", s)
	return s
}
func Exist(name string) bool {
	return Locate(name) != ""
}
