package ws

import (
	"backend-api/config"
	pb "backend-ws/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

var grpcConn pb.BroadcasterClient

func InitGRPCClient(cfg *config.GrpcClient) {
	var opts []grpc.DialOption

	if os.Getenv("ENV") == "prd" {
		creds := credentials.NewClientTLSFromCert(nil, "")
		opts = append(opts, grpc.WithTransportCredentials(creds))
		log.Println("gRPC client using TLS (production mode)")
	} else {
		opts = append(opts, grpc.WithInsecure())
		log.Println("gRPC client using insecure connection (non-production mode)")
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		opts...,
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC: %v", err)
	}

	grpcConn = pb.NewBroadcasterClient(conn)
}
