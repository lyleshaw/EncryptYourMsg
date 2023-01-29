package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lyleshaw/EncryptYourMsg/internal/pkg/service"
	"net/http"
)

// Handler entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.POST("/encrypt", service.Encrypt)
	router.POST("/decrypt", service.Decrypt)
	router.ServeHTTP(w, r)
}
