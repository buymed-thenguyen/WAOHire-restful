package handler

import (
	"backend-api/config"
	"backend-api/handler/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authCfg *config.Auth) *gin.Engine {
	r := gin.New()
	r.Use(middleware.CustomRecovery())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.ResponseWrapper())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/seed", SeedData)
	r.POST("/user/sign-up", Signup)
	r.POST("/user/log-in", Login)

	user := r.Group("/user", middleware.AuthMiddleware(authCfg))
	user.GET("/me", GetMe)

	session := r.Group("/session", middleware.AuthMiddleware(authCfg))
	session.POST("/create", CreateSessionWithQuizID)
	session.POST("/:code/join", JoinSessionByCode)
	session.POST("/:code/submit", SubmitAnswer)
	session.PUT("/:code/start", StartSession)
	session.PUT("/:code/leave", LeaveSessionByCode)
	session.GET("/:code/leaderboard", GetLeaderboardBySession)
	session.GET("/:code", GetSessionDetail)
	session.GET("/:code/quiz", GetQuizDetail)
	session.GET("/:code/participants", GetSessionParticipants)
	session.GET("/:code/participants/answers", GetSessionParticipantAnswers)

	quiz := r.Group("/quiz", middleware.AuthMiddleware(authCfg))
	quiz.GET("/all", GetListQuiz)

	return r
}
