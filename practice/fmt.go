package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	str := "i%3D165431679%26m%3D0%26n%3D%25CA%25D4%25CA%25D4a%26s%3Dc6087af34e23e03d4db61f78c42ca7a0%26st%3D143549%26t%3D1649763167%26u%3D165431679%26v%3D1.1"
	decodeURl, _ := url.QueryUnescape(str)
	params, err := url.ParseQuery(decodeURl)

	fmt.Println(params, err, params.Get("u"))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
