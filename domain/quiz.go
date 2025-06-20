package domain

import (
	"backend-api/db"
	dbModel "backend-api/model/db"
	"backend-api/model/response"
	"backend-api/utils"
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func GetListQuiz(c *gin.Context) []*response.Quiz {
	quizzes, err := db.GetListQuiz(c)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	return utils.MapSlice(quizzes, func(q *dbModel.Quiz) *response.Quiz {
		return &response.Quiz{
			ID:            q.ID,
			Title:         q.Title,
			TotalQuestion: len(q.Questions),
		}
	})
}

func GetQuizDetail(c *gin.Context, sessionCode string) *response.Quiz {
	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
		return nil
	}
	if session.StartAt == nil || session.StartAt.After(time.Now()) {
		logger.BadRequest(c, "session not start yet")
		return nil
	}

	quiz, err := db.GetQuizByID(c, session.QuizID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if quiz == nil {
		logger.NotFound(c, "quiz not found")
		return nil
	}

	questions, err := db.GetQuestionsFromQuizID(c, session.QuizID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if len(questions) == 0 {
		logger.NotFound(c, "questions not found")
		return nil
	}

	for _, q := range questions {
		// Shuffle options
		rand.Shuffle(len(q.Options), func(i, j int) {
			q.Options[i], q.Options[j] = q.Options[j], q.Options[i]
		})
	}

	resp := response.Quiz{
		ID:    quiz.ID,
		Title: quiz.Title,
	}
	resp.Questions = utils.MapSlice(questions, func(q *dbModel.Question) *response.Question {
		return &response.Question{
			ID:           q.ID,
			QuestionText: q.QuestionText,
			AnswerOptions: utils.MapSlice(q.Options, func(a *dbModel.AnswerOption) *response.AnswerOption {
				return &response.AnswerOption{
					ID:   a.ID,
					Text: a.Text,
				}
			}),
		}
	})
	return &resp
}
