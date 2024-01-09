package service

import (
	"github.com/vera-byte/cool-admin-go/cool"
	"github.com/vera-byte/cool-admin-go/modules/demo/model"
)

type DemoSampleService struct {
	*cool.Service
}

func NewDemoSampleService() *DemoSampleService {
	return &DemoSampleService{
		&cool.Service{
			Model: model.NewDemoSample(),
		},
	}
}
