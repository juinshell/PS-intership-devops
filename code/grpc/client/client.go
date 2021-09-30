package main

import (
	"context"
	"log"
	"net/http"

	pb "grpc/proto"
	"os"
	"time"

	"google.golang.org/grpc"
)
const (
	defaultName = "world"
	address     = "hello-grpc-server-service:80"
	//address     = ":50053"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageSenderClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rr, err := c.Send(ctx, &pb.MessageRequest{SaySomething: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rr.GetResponseSomething())
}	
func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":50060", nil)
	if err != nil {
		log.Fatal(err)
	}	
}
