package domain

import (
	"backend-api/db"
	"backend-api/model/constant"
	"backend-api/model/request"
	"backend-api/model/response"
	"backend-api/utils"
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strings"
	"time"

	dbModel "backend-api/model/db"
)

func CreateSessionWithQuizID(c *gin.Context, quizID uint) *response.Session {
	userID := c.GetUint(constant.USER_ID_CTX) // get user_id from jwt
	if userID == 0 {
		logger.Unauthorized(c)
		return nil
	}

	session := &dbModel.Session{
		QuizID:    quizID,
		Code:      generateSessionCode(constant.SESSION_CODE_LEN),
		CreatedBy: userID,
	}
	if err := db.CreateSession(c, session); err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	participant := &dbModel.Participant{
		UserID:    userID,
		QuizID:    quizID,
		SessionID: session.ID,
	}
	if err := db.CreateParticipant(c, participant); err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	return &response.Session{
		ID:        session.ID,
		QuizID:    quizID,
		Code:      session.Code,
		CreatedBy: userID,
	}
}

func generateSessionCode(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return strings.ToUpper(string(b))
}

func JoinSessionByCode(c *gin.Context, sessionCode string) *response.Participant {
	userID := c.GetUint(constant.USER_ID_CTX) // get user_id from jwt
	if userID == 0 {
		logger.Unauthorized(c)
		return nil
	}
	user, err := db.GetUsersByID(c, userID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if user == nil {
		logger.NotFound(c, "user not found")
		return nil
	}

	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
		return nil
	}
	if session.StartAt != nil && session.StartAt.Before(time.Now()) {
		logger.BadRequest(c, "session already started")
		return nil
	}

	// Check if participant already exists
	participant, err := db.GetParticipantByUserIDSessionID(c, userID, session.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if participant != nil {
		log.Printf("user already joined")
		return &response.Participant{
			UserID:    userID,
			QuizID:    session.QuizID,
			SessionID: session.ID,
			UserName:  user.Name,
		}
	}

	participant = &dbModel.Participant{
		UserID:    userID,
		QuizID:    session.QuizID,
		SessionID: session.ID,
	}
	if err = db.CreateParticipant(c, participant); err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	//ws.UserJoinedWs(c, session.Code)
	d.WsHttpClient.UserJoinedWs(c, session.Code)

	return &response.Participant{
		UserID:    userID,
		QuizID:    session.QuizID,
		SessionID: session.ID,
		UserName:  user.Name,
	}
}

func LeaveSessionByCode(c *gin.Context, sessionCode string) *response.Participant {
	userID := c.GetUint(constant.USER_ID_CTX)
	if userID == 0 {
		logger.Unauthorized(c)
		return nil
	}
	user, err := db.GetUsersByID(c, userID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if user == nil {
		logger.NotFound(c, "user not found")
		return nil
	}

	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
		return nil
	}
	if session.StartAt != nil && session.StartAt.Before(time.Now()) {
		logger.BadRequest(c, "session already started")
		return nil
	}

	// Check if participant already exists
	participant, err := db.GetParticipantByUserIDSessionID(c, userID, session.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if participant == nil {
		logger.BadRequest(c, "user not join yet")
		return nil
	}

	participant = &dbModel.Participant{
		UserID:    userID,
		QuizID:    session.QuizID,
		SessionID: session.ID,
	}
	if err = db.RemoveParticipant(c, participant); err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	//ws.UserLeavedWs(c, session.Code)
	d.WsHttpClient.UserLeavedWs(c, session.Code)
	return &response.Participant{
		UserID:    userID,
		QuizID:    session.QuizID,
		SessionID: session.ID,
		UserName:  user.Name,
	}
}

func GetLeaderboardBySession(c *gin.Context, sessionCode string) []*response.Participant {
	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
		return nil
	}

	rows, err := db.GetSessionLeaderboard(c, session.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	userIDs := utils.MapSlice(rows, func(p *dbModel.Participant) uint {
		return p.UserID
	})
	users, err := db.GetUsersByIDs(c, userIDs)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	userMap := utils.SliceToMap(users, func(u *dbModel.User) uint {
		return u.ID
	})

	return utils.MapSlice(rows, func(r *dbModel.Participant) *response.Participant {
		var name string
		if user, exists := userMap[r.UserID]; exists {
			name = user.Name
		}
		return &response.Participant{
			UserID:       r.UserID,
			QuizID:       r.QuizID,
			SessionID:    r.SessionID,
			TotalScore:   r.TotalScore,
			DoneAt:       r.DoneAt,
			TimeConsumed: r.TimeConsumed,
			UserName:     name,
		}
	})
}

func SubmitAnswer(c *gin.Context, sessionCode string, req *request.SubmitAnswer) *response.DefaultResponse {
	userID := c.GetUint(constant.USER_ID_CTX) // get user_id from jwt
	if userID == 0 {
		logger.Unauthorized(c)
		return nil
	}

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

	participant, err := db.GetParticipantByUserIDSessionID(c, userID, session.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if participant == nil {
		logger.NotFound(c, "participant not found")
		return nil
	}
	if participant.DoneAt != nil && !participant.DoneAt.IsZero() {
		logger.BadRequest(c, "participant already done")
		return nil
	}

	totalScore, correctAnswerMap, err := calculateScore(c, session, req.Answers)
	if err != nil {
		return nil
	}

	participant.DoneAt = utils.ToPointerTime(time.Now())
	participant.TotalScore = totalScore
	participant.TimeConsumed = participant.DoneAt.Sub(*participant.CreatedAt).Milliseconds()

	err = db.UpdateParticipant(c, participant)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	participantAnswers := utils.MapSlice(req.Answers, func(a *request.Answer) *dbModel.ParticipantAnswer {
		isCorrect := correctAnswerMap[a.QuestionId].ID == a.AnswerOptionId
		return &dbModel.ParticipantAnswer{
			ParticipantID:  participant.ID,
			QuestionID:     a.QuestionId,
			AnswerOptionID: a.AnswerOptionId,
			SessionID:      session.ID,
			IsCorrect:      isCorrect,
		}
	})
	err = db.CreateParticipantAnswers(c, participantAnswers)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	//ws.SubmitAnswerWs(c, userID, sessionCode)
	d.WsHttpClient.UserAnsweredWs(c, userID, sessionCode)

	return &response.DefaultResponse{
		Message: "ok",
	}
}

func calculateScore(c *gin.Context, session *dbModel.Session, answers []*request.Answer) (int, map[uint]*dbModel.AnswerOption, error) {
	questions, err := db.GetQuestionsFromQuizID(c, session.QuizID)
	if err != nil {
		logger.InternalServerError(c, err)
		return 0, nil, nil
	}
	questionMap := utils.SliceToMap(questions, func(q *dbModel.Question) uint {
		return q.ID
	})

	questionIDs := utils.MapSlice(questions, func(q *dbModel.Question) uint {
		return q.ID
	})
	correctAnswers, err := db.GetCorrectAnswerByQuestionIDs(c, questionIDs)
	if err != nil {
		logger.InternalServerError(c, err)
		return 0, nil, nil
	}

	correctAnswersMap := utils.SliceToMap(correctAnswers, func(a *dbModel.AnswerOption) uint {
		return a.QuestionID
	})

	var totalScore int
	for _, a := range answers {
		if a == nil {
			continue
		}
		question, exists := questionMap[a.QuestionId]
		if !exists {
			continue
		}
		if correctAnswersMap[a.QuestionId].ID == a.AnswerOptionId {
			totalScore += question.Score
		}
	}

	return totalScore, correctAnswersMap, nil
}

func StartSession(c *gin.Context, sessionCode string) *response.DefaultResponse {
	userID := c.GetUint(constant.USER_ID_CTX) // get user_id from jwt
	if userID == 0 {
		logger.Unauthorized(c)
		return nil
	}

	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
		return nil
	}
	if session.CreatedBy != userID {
		logger.BadRequest(c, "session not created by user")
		return nil
	}
	if session.StartAt != nil && session.StartAt.Before(time.Now()) {
		logger.BadRequest(c, "session already started")
		return nil
	}

	session.StartAt = utils.ToPointerTime(time.Now())
	if err = db.UpdateSession(c, session); err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	//ws.StartSessionWs(c, userID, sessionCode)
	d.WsHttpClient.StartSessionWs(c, userID, sessionCode)
	return &response.DefaultResponse{
		Message: "ok",
	}
}

func GetSessionDetail(c *gin.Context, sessionCode string) *response.Session {
	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
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

	return &response.Session{
		ID:        session.ID,
		QuizID:    session.QuizID,
		Code:      sessionCode,
		CreatedBy: session.CreatedBy,
		Quiz: &response.Quiz{
			ID:            quiz.ID,
			Title:         quiz.Title,
			TotalQuestion: len(quiz.Questions),
		},
	}
}

func GetSessionParticipants(c *gin.Context, sessionCode string) []*response.Participant {
	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
		return nil
	}

	participants, err := db.GetParticipantBySessionID(c, session.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	userIDs := utils.MapSlice(participants, func(p *dbModel.Participant) uint {
		return p.UserID
	})
	users, err := db.GetUsersByIDs(c, userIDs)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	userMap := utils.SliceToMap(users, func(u *dbModel.User) uint {
		return u.ID
	})

	return utils.MapSlice(participants, func(p *dbModel.Participant) *response.Participant {
		var userName string
		if user, exists := userMap[p.UserID]; exists {
			userName = user.Name
		}
		return &response.Participant{
			QuizID:       p.QuizID,
			SessionID:    p.SessionID,
			UserID:       p.UserID,
			UserName:     userName,
			TotalScore:   p.TotalScore,
			TimeConsumed: p.TimeConsumed,
		}
	})
}

func GetSessionParticipantAnswers(c *gin.Context, sessionCode string) []*response.ParticipantAnswer {
	userID := c.GetUint(constant.USER_ID_CTX) // get user_id from jwt
	if userID == 0 {
		logger.Unauthorized(c)
		return nil
	}

	session, err := db.GetSessionByCode(c, sessionCode)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if session == nil {
		logger.NotFound(c, "session not found")
		return nil
	}

	participant, err := db.GetParticipantByUserIDSessionID(c, userID, session.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if participant == nil {
		logger.NotFound(c, "participant not found")
		return nil
	}

	participantAnswers, err := db.GetParticipantAnswersByParticipantID(c, participant.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	return utils.MapSlice(participantAnswers, func(a *dbModel.ParticipantAnswer) *response.ParticipantAnswer {
		return &response.ParticipantAnswer{
			ID:             a.ID,
			QuestionID:     a.QuestionID,
			AnswerOptionID: a.AnswerOptionID,
			CreatedAt:      a.CreatedAt,
		}
	})
}
