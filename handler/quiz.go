package handler

import (
	"backend-api/domain"
	"backend-api/model/constant"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetQuizDetail(c *gin.Context) {
	sessionCode := strings.TrimSpace(c.Param("code"))
	c.Set(constant.DATA_CTX, domain.GetQuizDetail(c, sessionCode))
}

func GetListQuiz(c *gin.Context) {
	c.Set(constant.DATA_CTX, domain.GetListQuiz(c))
}
