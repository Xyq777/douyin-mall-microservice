package core

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/pkg/log"
	"github.com/All-Done-Right/douyin-mall-microservice/common/mtl"
	logger "github.com/All-Done-Right/douyin-mall-microservice/common/mtl/log"
	"github.com/spf13/viper"
)

const (
	logPath   = "logs"
	logPrefix = "user"
)

func StartMtl() {
	log.RegisterLogger(logger.NewLogrusLogger(logPath, logPrefix))
	mtl.InitTracing(viper.GetString("service.name"))
	mtl.InitMetric(viper.GetString("service.name"),
		viper.GetString("core.address"),
		viper.GetString("consul.address"))
}
