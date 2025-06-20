package db

import (
	dbModel "backend-api/model/db"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetSessionByCode(c *gin.Context, code string) (*dbModel.Session, error) {
	var existing *dbModel.Session
	err := DB.WithContext(c.Request.Context()).
		Where("code = ?", code).
		First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func GetSessionByCodeForUpdate(c *gin.Context, code string, tx *gorm.DB) (*dbModel.Session, error) {
	var existing *dbModel.Session
	err := tx.WithContext(c.Request.Context()).
		Where("code = ?", code).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func CreateSession(c *gin.Context, session *dbModel.Session) error {
	return DB.WithContext(c.Request.Context()).Create(&session).Error
}

func UpdateSession(c *gin.Context, session *dbModel.Session) error {
	return DB.WithContext(c.Request.Context()).Save(session).Error
}

func UpdateSessionTx(c *gin.Context, session *dbModel.Session, tx *gorm.DB) error {
	return tx.WithContext(c.Request.Context()).Save(session).Error
}
