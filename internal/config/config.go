package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	rest.RestConf
	ListenOn string
	Consul   consul.Conf
}
