package token

import (
	"context"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var ctx = context.Background()

func init() {
	os.Setenv("ACCESS_SECRET", "secretsecret")
	fmt.Println("init function!")
}

type AccessDetails struct {
	Email string
}

func CreateToken(email string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println("is going well")
		fmt.Println(os.Getenv("ACCESS_SECRET"))
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenEmail(tokenString string) (*AccessDetails, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			Email: email,
		}, nil
	}
	return nil, err
}
