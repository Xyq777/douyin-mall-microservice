package dal

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
