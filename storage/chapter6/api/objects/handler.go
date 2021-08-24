package objects

import (
	"fmt"
	"github.com/api/constant"
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
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodPost {
		post(w, r)
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
	size := utils.GetSizeFromHeader(r.Header)
	c, e := storeObject(r.Body, hash, size)
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
	e = database.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(c)
}
func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if location.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}
	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusInternalServerError, e
	}
	//调用reader的时候会自动调用stream的write方法
	reader := io.TeeReader(r, stream)
	d := utils.CalculateHash(reader)
	fmt.Println("store object hash:", d)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("api store object hash mismatched %v, %v", d, hash)
	}
	//buf := make([]byte, size)
	//r.Read(buf)
	//fmt.Println("content:", string(buf))
	//_, err := stream.Write(buf)
	//if err != nil {
	//	stream.Commit(false)
	//	return http.StatusBadRequest,fmt.Errorf("api store object write failed %v", err)
	//}
	stream.Commit(true)
	return http.StatusOK, nil
}
func putStream(hash string, size int64) (*RSPutStream, error) {
	servers := heartbeat.ChooseRandomDataServer(constant.AllShards, nil)
	if len(servers) != constant.AllShards {
		return nil, fmt.Errorf("can not find enough dataServer")
	}
	return NewRSPutStream(servers, hash, size)
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
	fmt.Println("meta:", meta)
	if me != nil {
		log.Println(me)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if meta.Hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hash := meta.Hash
	stream, e := getStream(hash, meta.Size)
	if e != nil {
		fmt.Println("getStream: ", e)
	}
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//_, e = io.Copy(w, stream)
	//if e != nil {
	//	log.Println(e)
	//	w.WriteHeader(http.StatusNotFound)
	//	return
	//}
	offset := utils.GetOffsetFromHeader(r.Header)
	if offset != 0 {
		stream.Seek(offset, io.SeekCurrent)
		w.Header().Set("content-range", fmt.Sprintf("bytes %d-%d/%d", offset, meta.Size-1, meta.Size))
		w.WriteHeader(http.StatusPartialContent)
	}
	io.Copy(w, stream)
	stream.Close()
}
func getStream(hash string, size int64) (*RSGetStream, error) {
	locateInfo := location.Locate(hash)
	if len(locateInfo) < constant.DataShard {
		return nil, fmt.Errorf("object %s locate fail, result %v", hash, locateInfo)
	}
	dataServers := make([]string, 0)
	if len(locateInfo) != constant.AllShards {
		dataServers = heartbeat.ChooseRandomDataServer(constant.AllShards-len(locateInfo), locateInfo)
	}
	return NewRSGetStream(locateInfo, dataServers, hash, size)
}

func del(w http.ResponseWriter, r *http.Request) {
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

func post(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size, e := strconv.ParseInt(r.Header.Get("size"), 0, 64)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hsh in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if location.Exist(url.PathEscape(hash)) {
		e = database.AddVersion(name, hash, size)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		return
	}
	ds := heartbeat.ChooseRandomDataServer(constant.AllShards, nil)
	if len(ds) != constant.AllShards {
		log.Println("can not find enough dataServer")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	stream, e := NewRSResumablePutStream(ds, name, url.PathEscape(hash), size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("location", "/temp/"+stream.ToToken())
	w.WriteHeader(http.StatusCreated)
}
