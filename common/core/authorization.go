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
	CtxAuthKeyByJwtTenantId = "JwtTenantId"
	CtxAuthKeyByJwtUserId   = "JwtUserId"
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

// GetAuthJwtKeyTenantId 根据ctx解析token中的JwtTenantId
// Params:
//
//	ctx:
func GetAuthJwtKeyTenantId(ctx context.Context) string {
	tenantId := ctx.Value(CtxAuthKeyByJwtTenantId).(string)
	return tenantId
}

// GetAuthJwtKeyUserId 根据ctx解析token中的jwtTenantId
// Params:
//
//	ctx:
func GetAuthJwtKeyUserId(ctx context.Context) string {
	userId := ctx.Value(CtxAuthKeyByJwtUserId).(string)
	return userId
}

// GetAuthJwtKeyValue 根据ctx解析token中自定义的值
// Params:
//
//	ctx:
//	val: (自定义值)
func GetAuthJwtKeyValue(ctx context.Context, val string) string {
	value := ctx.Value(val).(string)
	return value
}
