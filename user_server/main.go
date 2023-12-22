package main

import (
	"my_user/internal/controller/user"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

func main() {
	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))
	c := grpcx.Server.NewConfig()
	c.Options = append(c.Options, []grpc.ServerOption{
		grpcx.Server.ChainUnary(
			grpcx.Server.UnaryValidate,
		)}...,
	)
	s := grpcx.Server.New(c)
	user.Register(s)
	s.Run()
}
