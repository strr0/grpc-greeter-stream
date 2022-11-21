package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"greeter/proto"
	"io"
	"log"
)

func lotsRequest(c greeter.SayClient) error {
	stream, err := c.LotsRequest(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		if err := stream.Send(&greeter.Request{Name: "hello john~"}); err != nil {
			return err
		}
	}
	recv, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Println(recv)
	return nil
}

func lotsResponse(c greeter.SayClient) error {
	stream, err := c.LotsResponse(context.Background(), &greeter.Request{Name: "stream test"})
	if err != nil {
		return err
	}
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println(recv)
	}
	return nil
}

func main() {
	dial, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	cli := greeter.NewSayClient(dial)
	//if err = lotsRequest(cli); err != nil {
	//	log.Fatalln(err)
	//}
	if err = lotsResponse(cli); err != nil {
		log.Fatalln(err)
	}
}
