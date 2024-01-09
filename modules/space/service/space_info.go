package service

import (
	"github.com/vera-byte/cool-admin-go/cool"
	"github.com/vera-byte/cool-admin-go/modules/space/model"
)

type SpaceInfoService struct {
	*cool.Service
}

func NewSpaceInfoService() *SpaceInfoService {
	return &SpaceInfoService{
		&cool.Service{
			Model: model.NewSpaceInfo(),
		},

		// Service: cool.NewService(model.NewSpaceInfo()),
	}
}
