package tokens

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"priyanshu.com/jwt/initializer"
)

func GenerateJwt() (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.StandardClaims{Issuer: "my-server", ExpiresAt: time.Now().AddDate(0, 0, 5).Unix()},
	)

	tokenString, err := token.SignedString(initializer.HMacSigningingSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (error) {
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return initializer.HMacSigningingSecret, nil
	})

	if err != nil {
		return err
	}

	claims := parsedToken.Claims
	if err := claims.Valid(); err != nil {
		return err
	}

	return nil
}
