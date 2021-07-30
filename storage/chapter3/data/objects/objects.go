package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func put(w http.ResponseWriter, r *http.Request) {
	path := os.Getenv("STORAGE_ROOT") + "/" +
		strings.Split(r.URL.EscapedPath(), "/")[2]
	fmt.Println("data put path:", path)
	//tmp, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	 fmt.Println(err.Error())
	//}
	//fmt.Println("put tmp:", string(tmp))
	f, e := os.Create(path)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
func get(w http.ResponseWriter, r *http.Request) {
	path := os.Getenv("STORAGE_ROOT") + "/" +
		strings.Split(r.URL.EscapedPath(), "/")[2]
	f, e := os.Open(path)
	fmt.Println("open path:", path)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
