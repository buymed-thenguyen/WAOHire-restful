package ws

import (
	"backend-api/config"
	pb "backend-ws/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

var grpcConn pb.BroadcasterClient

func InitGRPCClient(cfg *config.Websocket) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC: %v", err)
	}
	grpcConn = pb.NewBroadcasterClient(conn)
}
