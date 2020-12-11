package main

import (
	"context"
	proto "demo/proto/demo"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("demo"))
	service.Init()

	greeter := proto.NewGreeterService("demo", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
