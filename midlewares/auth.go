package midlewares

import (
	"golang_basic_gin_sept_2023/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Request need access token",
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		// validate token
		email, role, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		// set email
		c.Set("x-email", email)
		c.Set("x-role", role)

		// next
		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Request need access token",
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		// validate token
		email, role, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		if role != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Role Anda Tidak Dapat Mengkases Enpoint Ini",
				"status":  http.StatusUnauthorized,
			})

			c.Abort()
			return
		}

		// set email
		c.Set("x-email", email)
		c.Set("x-role", role)

		// next
		c.Next()
	}
}
