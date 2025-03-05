package main

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/product/biz/dal"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/sirupsen/logrus"
	"log"
	"net"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dal.Init()

	r, err := consul.NewConsulRegister("consul:8500")
	if err != nil {
		klog.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", ":8887")
	if err != nil {
		logrus.Fatal(err)
	}
	svr := product.NewServer(new(ProductCatalogServiceImpl),
		server.WithServiceAddr(addr),
		//指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "product",
			},
		),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
