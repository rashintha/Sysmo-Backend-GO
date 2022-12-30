package jwtUtil

import (
	"fmt"
	"log"
	"sysmo/lib/env"
	usersService "sysmo/services/users"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(user *usersService.User) string {
	claims := &jwt.MapClaims{
		"iss": "Rosetta Software",
		"exp": jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
		"nbf": jwt.NewNumericDate(time.Now()),
		"iat": jwt.NewNumericDate(time.Now()),
		"data": map[string]string{
			"id":    fmt.Sprintf("%d", user.ID),
			"email": user.Email,
			"name":  user.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(env.CONF["JWT_ACCESS_SECRET"]))

	if err != nil {
		log.Fatalln("Generating JWT token failed.")
	}
	return signed
}

func ValidateAccessToken(tokenEncoded string) (*jwt.Token, error) {
	return jwt.Parse(tokenEncoded, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		return []byte(env.CONF["JWT_ACCESS_SECRET"]), nil
	})
}

func GenerateRefreshToken(user *usersService.User) string {
	claims := &jwt.MapClaims{
		"iss": "Rosetta Software",
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		"nbf": jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		"iat": jwt.NewNumericDate(time.Now()),
		"data": map[string]string{
			"id":    fmt.Sprintf("%d", user.ID),
			"email": user.Email,
			"name":  user.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(env.CONF["JWT_REFRESH_SECRET"]))

	if err != nil {
		log.Fatalln("Generating JWT refresh token failed.")
	}
	return signed
}

func ValidateRefreshToken(tokenEncoded string) (*jwt.Token, error) {
	return jwt.Parse(tokenEncoded, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		return []byte(env.CONF["JWT_REFRESH_SECRET"]), nil
	})
}
