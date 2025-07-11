package handler

import (
	"backend-api/domain"
	"backend-api/model/constant"
	reqModel "backend-api/model/request"
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
	"strings"
)

func Login(c *gin.Context) {
	var req *reqModel.User
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	c.Set(constant.DATA_CTX, domain.Login(c, req))
}

func Signup(c *gin.Context) {
	var req *reqModel.User
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	req.Name = strings.TrimSpace(req.Name)

	c.Set(constant.DATA_CTX, domain.Signup(c, req))
}

func GetMe(c *gin.Context) {
	c.Set(constant.DATA_CTX, domain.GetMe(c))
}
