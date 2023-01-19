/**
 * @author:       wangxuebing
 * @fileName:     authorization.go
 * @date:         2023/1/14 16:33
 * @description:
 */

package core

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

const (
	CtxAuth = "Authorization"
)
const (
	CtxAuthKeyByJwtUserId = "JwtUserId" //default
)

// GeneratorJwtToken Generator Token
// Params:
//
//	val: (默认可以传:CtxAuthKeyByJwtUserId <如需其他可以自定义>)
//	secretKey:
//	iat:
//	seconds:
//	userId:
func GeneratorJwtToken(val string, secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["iss"] = "lynn"
	claims[val] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// GetAuthJwtKeyUserId 根据ctx解析token中的jwtUserId
// Params:
//
//	ctx:
//	val: (默认可以传:CtxAuthKeyByJwtUserId <如需其他可以自定义>)
func GetAuthJwtKeyUserId(ctx context.Context, val string) string {
	userId := ctx.Value(val).(string)
	return userId
}
