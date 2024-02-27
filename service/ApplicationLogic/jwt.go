package ApplicationLogic

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtKey = []byte("FDr1VjVQiSiybYJrQZNt8Vfd7bFEsKP6vNX1brOSiWl0mAIVCxJiR4/T3zpAlBKc2/9Lw2ac4IwMElGZkssfj3dqwa7CQC7IIB+nVxiM1c9yfowAZw4WQJ86RCUTXaXvRX8JoNYlgXcRrK3BK0E/fKCOY1+izInW3abf0jEeN40HJLkXG6MZnYdhzLnPgLL/TnIFTTAbbItxqWBtkz6FkZTG+dkDSXN7xNUxlg==")

type authClaims struct {
	jwt.StandardClaims
}

func GenerateToken() (tokenString string, err error) {
	expiresAt := time.Now().Add(10 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	})
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}

//func validateToken(tokenString string) (uint, string, error) {
//	var claims authClaims
//	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return jwtKey, nil
//	})
//	if err != nil {
//		return 0, "", err
//	}
//	if !token.Valid {
//		return 0, "", errors.New("invalid token")
//	}
//	id := claims.UserID
//	username := claims.Subject
//	return id, username, nil
//}
