package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendJsonError(c *gin.Context, status int, message string, err error) {
	msg := fmt.Sprintf("%v: %v", message, err)
	LoggerError(msg)

	c.JSON(status, gin.H{"message": msg})
}

func SendJson(c *gin.Context, status int, obj interface{}) {
	c.JSON(status, obj)
}