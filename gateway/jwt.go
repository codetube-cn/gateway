package gateway

import (
	"codetube.cn/gateway/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//UserJwtClaims 用户 JWT 声明
type UserJwtClaims struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	jwt.StandardClaims
}

// ParseJwt 解析 JWT
func ParseJwt(j string) (*jwt.Token, *UserJwtClaims, error) {
	claims := &UserJwtClaims{}
	token, err := jwt.ParseWithClaims(j, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GatewayConfig.JwtKey), nil
	})
	return token, claims, err
}
