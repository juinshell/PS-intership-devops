package main

import (
	"context"
	"log"
	"net/http"

	"os"
	"sync"
	"time"
	pb "grpc/proto"
	"google.golang.org/grpc"
)
const (
	defaultName = "world"
	address     = "hello-grpc-service:80"
)
var sig = 0
var mu sync.Mutex
func echo(wr http.ResponseWriter, r *http.Request) {
	mu.Lock()
	sig = 1
	mu.Unlock()
}	
func main() {
	go func(){
		http.HandleFunc("/", echo)
		err := http.ListenAndServe(":50060", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	for ; ; {
		if sig == 1 {
			// Set up a connection to the server.
			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			c := pb.NewMessageSenderClient(conn)

			// Contact the server and print out its response.
			name := defaultName
			if len(os.Args) > 1 {
				name = os.Args[1]
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			r, err := c.Send(ctx, &pb.MessageRequest{SaySomething: name})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.GetResponseSomething())
			mu.Lock()
			sig = 0
			conn.Close()
			cancel()
			mu.Unlock()
		}
	}
}
