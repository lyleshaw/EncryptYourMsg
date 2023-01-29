package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/EncryptYourMsg/internal/pkg/service"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("public/*")
	router.GET("/", service.Index)
	router.POST("/encrypt", service.Encrypt)
	router.POST("/decrypt", service.Decrypt)
	err := router.Run(":3000")
	if err != nil {
		return
	}
	return
}
