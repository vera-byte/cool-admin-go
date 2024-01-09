package demo

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/vera-byte/cool-admin-go/modules/space/controller"
	_ "github.com/vera-byte/cool-admin-go/modules/space/middleware"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module space init start ...")
	g.Log().Debug(ctx, "module space init finished ...")
}
