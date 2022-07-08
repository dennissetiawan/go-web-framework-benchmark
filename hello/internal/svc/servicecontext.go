package svc

import (
	"github.com/dennissetiawan/go-web-framework-benchmark/hello/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
