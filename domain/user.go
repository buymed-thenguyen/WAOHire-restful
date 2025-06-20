package domain

import (
	"backend-api/config"
	"backend-api/db"
	"backend-api/model/constant"
	dbModel "backend-api/model/db"
	reqModel "backend-api/model/request"
	"backend-api/model/response"
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func Login(c *gin.Context, user *reqModel.User) *response.Token {
	if user == nil {
		logger.BadRequest(c, "invalid user info")
		return nil
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	if user.Username == "" || user.Password == "" {
		logger.BadRequest(c, "missing info")
		return nil
	}

	// Check if username exists
	existsUser, err := db.GetUserByUsername(c, user.Username)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if existsUser == nil {
		logger.NotFound(c, "user not found")
		return nil
	}

	if err = bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(user.Password)); err != nil {
		logger.BadRequest(c, "wrong password")
		return nil
	}

	token, expireAt, err := config.GenerateToken(existsUser.ID)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	return &response.Token{
		Token:    token,
		ExpireAt: expireAt,
	}
}

func Signup(c *gin.Context, user *reqModel.User) *response.User {
	if user == nil {
		logger.BadRequest(c, "invalid user info")
		return nil
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.Name = strings.TrimSpace(user.Name)
	if user.Username == "" || user.Password == "" || user.Name == "" {
		logger.BadRequest(c, "missing info")
		return nil
	}

	// Check if username exists
	existsUser, err := db.GetUserByUsername(c, user.Username)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}
	if existsUser != nil {
		logger.BadRequest(c, "user already exists")
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	dbUser := dbModel.User{
		Name:     user.Name,
		Username: user.Username,
		Password: string(hash),
	}
	if err = db.CreateUser(c, &dbUser); err != nil {
		logger.InternalServerError(c, err)
		return nil
	}

	return &response.User{
		Username: user.Username,
		Name:     user.Name,
	}
}

func GetMe(c *gin.Context) *response.User {
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

	// Chuyển đổi từ db model sang response model
	return &response.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}
