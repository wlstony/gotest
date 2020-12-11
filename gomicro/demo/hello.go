package main

import (
	"fmt"
	"net/http"
)
//micro service --name hellomicro --endpoint http://localhost:9090 go run hello.go
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`hello micro`))
	})
	fmt.Println("listen")
	http.ListenAndServe(":9090", nil)
}