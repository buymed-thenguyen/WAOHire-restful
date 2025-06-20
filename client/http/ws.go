package http

import (
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
)

type BroadcastRequest struct {
	UserID      uint   `json:"user_id"`
	SessionCode string `json:"session_code"`
	QuizID      uint   `json:"quiz_id"`
}

func (c *WSClient) UserJoinedWs(ctx *gin.Context, sessionCode string) error {
	req := &BroadcastRequest{
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/user-joined", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		return err
	}
	return nil
}

func (c *WSClient) UserLeavedWs(ctx *gin.Context, sessionCode string) error {
	req := &BroadcastRequest{
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/user-leaved", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		return err
	}
	return nil
}

func (c *WSClient) UserAnsweredWs(ctx *gin.Context, userID uint, sessionCode string) error {
	req := &BroadcastRequest{
		UserID:      userID,
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/user-answered", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		return err
	}
	return nil
}

func (c *WSClient) StartSessionWs(ctx *gin.Context, quizID uint, sessionCode string) error {
	req := &BroadcastRequest{
		QuizID:      quizID,
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/start-session", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		return err
	}
	return nil
}
