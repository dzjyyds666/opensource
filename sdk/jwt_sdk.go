package sdk

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	ErrTokenIsInvalid = errors.New("token is invalid")
)

func CreateJwtToken(secretKey string, jsondata []byte) (string, error) {
	claims := jwt.MapClaims{
		"data": string(jsondata),
		"exp":  time.Now().Add(time.Hour * 24 * 3).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  "learnx",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseJwtToken(secretKey string, tokenString string) (*jwt.MapClaims, error) {
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是相同的签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	// 检查解析结果
	if err != nil {
		return nil, err
	}

	// 验证 Token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
