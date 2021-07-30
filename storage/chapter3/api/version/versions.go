package version

import (
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request)  {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	from := 0
	size := 1000
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	for  {
		from += size
	}
}