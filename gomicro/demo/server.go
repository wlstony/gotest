package main

import (
	"context"
	proto "demo/proto/demo"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type Greeter struct {

}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error  {
	rsp.Greeting = "Welcome aa:" + req.Name
	return nil
}
func main() {
	registry := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("demo"),
		micro.Registry(registry),
		//micro.Broker(broker),
		//micro.Transport(transport),
	)
	//service := micro.NewService(micro.Name("demo"))
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}