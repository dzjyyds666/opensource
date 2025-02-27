package sdk

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func createJwtToken(secretKey string, jsondata []byte) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(jsondata),
	})

	// 使用密钥签名并获取完整的编码令牌作为字符串
	signedString, err := claims.SignedString(secretKey)
	if nil != err {
		return "", err
	}
	return signedString, nil
}

func parseJwtToken(secretKey string, tokenString string) (string, error) {
	// 解析并验证令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if nil != err {
		return "", err
	}

	// 检查令牌是否正确
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["data"].(string), nil
	}
	return "", errors.New("token is invalid")
}
