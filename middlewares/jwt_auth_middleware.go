package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware(c *gin.Context) {
	ts, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	type Claims struct {
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(ts, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		c.Next()

		return
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
