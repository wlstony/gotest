package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
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
func GetSizeFromHeader(h http.Header) int64  {
	size, _ := strconv.ParseInt(h.Get("content-length"), 0, 64)
	return size

}
func CalculateHash(r io.Reader) string  {
	h := sha256.New()
	io.Copy(h, r)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}