package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/vera-byte/cool-admin-go/cool-tools/internal/service"
)

var (
	Install = gcmd.Command{
		Name:  "install",
		Usage: "cool-tools install",
		Brief: "Install cool-tools to the system.",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			err = service.Install.Run(ctx)
			return
		},
	}
)

// init
func init() {
	Main.AddCommand(&Install)
}
