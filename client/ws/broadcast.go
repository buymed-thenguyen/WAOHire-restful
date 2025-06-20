package ws

import (
	pb "backend-ws/proto"
	"github.com/gin-gonic/gin"
	"log"
)

func SubmitAnswerWs(c *gin.Context, userID uint, sessionCode string) {
	_, err := grpcConn.BroadcastUserAnswered(c, &pb.UserAnsweredRequest{
		UserId:      uint32(userID),
		SessionCode: sessionCode,
	})
	if err != nil {
		log.Printf("gRPC call failed: %v", err)
	}
}

func UserJoinedWs(c *gin.Context, sessionCode string) {
	_, err := grpcConn.BroadcastUserJoined(c, &pb.UserJoinedRequest{
		SessionCode: sessionCode,
	})
	if err != nil {
		log.Printf("gRPC call failed: %v", err)
	}
}

func UserLeavedWs(c *gin.Context, sessionCode string) {
	_, err := grpcConn.BroadcastUserLeaved(c, &pb.UserLeavedRequest{
		SessionCode: sessionCode,
	})
	if err != nil {
		log.Printf("gRPC call failed: %v", err)
	}
}

func StartSessionWs(c *gin.Context, quizID uint, sessionCode string) {
	_, err := grpcConn.BroadcastStartSession(c, &pb.StartSessionRequest{
		SessionCode: sessionCode,
		QuizId:      uint32(quizID),
	})
	if err != nil {
		log.Printf("gRPC call failed: %v", err)
	}
}
