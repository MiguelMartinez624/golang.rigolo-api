package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtSecret string

//jwt service
type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthClaims struct {
	UserID     string `json:"user_id"`
	jwt.StandardClaims
}

type DefaultJWTServices struct {
	secretKey string
	issure    string
}

//SetSecret default value for secret "secret"
func SetSecret(secret string) {
	if secret == "" {
		jwtSecret = "secret"
	}
	jwtSecret = os.Getenv("SECRET")
}

// validateIDToken returns the token's claims if the token is valid
func ValidateToken(tokenString string) (*AuthClaims, error) {
	claims := &AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	// For now we'll just return the error and handle logging in service level
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidTokenID
	}

	claims, ok := token.Claims.(*AuthClaims)

	if !ok {
		return nil, ErrCouldNotParseClains
	}

	return claims, nil
}
