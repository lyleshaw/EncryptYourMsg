package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/EncryptYourMsg/internal/pkg/DES"
)

type Request struct {
	Key   string `json:"key"`
	Input string `json:"input"`
}

// Index .
// @router / [GET]
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

// Encrypt .
// @router /Encrypt [POST]
func Encrypt(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ciphertext, err := DES.Encrypt([]byte(req.Input), []byte(req.Key))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.String(200, ciphertext)
}

// Decrypt .
// @router /Decrypt [POST]
func Decrypt(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	plaintext, err := DES.Decrypt(req.Input, []byte(req.Key))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.String(200, plaintext)
}
