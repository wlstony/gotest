package main

import (
	"context"
	proto "demo/proto/demo"
	"fmt"
	"github.com/micro/go-micro/v2"
)

type Greeter struct {

}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error  {
	rsp.Greeting = "Welcome " + req.Name
	return nil
}
func main() {
	service := micro.NewService(micro.Name("demo"))
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}