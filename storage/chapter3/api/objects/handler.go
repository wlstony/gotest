package objects

import (
	"fmt"
	"github.com/api/database"
	"github.com/api/heartbeat"
	"github.com/api/location"
	"github.com/api/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		//tmp, e := ioutil.ReadAll(r.Body)
		//fmt.Println("tmp:", string(tmp))
		//if e != nil{
		//	 fmt.Println("put error:", e.Error())
		//}
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func put(w http.ResponseWriter, r *http.Request) {
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, e := storeObject(r.Body, url.PathEscape(hash))
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size := utils.GetSiseFromHeader(r.Header)
	e = database.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(c)
}
func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusInternalServerError, e
	}
	//a1, a2 := ioutil.ReadAll(r)
	io.Copy(stream, r)

	//fmt.Println("a1:", string(a1))
	//if a2 != nil {
	//	fmt.Println("a2:", a2.Error())
	//}

	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}
func putStream(object string) (*PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return NewPutStream(server, object), nil
}
func get(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	versionId := r.URL.Query()["version"]
	version := 0
	var ve error
	if len(versionId) != 0 {
		version, ve = strconv.Atoi(versionId[0])
		if ve != nil {
			log.Println(ve)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	meta, me := database.GetMetadata(name, version)
	if me != nil {
		log.Println(me)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if meta.Hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	object := url.PathEscape(meta.Hash)
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}
func getStream(object string) (io.Reader, error) {
	server := location.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return NewGetStream(server, object)
}

func del(w http.ResponseWriter, r *http.Request)  {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	version, e := database.SearchLatestVersion(name)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = database.PutMetadata(name, version.Version+1, 0, "")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}