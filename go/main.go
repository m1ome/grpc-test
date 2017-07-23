package main

import (
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/m1ome/grpc-test/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	counter int64
	mu      sync.Mutex
}

func (s *server) GetCounter(ctx context.Context, in *pb.Empty) (*pb.CounterReply, error) {
	// Sleeping random time
	n := time.Duration(rand.Intn(500))
	time.Sleep(time.Millisecond * n)

	// Incrementing counter
	s.mu.Lock()
	s.counter++
	cnt := s.counter
	s.mu.Unlock()

	return &pb.CounterReply{Counter: cnt}, nil
}

// This server simulates big latency stuff
// Sending all replies with a latency of 50-200ms
// It's a free toll to Ruby application
func main() {
	conn, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
