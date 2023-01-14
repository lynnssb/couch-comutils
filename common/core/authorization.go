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
	CtxAuth               = "Authorization"
	CtxAuthKeyByJwtUserId = "JwtUserId"
)

// GeneratorJwtToken Generator Token
func GeneratorJwtToken(secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["iss"] = "lynn"
	claims[CtxAuthKeyByJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// GetAuthJwtKeyUserId 根据ctx解析token中的jwtUserId
func GetAuthJwtKeyUserId(ctx context.Context) string {
	userId := ctx.Value(CtxAuthKeyByJwtUserId).(string)
	return userId
}
