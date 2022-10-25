package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type User struct {
	Name string
	// Age      int
	// Password string
	Data map[string]interface{}
}

// type ISecretFunc interface {
// 	GetSecret() []byte
// }

func (u *User) GetSecret() []byte {
	return []byte(u.Name + "secret")
}

type UserClaims struct {
	User
	jwt.StandardClaims
}

func GenerateToken(user User, expires time.Duration) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	userClaims := UserClaims{User: user, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(expires).Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(user.GetSecret())

	return tokenString, errors.WithStack(err)
}

func AuthToken(tokenString string) (*UserClaims, error) {
	jwtParser := new(jwt.Parser)
	claims := UserClaims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if uc, ok := token.Claims.(*UserClaims); ok {
			return uc.GetSecret(), nil
		}
		return nil, errors.Errorf("token claim is not UserClaims")
	}
	token, err := jwtParser.ParseWithClaims(tokenString, &claims, keyFunc)
	if err != nil {
		return &claims, err
	}

	if _, ok := token.Claims.(*UserClaims); !(ok && token.Valid) {
		return nil, errors.Errorf("token claim is not UserClaims or token is not Valid")
	}
	return &claims, nil
}
func main() {
	// user := User{Name: "chen", Age: 25, Password: "matrix"}
	// token, err := GenerateToken(user, time.Hour*24*7)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(token)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyVHlwZSI6MSwiVXNlcklEIjowLCJBY2NvdW50IjoiMTMyMTM4MDIzNzMiLCJVc2VyTmFtZSI6IkljaGVuZyIsIk1vYmlsZSI6IiIsIlNsYXQiOiIyMzczIiwiT3JnS2V5IjoiZGV2LW9yZyIsImV4cCI6MTY2NTEyNzM4NywiaXNzIjoic3BlY2RpcyJ9.A5AsVCSTmLM-Qe8DkOKJwSY1WZ4ojhCKL9rB0KX3NdA"
	userClaims, err := AuthToken(token)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(userClaims.User)
}
