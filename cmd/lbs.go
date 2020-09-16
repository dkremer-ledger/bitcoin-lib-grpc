package main

import (
	"fmt"
	"net"

	controllers "github.com/ledgerhq/bitcoin-svc/grpc"
	"github.com/ledgerhq/bitcoin-svc/log"
	pb "github.com/ledgerhq/bitcoin-svc/pb/bitcoin"
	"google.golang.org/grpc"
)

func serve() {
	addr := fmt.Sprintf(":%d", 50051)

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cannot listen to address %s", addr)
	}

	s := grpc.NewServer()
	bitcoinController := controllers.NewBitcoinController()
	pb.RegisterCoinServiceServer(s, bitcoinController)

	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	serve()
}
