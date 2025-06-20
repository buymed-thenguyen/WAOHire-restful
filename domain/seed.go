package domain

import (
	"backend-api/db"
	"backend-api/model/response"
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func SeedData(c *gin.Context) *response.DefaultResponse {
	if err := db.SeedQuizzesFullSet(); err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	return &response.DefaultResponse{Message: "ok"}
}
