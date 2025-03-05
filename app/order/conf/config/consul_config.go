package config

import "strconv"

type Consul struct {
	Host string
	Port int
}

func (e Consul) Addr() string {
	return e.Host + ":" + strconv.Itoa(e.Port)
}
