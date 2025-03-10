package usecase

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// Run create note info
func (u *ConcreteAuthUsecase) DeliverToken(ctx context.Context, userID int32) (tokenStr string, err error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(jwtExpire) * time.Second), // 过期时间
		"iat":     time.Now().Unix(),                                      // 创建时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
