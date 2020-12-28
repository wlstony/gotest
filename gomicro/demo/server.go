package main

import (
	"context"
	proto "demo/proto/demo"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"strconv"
)

type Greeter struct {

}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error  {
	err := req.Validate()
	//fmt.Println("err:", err.Error(), ctx.Err())
	if err != nil {
		return  err
	}
	fmt.Printf("%+v\n", req)
	rsp.Greeting = "Welcome aa:" + req.Name + ", are you " + strconv.Itoa(int(req.Age)) + " years old?"
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
	proto.RegisterGreeterHandler(service.Server(), new(Greeter), )
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}