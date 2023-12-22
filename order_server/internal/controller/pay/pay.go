package pay

import (
	"context"
	v1 "myorder/api/pay/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedPayServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPayServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.PayReq) (res *v1.PayRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
