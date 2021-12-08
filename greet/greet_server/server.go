package main

import (
	"context"
	"fmt"
	"go-grpc-udemy/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Hello From Greet Stream")
	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 100; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		stream.SendMsg(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// stop reading client stream request
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}

		firtName := req.GetGreeting().GetFirstName()
		result += "Hello " + firtName + "! "
	}
}

func main() {
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
