package db

import (
	dbModel "backend-api/model/db"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateParticipant(c *gin.Context, participant *dbModel.Participant) error {
	return DB.WithContext(c.Request.Context()).Create(&participant).Error
}

func CreateParticipantTx(c *gin.Context, participant *dbModel.Participant, tx *gorm.DB) error {
	return tx.WithContext(c.Request.Context()).Create(&participant).Error
}

func RemoveParticipant(c *gin.Context, participant *dbModel.Participant) error {
	return DB.WithContext(c.Request.Context()).
		Where("session_id = ?", participant.SessionID).
		Where("user_id = ?", participant.UserID).
		Delete(&participant).Error
}

func RemoveParticipantTx(c *gin.Context, participant *dbModel.Participant, tx *gorm.DB) error {
	return tx.WithContext(c.Request.Context()).
		Where("session_id = ?", participant.SessionID).
		Where("user_id = ?", participant.UserID).
		Delete(&participant).Error
}

func GetParticipantByUserIDSessionID(c *gin.Context, userID uint, sessionID uint) (*dbModel.Participant, error) {
	var participant *dbModel.Participant
	err := DB.WithContext(c.Request.Context()).
		Where("user_id = ?", userID).
		Where("session_id = ?", sessionID).
		First(&participant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return participant, nil
}

func GetParticipantBySessionID(c *gin.Context, sessionID uint) ([]*dbModel.Participant, error) {
	var participants []*dbModel.Participant
	err := DB.WithContext(c.Request.Context()).
		Where("session_id = ?", sessionID).
		Find(&participants).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func UpdateParticipant(c *gin.Context, participant *dbModel.Participant) error {
	return DB.WithContext(c.Request.Context()).Save(participant).Error
}

func UpdateParticipantTx(c *gin.Context, participant *dbModel.Participant, tx *gorm.DB) error {
	return tx.WithContext(c.Request.Context()).Save(participant).Error
}

func GetSessionLeaderboard(c *gin.Context, sessionID uint) ([]*dbModel.Participant, error) {
	var rows []*dbModel.Participant
	err := DB.WithContext(c.Request.Context()).
		Table("participants").
		Where("participants.session_id = ?", sessionID).
		Order("participants.total_score DESC, participants.time_consumed ASC").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}
