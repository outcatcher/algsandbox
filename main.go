package main

import (
	"context"
	"fmt"
	"net"
	"time"

	p "github.com/outcatcher/algsandbox/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listenRaw(port int, ready chan<- string) {
	lsn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	defer lsn.Close()
	conn, err := lsn.Accept()
	if err != nil {
		panic(err)
	}
	buff := make([]byte, 100)
	cnt, err := conn.Read(buff)
	if err != nil {
		panic(err)
	}
	fmt.Printf("input bytes(%d): %s", cnt, string(buff))
	ready <- string(buff)
}

func main() {
	port := 8080
	readyChan := make(chan string)
	go listenRaw(port, readyChan)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Millisecond)
	cc, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := p.NewTestServiceClient(cc)
	time.Sleep(1 * time.Millisecond)
	_, _ = client.Smth(context.Background(), &emptypb.Empty{})
	fmt.Printf("Result: %s", <-readyChan)
}
