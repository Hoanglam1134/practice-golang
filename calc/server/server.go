package main

import (
	"context"
	"fmt"
	"log"
	"net"
	__calcpb "todolist-grpc/calc/calcpb"

	"google.golang.org/grpc"
)

type server struct {
	__calcpb.UnimplementedCalcServiceServer
}

func (*server) Sum(ctx context.Context, req *__calcpb.SumRequest) (*__calcpb.SumResponse, error) {
	resp := &__calcpb.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}
	log.Printf("Sum calc ...")
	return resp, nil
}

func (*server) Prime(ctx context.Context)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()
	__calcpb.RegisterCalcServiceServer(s, &server{})

	fmt.Printf("Server is running ...\n")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
