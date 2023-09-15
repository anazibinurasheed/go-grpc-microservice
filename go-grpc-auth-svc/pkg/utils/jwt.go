package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/anazibinurasheed/go-grpc-microservice/go-grpc-auth-svc/pkg/models"
	"github.com/golang-jwt/jwt"
)

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
}

// Token auth methods implemented on JwtWrapper
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

func (w *JwtWrapper) GenerateToken(user models.User) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		})

	if err != nil {
		fmt.Println("u1")
		return
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		fmt.Println("u2")
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		fmt.Println("u3")
		return nil, errors.New("JWT is expired")

	}

	return claims, nil

}
