package usecase

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/internal/domain"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ProviderSet = wire.NewSet(wire.Bind(new(domain.Usecase), new(*ConcreteAuthUsecase)), NewConcreteAuthUsecase)
var _ domain.Usecase = &ConcreteAuthUsecase{}

var (
	jwtSecret = viper.GetString("jwt.secret")
	jwtExpire = viper.GetInt("jwt.expiredSecond")
)

type ConcreteAuthUsecase struct {
}

func NewConcreteAuthUsecase() *ConcreteAuthUsecase {
	return &ConcreteAuthUsecase{}
}
