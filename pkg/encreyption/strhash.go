package encreyption

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HashPassword(c * gin.Context,pass string) string  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
			return ""
	}
	return string(hashedPassword)
}