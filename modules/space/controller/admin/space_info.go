package admin

import (
	"github.com/vera-byte/cool-admin-go/cool"
	"github.com/vera-byte/cool-admin-go/modules/space/service"
)

type SpaceInfoController struct {
	*cool.Controller
}

func init() {
	var space_info_controller = &SpaceInfoController{
		&cool.Controller{
			Perfix:  "/admin/space/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewSpaceInfoService(),
		},
	}
	// 注册路由
	cool.RegisterController(space_info_controller)
}
