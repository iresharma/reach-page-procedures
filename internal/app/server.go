package app

import (
	"Reach-page/internal/pkg/RPC/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":3000"
)

type PageServer struct {
	pb.PagePackageServer
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to setup tcp server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPagePackageServer(grpcServer, &PageServer{})
	log.Printf("Server started on this %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err)
	}
}
