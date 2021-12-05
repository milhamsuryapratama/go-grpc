package main

import (
	"context"
	"fmt"
	"go-grpc-udemy/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect %v", err)
	}

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Println(c)

	doUnary(c)

	defer conn.Close()
}

func doUnary(c greetpb.GreetServiceClient)  {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ilham",
			LastName: "Surya",
		},
	}

	response, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet %v", err)
	}

	log.Printf("response %v", response)
}