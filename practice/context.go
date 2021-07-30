package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go handle(ctx, 6*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("main done", ctx.Err(), time.Now())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle done", ctx.Err(), time.Now())
	case <-time.After(duration):
		fmt.Println("process request with", duration, time.Now())
	}
}
