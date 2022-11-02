package main

import (
	"context"
	"log"
	__calcpb "practice-golang/calc/calcpb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	defer cc.Close()

	client := __calcpb.NewCalcServiceClient(cc)
	callSum(client)
	log.Printf("service client %v", client)
}

func callSum(c __calcpb.CalcServiceClient) {
	log.Printf("calling service ...")
	resp, err := c.Sum(context.Background(), &__calcpb.SumRequest{
		Num1: 10,
		Num2: 4,
	})

	if err != nil {
		log.Fatalf("Error when calc sum: %v", err)
	}

	log.Printf("sum of api call: %v", resp.GetResult())
}
