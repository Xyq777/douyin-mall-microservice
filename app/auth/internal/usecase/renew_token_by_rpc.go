package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Run create note info
func (u *ConcreteAuthUsecase) RenewToken(ctx context.Context, oldToken string) (tokenStr string, err error) {
	// Finish your business logic.
	// 1. 从请求中获取旧的 Token

	// 2. 解析旧的 JWT Token
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	// 3. 提取旧 Token 中的用户 ID
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to extract claims from token")
	}
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return "", errors.New("failed to extract user_id from claims")
	}

	// 4. 生成新的 JWT Token
	newClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(jwtExpire)), // 过期时间
		"iat":     time.Now().Unix(),                                    // 创建时间
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err = newToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	// 5. 返回新的 JWT Token
	return tokenStr, nil
}
