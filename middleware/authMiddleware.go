package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status" : false,"message" : "Unauthorized1"})
		return
	}
	tokenString := strings.Replace(header,"Bearer ","",1)
	validateJwtToken(tokenString, c)
}

func validateJwtToken(tokenString string, c *gin.Context)  {
	var hmacSampleSecret []byte
	secret := os.Getenv("JWT_SECRET")
	hmacSampleSecret = []byte(secret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status" : false,"message" : "Unauthorized"})
		}
	
		return hmacSampleSecret, nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("userId", claims["userId"])
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status" : false,"message" : err.Error()})
	}

}
