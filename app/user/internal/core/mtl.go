package core

import (
	"github.com/All-Done-Right/douyin-mall-microservice/common/mtl"
	"github.com/spf13/viper"
)

func StartMtl() {
	mtl.InitTracing(viper.GetString("service.name"))
	mtl.InitMetric(viper.GetString("service.name"),
		viper.GetString("mtl.port"),
		viper.GetString("consul.address"))
}
