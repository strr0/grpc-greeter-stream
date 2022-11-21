package main

import (
	"google.golang.org/grpc"
	"greeter/proto"
	"io"
	"log"
	"net"
)

type Say struct {
	greeter.UnimplementedSayServer
}

func (Say) LotsRequest(stream greeter.Say_LotsRequestServer) error {
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greeter.Response{Msg: "recv over..."})
		}
		if err != nil {
			return err
		}
		log.Println(recv)
	}
}

func (Say) LotsResponse(req *greeter.Request, stream greeter.Say_LotsResponseServer) error {
	log.Println(req)
	for i := 0; i < 5; i++ {
		if err := stream.Send(&greeter.Response{Msg: "hello~"}); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	srv := grpc.NewServer()
	greeter.RegisterSayServer(srv, new(Say))
	if err := srv.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
