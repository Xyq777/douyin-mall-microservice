package core

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/pkg/log"
	logger "github.com/All-Done-Right/douyin-mall-microservice/common/mtl/log"
)

const (
	logPath   = "logs"
	logPrefix = "user"
)

func LoadLogger() {
	log.RegisterLogger(logger.NewLogrusLogger(logPath, logPrefix))
}
