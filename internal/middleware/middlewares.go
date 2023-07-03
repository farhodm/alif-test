package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/farhodm/ewallet/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
)

// AuthMiddleware is a middleware function to authenticate requests
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-UserID")
		digest := c.GetHeader("X-Digest")

		// Get user from the database using userID
		user := models.User{}
		result := db.First(&user, "id =?", userID)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// Verify the request body digest
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Generate HMAC-SHA1 digest from the request body
		hc := hmac.New(sha1.New, []byte(user.Password))
		hc.Write(body)
		expectedDigest := hex.EncodeToString(hc.Sum(nil))

		// Compare digests
		if digest != expectedDigest {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Continue processing the request
		c.Next()
	}
}
