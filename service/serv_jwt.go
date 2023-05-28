package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 私钥
var Secret = "4udLpfC0oebtbdxjYOC1O99ZvZ9cz888OyUFwNLeKilDySduFMrDl7ShGTisbEWC"
var hmacSampleSecret = []byte(Secret)

// 过期时间, 默认1小时
const TokenExpireDuration = 1 * time.Hour

type AuthClaim struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

// 生成token
func GenToken(uid string) string {
	var authClaim AuthClaim
	authClaim.UID = uid
	authClaim.StandardClaims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaim)
	tokenString, _ := token.SignedString(hmacSampleSecret) //私钥加密
	return tokenString
}

// 解析token
func Parse(tokenString string) (auth AuthClaim, Valid bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil // hmacSampleSecret为[]byte("my_secret_key")
	})
	//token是否有效, true=有效 false=无效
	Valid = token.Valid

	//返回数据
	if claims, ok := token.Claims.(jwt.MapClaims); ok && Valid {
		auth.UID = claims["uid"].(string)               //自定义的UID
		auth.ExpiresAt = int64(claims["exp"].(float64)) //过期时间
	}
	return
}
