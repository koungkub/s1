package main

import (
	"context"
	"fmt"

	pb "github.com/koungkub/s1/proto"
	micro "github.com/micro/go-micro/v2"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	resp.Greeting = "hello"
	return nil
}

func main() {

	s := micro.NewService(
		micro.Name("koung-server"),
	)

	s.Init()

	pb.RegisterGreeterHandler(s.Server(), new(Greeter))

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
