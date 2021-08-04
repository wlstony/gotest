package utils

import (
	"net/http"
	"strconv"
)

func GetHashFromHeader(h http.Header) string {
	digest := h.Get("digest")
	if len(digest) < 9 {
		return ""
	}
	if digest[:7] != "SHA-256" {
		return ""
	}
	return digest[7:]

}
func GetSiseFromHeader(h http.Header) int64  {
	size, _ := strconv.ParseInt(h.Get("content-length"), 0, 64)
	return size

}