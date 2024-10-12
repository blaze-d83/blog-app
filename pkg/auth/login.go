package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/blaze-d83/blog-app/env"
	"github.com/blaze-d83/blog-app/pkg/types"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetJWTConfig() string {
	jwtSecret := env.LoadConfig().JWTSecret
	return jwtSecret
}


type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := GetJWTConfig()
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error){
	claims := &Claims{}
	jwtSecret := GetJWTConfig()
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token signature")
		}
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}


func HashPassword(p string) []byte {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}
	return hashPass
}

func CompareHashPassword(admin *types.Admin, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Pass), []byte(pass)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return fmt.Errorf("incorrect password")
		}
		return fmt.Errorf("failed to compare passwords: %v", err)
	}
	return nil
}
