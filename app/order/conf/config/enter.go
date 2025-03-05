package config

type Config struct {
	Mysql       Mysql
	Consul      Consul
	Redis       Redis
	ServiceInfo ServiceInfo
	Logger      Logger
}
