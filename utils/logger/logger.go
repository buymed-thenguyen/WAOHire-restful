package logger

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func InitLogger(path string) {
	var err error
	LogFile, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("missing log file: %v", err)
	}

	Logger = log.New(io.MultiWriter(LogFile, os.Stdout), "", log.Ldate|log.Ltime|log.Lshortfile)
}

func ErrLog(c *gin.Context, status int, err error) {
	if err == nil {
		_ = c.Error(nil)
		c.Status(http.StatusOK)
		return
	}

	Logger.Printf("[GIN] %d | %s", status, err.Error())
	_ = c.Error(err)
	c.Status(status)
}

func InternalServerError(c *gin.Context, err error) {
	_ = c.Error(err)
	Logger.Printf("[GIN] %d | %s", http.StatusInternalServerError, err.Error())
	c.Status(http.StatusInternalServerError)
}

func BadRequest(c *gin.Context, msg string) {
	_ = c.Error(errors.New(msg))
	Logger.Printf("[GIN] %d | %s", http.StatusBadRequest, msg)
	c.Status(http.StatusBadRequest)
}

func Unauthorized(c *gin.Context) {
	_ = c.Error(errors.New("unauthorized"))
	Logger.Printf("[GIN] %d | %s", http.StatusUnauthorized, "unauthorized")
	c.Status(http.StatusUnauthorized)
}

func NotFound(c *gin.Context, msg string) {
	_ = c.Error(errors.New(msg))
	Logger.Printf("[GIN] %d | %s", http.StatusNotFound, errors.New(msg))
	c.Status(http.StatusNotFound)
}
