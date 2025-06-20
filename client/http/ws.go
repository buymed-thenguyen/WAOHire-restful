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

func (c *WSClient) UserJoinedWs(ctx *gin.Context, sessionCode string) {
	req := &BroadcastRequest{
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/user-joined", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		// continue
	}
}

func (c *WSClient) UserLeavedWs(ctx *gin.Context, sessionCode string) {
	req := &BroadcastRequest{
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/user-leaved", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		// continue
	}
}

func (c *WSClient) UserAnsweredWs(ctx *gin.Context, userID uint, sessionCode string) {
	req := &BroadcastRequest{
		UserID:      userID,
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/user-answered", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		// continue
	}
}

func (c *WSClient) StartSessionWs(ctx *gin.Context, quizID uint, sessionCode string) {
	req := &BroadcastRequest{
		QuizID:      quizID,
		SessionCode: sessionCode,
	}
	err := c.Post("/ws/start-session", req)
	if err != nil {
		logger.InternalServerError(ctx, err)
		// continue
	}
}
