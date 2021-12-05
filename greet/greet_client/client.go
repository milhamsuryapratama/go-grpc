package main

import (
	"context"
	"fmt"
	"go-grpc-udemy/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect %v", err)
	}

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Println(c)

	doStream(c)

	defer conn.Close()
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ilham",
			LastName:  "Surya",
		},
	}

	response, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet %v", err)
	}

	log.Printf("response %v", response)
}

func doStream(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ilham",
			LastName:  "Surya",
		},
	}

	response, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet %v", err)
	}

	for {
		msg, err := response.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while stream %v", err)
		}

		log.Printf("response %v", msg.GetResult())
	}
}
