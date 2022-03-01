package main

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)
//func slowHandler(w http.ResponseWriter, req *http.Request) {
//	time.Sleep(2 * time.Second)
//	io.WriteString(w, "I am slow!\n")
//}

//func slowAPICall() string {
//	d := rand.Intn(5)
//	select {
//	case <-time.After(time.Duration(d) * time.Second):
//		log.Printf("Slow API call done after %s seconds.\n", d)
//		return "foobar"
//	}
//}
//
//func slowHandler(w http.ResponseWriter, r *http.Request) {
//	result := slowAPICall()
//	io.WriteString(w, result+"\n")
//}

func slowAPICall(ctx context.Context) string {
	d := rand.Intn(5)
	select {
	case <-ctx.Done():
		log.Printf("slowAPICall was supposed to take %s seconds, but was canceled.", d)
		return ""
		//time.After() 可能会导致内存泄漏
	case <-time.After(time.Duration(d) * time.Second):
		log.Printf("Slow API call done after %d seconds.\n", d)
		return "foobar"
	}
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	result := slowAPICall(r.Context())
	io.WriteString(w, result+"\n")
}


func main() {
	//srv := http.Server{
	//	Addr:         ":8888",
	//	WriteTimeout: 1 * time.Second,
	//	Handler:      http.HandlerFunc(slowHandler),
	//}
	srv := http.Server{
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "Timeout!\n"),
	}


	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
