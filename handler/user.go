package handler

import (
	"backend-api/domain"
	"backend-api/model/constant"
	reqModel "backend-api/model/request"
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req *reqModel.User
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	c.Set(constant.DATA_CTX, domain.Login(c, req))
}

func Signup(c *gin.Context) {
	var req *reqModel.User
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	c.Set(constant.DATA_CTX, domain.Signup(c, req))
}

func GetMe(c *gin.Context) {
	c.Set(constant.DATA_CTX, domain.GetMe(c))
}
