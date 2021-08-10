package objects

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type PutStream struct {
	writer *io.PipeWriter
	c      chan error
}

func NewPutStream(server, object string) *PutStream {
	reader, writer := io.Pipe()
	c := make(chan error)
	go func() {
		fmt.Println("NewPutStream:", "http://"+server+"/objects/"+object)
		request, _ := http.NewRequest("PUT", "http://"+server+"/objects/"+object, reader)
		//n, e := ioutil.ReadAll(reader)
		//fmt.Println("reader:", string(n), e)
		client := http.Client{}
		r, e := client.Do(request)
		fmt.Println("r:", r, "e:", e)
		if e != nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{writer, c}
}
func (w *PutStream) Write(p []byte) (n int, err error) {
	fmt.Println("PutStream: Write:", string(p))
	return w.writer.Write(p)
}
func (w *PutStream) Close() error {
	w.writer.Close()
	return <-w.c
}

type GetStream struct {
	reader io.Reader
}

func newGetStream(url string) (*GetStream, error) {
	fmt.Println("newGetStream url:", url)
	r, e := http.Get(url)
	if e != nil {
		fmt.Println("get stream error:", e.Error())
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %d", r.StatusCode)
	}
	return &GetStream{r.Body}, nil
}

func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, object)
	}
	return newGetStream("http://" + server + "/objects/" + object)
}
func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

type TempPutStream struct {
	Server string
	Uuid string
}

func NewTempPutStream(server, hash string, size int64) (*TempPutStream, error)  {
	fmt.Println("NewTempPutStream:", "http://" + server + "/temp/" + hash)
	req, e := http.NewRequest("POST", "http://" + server + "/temp/" + hash, nil)
	if e != nil {
		return nil, e
	}
	req.Header.Set("size", fmt.Sprintf("%d", size))
	client := http.Client{}
	response, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code not ok, %v", response)
	}
	fmt.Println("NewTempPutStream response.Body:", response.Body)
	uuid, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, e
	}
	return &TempPutStream{Server:server, Uuid:string(uuid)}, nil
}
func (w *TempPutStream) Write(p []byte)(n int, err error)  {
	//fmt.Println("write:",  "http://" + w.Server + "/temp/" + w.Uuid, string(debug.Stack()))
	request, e := http.NewRequest("PATCH", "http://" + w.Server + "/temp/" + w.Uuid, strings.NewReader(string(p)))
	if e != nil {
		return 0, e
	}
	client := http.Client{}
	r, e := client.Do(request)
	if e != nil {
		return 0, e
	}
	if r.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("dataServer return response %v, http code %d",r, r.StatusCode)
	}
	return len(p), nil
}
func (w *TempPutStream) Commit(good bool)  {
	method := "DELETE"
	if good {
		method = "PUT"
	}
	fmt.Println("api commit ", "http://" + w.Server + "/temp/" + w.Uuid, w)
	req, _ := http.NewRequest(method, "http://" + w.Server + "/temp/" + w.Uuid, nil)
	client := http.Client{}
	res, rerr := client.Do(req)
	fmt.Println("commit res:", res, ", error:", rerr)
}