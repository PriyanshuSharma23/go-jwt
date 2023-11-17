package users

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"priyanshu.com/jwt/constants"
)

type customClaims struct {
	jwt.StandardClaims
	User
}

func GenerateJwt(user User) (string, error) {
	claims := customClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "my-server",
			ExpiresAt: time.Now().AddDate(0, 0, 5).Unix(),
		},
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(constants.HMacSigningingSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*User, error) {
	parsedToken, err := jwt.ParseWithClaims(
		tokenString,
		&customClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return constants.HMacSigningingSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*customClaims)
	if err := claims.Valid(); err != nil || !ok {
		return nil, fmt.Errorf("Token not valid")
	}

	return &claims.User, nil
}
