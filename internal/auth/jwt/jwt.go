package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/ChandraWahyuR/be-latihan_mkp/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	signKey string
}

type JWTCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewJwt(signKey string) *JWTService {
	return &JWTService{
		signKey: signKey,
	}
}

func (j *JWTService) GenerateToken(data *model.Login) (string, error) {
	claims := JWTCustomClaims{
		ID:    data.ID,
		Email: data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "soal_mkp",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.signKey))
	if err != nil {
		return "", fmt.Errorf("error when signing :%s", err)
	}

	return tokenString, nil
}

func (j *JWTService) VerifyToken(tokenString string) (*JWTCustomClaims, error) {
	claims := &JWTCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.signKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
