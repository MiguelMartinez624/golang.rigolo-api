package authentication

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const BearerSchema = "Bearer"
const TokenContextKey = "token"

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}

func AuthorizeJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authHeader AuthHeader
		err := c.ShouldBindHeader(&authHeader)
		if err != nil || authHeader.Authorization == "" {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader.Authorization[len(BearerSchema):]

		tokenClaims, err := ValidateToken(tokenString)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		log.Println(tokenClaims)
		c.Set(TokenContextKey, tokenClaims)

	}
}
