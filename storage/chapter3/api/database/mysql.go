package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Metadata struct {
	Name string
	Version int
	Size int64
	Hash string
}

func getMetadata(name string, versionId int)  (meta Metadata, e error) {
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d", os.Getenv("META_SERVER"), name, versionId)
	r, e := http.Get(url)
	if e != nil {
		return
	}
	result, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(result, &meta)
	return
}

type hit struct {
	Source Metadata `json:"_source"`
}
type searchResult struct {
	Hits struct{
		Total int
		Hits []hit
	}
}
func SearchLatestVersion(name string) (meta Metadata, e error)  {
	url := fmt.Sprintf("http://%s/metadata/_search?q=, ", os.Getenv("META_SERVER"))
	r, e := http.Get(url)
	if e != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		e = fmt.Errorf("fail to search lastest")
		return
	}
	result, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(result, &sr)
	if len(sr.Hits.Hits) != 0 {
		meta = sr.Hits.Hits[0].Source
	}
	return
}
func GetMetadata(name string, version int) (meta Metadata, e error)  {
	if version ==0 {
		return SearchLatestVersion(name)
	}
	return getMetadata(name, version)
}
func PutMetadata(name string, version int, size init64, hash string)  {

}




















