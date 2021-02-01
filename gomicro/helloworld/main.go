package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"helloworld/handler"
	"helloworld/subscriber"

	helloworld "helloworld/proto/helloworld"
)

func main() {

	consulRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.helloworld"),
		micro.Registry(consulRegistry),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.helloworld", service.Server(), new(subscriber.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
