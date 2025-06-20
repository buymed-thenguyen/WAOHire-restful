package ws

import (
	"backend-api/utils/logger"
	pb "backend-ws/proto"
	"github.com/gin-gonic/gin"
)

func SubmitAnswerWs(c *gin.Context, userID uint, sessionCode string) error {
	_, err := grpcConn.BroadcastUserAnswered(c, &pb.UserAnsweredRequest{
		UserId:      uint32(userID),
		SessionCode: sessionCode,
	})
	if err != nil {
		logger.InternalServerError(c, err)
		return err
	}
	return nil
}

func UserJoinedWs(c *gin.Context, sessionCode string) error {
	_, err := grpcConn.BroadcastUserJoined(c, &pb.UserJoinedRequest{
		SessionCode: sessionCode,
	})
	if err != nil {
		logger.InternalServerError(c, err)
		return err
	}
	return nil
}

func UserLeavedWs(c *gin.Context, sessionCode string) error {
	_, err := grpcConn.BroadcastUserLeaved(c, &pb.UserLeavedRequest{
		SessionCode: sessionCode,
	})
	if err != nil {
		logger.InternalServerError(c, err)
		return err
	}
	return nil
}

func StartSessionWs(c *gin.Context, quizID uint, sessionCode string) error {
	_, err := grpcConn.BroadcastStartSession(c, &pb.StartSessionRequest{
		SessionCode: sessionCode,
		QuizId:      uint32(quizID),
	})
	if err != nil {
		logger.InternalServerError(c, err)
		return err
	}
	return nil
}
