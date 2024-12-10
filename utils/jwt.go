package utils

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Name string
	jwt.RegisteredClaims
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, true)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, true)
	}
}

// GenerateToken 生成jwt签名
func GenerateToken(name string) (string, error) {
	claims := CustomClaims{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "codepzj is handsome",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("codepzj is handsome"))
}

func ParseToken(jwtStr string) (jwt.Claims, error) {
	byteKey := []byte("codepzj is handsome")
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (any, error) {
		return byteKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("令牌过期")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("断言失败")
	}
	return claims, nil
}

// ClearToken 清除Token
func ClearToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, true)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, true)
	}
}
