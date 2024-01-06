package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	"os"
)

var key = os.Getenv("JWT_SECRET")

type Service interface {
	GenerateToken(userID int, role string) (string, error)
	GenerateRefreshToken(userID int, role string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int, role string) (string, error) {
	expirationTime := time.Now().Add(3 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime,
	}
	log.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) GenerateRefreshToken(userID int, role string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime,
	}
	log.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
