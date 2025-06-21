package domain

import (
	"backend-api/db"
	"backend-api/model/response"
	"github.com/gin-gonic/gin"
)

func SeedData(c *gin.Context) *response.DefaultResponse {
	if err := db.SeedQuizzesFullSet(c); err != nil {
		return nil
	}
	return &response.DefaultResponse{Message: "ok"}
}
