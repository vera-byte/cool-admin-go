package main

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/example/rpc/grpcx/basic/controller"
)

func main() {
	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))
	s := grpcx.Server.New()
	controller.Register(s)
	s.Run()
}
