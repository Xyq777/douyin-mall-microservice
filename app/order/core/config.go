package core

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/Mmx233/EnvConfig"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error(err)
		return
	}
	EnvConfig.Load("ORDER_SERVICE_", &global.Config.ServiceInfo)
	EnvConfig.Load("ORDER_CONSUL_", &global.Config.Consul)
	EnvConfig.Load("ORDER_MYSQL_", &global.Config.Mysql)
	EnvConfig.Load("ORDER_REDIS_", &global.Config.Redis)
	EnvConfig.Load("ORDER_LOGGER_", &global.Config.Logger)

}
