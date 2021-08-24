package temp

import (
	"fmt"
	"github.com/api/constant"
	"github.com/api/database"
	"github.com/api/location"
	"github.com/api/objects"
	"github.com/api/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodHead {
		head(w, r)
		return
	}
	if m == http.MethodPut {
		put(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func head(w http.ResponseWriter, r *http.Request)  {
	token := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := objects.NewRSResumablePutStreamFromToken(token)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	current := stream.CurrentSize()
	if current == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("content-length", fmt.Sprintf("%d", current))
}

func put(w http.ResponseWriter, r *http.Request)  {
	token := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := objects.NewRSResumablePutStreamFromToken(token)
	if e != nil {
		log.Println("token解析错误", e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	current := stream.CurrentSize()
	if current == -1 {
		fmt.Println("stream CurrentSize:", stream.CurrentSize())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	offset := utils.GetOffsetFromHeader(r.Header)
	if current != offset {
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		return
	}
	bytes := make([]byte, constant.BlockSize)
	for {
		n, e := io.ReadFull(r.Body, bytes)
		if e != nil && e != io.EOF && e != io.ErrUnexpectedEOF {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		current += int64(n)
		if current > stream.Size {
			stream.Commit(false)
			log.Println("resumable put exceed size current:%d, stream:%d", current, stream.Size)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if n != constant.BlockSize && current != stream.Size {
			return
		}
		stream.Write(bytes[:n])
		if current == stream.Size {
			stream.Flush()
			getStream, e := objects.NewRSResumableGetStream(stream.Servers, stream.Uuids, stream.Size)
			hash := url.PathEscape(utils.CalculateHash(getStream))
			if hash != stream.Hash {
				stream.Commit(false)
				log.Printf("resumable put done but hash mismatch hash:%s, stream:%s", hash, stream.Hash)
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if location.Exist(url.PathEscape(hash)) {
				stream.Commit(false)
			} else {
				stream.Commit(true)
			}
			e = database.AddVersion(stream.Name, stream.Hash, stream.Size)
			if e != nil {
				log.Println(e)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}

}