package main

import (
	v1 "my_user/api/user/v1"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	grpcx.Resolver.Register(etcd.New("127.0.0.1:2379"))

	var (
		ctx  = gctx.New()
		conn = grpcx.Client.MustNewGrpcClientConn("user", grpcx.Client.ChainUnary(
			grpcx.Client.UnaryTracing,
		))
		client = v1.NewUserClient(conn)
	)

	res, err := client.GetOne(ctx, &v1.GetOneReq{Id: 1})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, "Response:", res.User)
}
