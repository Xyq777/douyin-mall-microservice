//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/internal/core"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/internal/delivery"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/internal/repo"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/internal/usecase"

	"github.com/google/wire"
)

func wireApp() *delivery.UserDelivery {
	panic(wire.Build(
		core.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		delivery.ProviderSet,
	))
}
