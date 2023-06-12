package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ikadgzl/inventory/authsvc/internal/config"
)

type JwtUtil interface {
	GenerateToken(id uint64) (string, error)
	ValidateToken(token string) (*jwtClaims, error)
}

type jwtUtil struct {
	jwtConfig *config.JWTConfig
}

func NewJwtUtil(jwtConfig *config.JWTConfig) JwtUtil {
	return &jwtUtil{
		jwtConfig: jwtConfig,
	}
}

type jwtClaims struct {
	*jwt.StandardClaims
	Id uint64
}

func (j *jwtUtil) GenerateToken(id uint64) (string, error) {
	sClaims := &jwt.StandardClaims{
		ExpiresAt: j.jwtConfig.ExpiresAt,
		Issuer:    j.jwtConfig.Issuer,
	}
	claims := &jwtClaims{
		StandardClaims: sClaims,
		Id:             id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.jwtConfig.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func (j *jwtUtil) ValidateToken(token string) (*jwtClaims, error) {
	validatedToken, err := jwt.ParseWithClaims(token, &jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.jwtConfig.Secret), nil
		})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := validatedToken.Claims.(*jwtClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token is expired")
	}

	return claims, nil
}
