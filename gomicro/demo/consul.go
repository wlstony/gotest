package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	registry := consul.NewRegistry()
	//broker := kafka.NewBroker()
	//transport := rabbitmq.NewTransport()

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(registry),
		//micro.Broker(broker),
		//micro.Transport(transport),
	)

	service.Init()
	service.Run()

}