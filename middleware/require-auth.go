package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No token provided"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			log.Fatal(err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Check if the token is valid
			if exp, ok := claims["exp"].(float64); ok {
				if exp < float64(time.Now().Unix()) {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Token expired"})
					return
				}

				// Get the user from the sub claim
				userIDFloat, ok := claims["sub"].(float64)
				if !ok {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid token"})
					return
				}
				userID := uint(userIDFloat)
				ctx.Set("userID", userID)

			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid token"})
				return
			}

			ctx.Next()
		}
	}
}
