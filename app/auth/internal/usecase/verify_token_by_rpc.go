package usecase

import (
	"context"
	"errors"
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/pkg/log"
	"github.com/golang-jwt/jwt/v4"
)

func (u *ConcreteAuthUsecase) VerifyToken(ctx context.Context, token string) (isValid bool, err error) {

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		log.Log().Error(err)
		return false, nil
	})
	if err != nil {
		return false, err
	}

	// 检查令牌是否有效
	if !t.Valid {
		return true, errors.New("token is invalid")
	}

	return true, nil
}
