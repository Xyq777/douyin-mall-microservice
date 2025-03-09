package usecase

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/internal/domain"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(wire.Bind(new(domain.UserUsecase), new(*ConcreteUserUsecase)), NewConcreteUserUsecase)
