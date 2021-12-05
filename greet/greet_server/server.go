package main

import (
	"context"
	"fmt"
	"go-grpc-udemy/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Hello From Greet")
	firstName := req.GetGreeting().GetFirstName()
	//lastName := req.GetGreeting().GetLastName()

	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main()  {
	fmt.Println("Hello World from Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed %v", err)
	}
}
