package handler

import (
	"backend-api/domain"
	"backend-api/model/constant"
	reqModel "backend-api/model/request"
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func CreateSessionWithQuizID(c *gin.Context) {
	var req *reqModel.CreateSession
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	c.Set(constant.DATA_CTX, domain.CreateSessionWithQuizID(c, req.QuizID))
}

func JoinSessionByCode(c *gin.Context) {
	sessionCode := c.Param("code")
	c.Set(constant.DATA_CTX, domain.JoinSessionByCode(c, sessionCode))
}

func LeaveSessionByCode(c *gin.Context) {
	sessionCode := c.Param("code")
	c.Set(constant.DATA_CTX, domain.LeaveSessionByCode(c, sessionCode))
}

func GetLeaderboardBySession(c *gin.Context) {
	sessionCode := c.Param("code")
	c.Set(constant.DATA_CTX, domain.GetLeaderboardBySession(c, sessionCode))
}

func SubmitAnswer(c *gin.Context) {
	sessionCode := c.Param("code")
	var req *reqModel.SubmitAnswer
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	c.Set(constant.DATA_CTX, domain.SubmitAnswer(c, sessionCode, req))
}

func StartSession(c *gin.Context) {
	sessionCode := c.Param("code")
	c.Set(constant.DATA_CTX, domain.StartSession(c, sessionCode))
}

func GetSessionDetail(c *gin.Context) {
	sessionCode := c.Param("code")
	c.Set(constant.DATA_CTX, domain.GetSessionDetail(c, sessionCode))
}

func GetSessionParticipants(c *gin.Context) {
	sessionCode := c.Param("code")
	c.Set(constant.DATA_CTX, domain.GetSessionParticipants(c, sessionCode))
}

func GetSessionParticipantAnswers(c *gin.Context) {
	sessionCode := c.Param("code")
	c.Set(constant.DATA_CTX, domain.GetSessionParticipantAnswers(c, sessionCode))
}
