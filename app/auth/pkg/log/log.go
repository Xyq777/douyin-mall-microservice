package log

import logger "github.com/All-Done-Right/douyin-mall-microservice/common/mtl/log"

var log logger.Logger

func RegisterLogger(logger logger.Logger) {
	log = logger
}
func Log() logger.Logger {
	if log == nil {
		panic("implement not found for interface Logger, please register")
	}
	return log
}
