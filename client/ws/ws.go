package ws

import (
	"backend-api/config"
	pb "backend-ws/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

var grpcConn pb.BroadcasterClient

func InitGRPCClient(cfg *config.GrpcClient) {
	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC: %v", err)
	}
	grpcConn = pb.NewBroadcasterClient(conn)
}
