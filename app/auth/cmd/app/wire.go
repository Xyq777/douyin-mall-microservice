//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/internal/delivery"
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/internal/usecase"

	"github.com/google/wire"
)

func wireApp() *delivery.AuthDelivery {
	panic(wire.Build(
		usecase.ProviderSet,
		delivery.ProviderSet,
	))
}
