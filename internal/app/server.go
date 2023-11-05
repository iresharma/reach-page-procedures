package app

import (
	pkg "Reach-page/internal/pkg"
	"Reach-page/internal/pkg/RPC/pb"
	"Reach-page/internal/pkg/database"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":4001"
)

func Main() {
	database.DB = database.CreateConnection()
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to setup tcp server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPagePackageServer(grpcServer, &pkg.PageServer{})
	log.Printf("Server started on this %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err)
	}
}
