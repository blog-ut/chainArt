// Package token
// Time : 2023/8/17 23:17
// Author : PTJ
// File : token
// Software: GoLand
package token

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var secret = []byte("It is my secret")

const TokenExpireDuration = 2 * time.Hour

type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	jwt.RegisteredClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenToken 生成token
func GenToken(username string, userId uint) (string, error) {
	c := Claims{
		Username: username,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration * time.Duration(1))), // 过期时间12小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                             // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                             // 生效时间
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*Claims, error) {
	s := strings.Fields(tokenStr)[1]
	token, err := jwt.ParseWithClaims(s, &Claims{}, func(tk *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claim, ok := token.Claims.(*Claims); ok && token.Valid {
			return claim, err
		}
	}
	return nil, err

}

// VerifyToken 验证token
func VerifyToken(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}

// GetRand
// 生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
