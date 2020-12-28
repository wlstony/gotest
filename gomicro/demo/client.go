package main

import (
	"context"
	proto "demo/proto/demo"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {

	registry := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("demo1"),
		micro.Registry(registry),
		//micro.Broker(broker),
		//micro.Transport(transport),
	)

	//service := micro.NewService(micro.Name("demo"))
	service.Init()

	greeter := proto.NewGreeterService("demo", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "", Age:10})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
