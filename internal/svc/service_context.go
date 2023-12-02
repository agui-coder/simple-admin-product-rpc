package svc

import (
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	"github.com/agui-coder/simple-admin-product-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"

	//需要导入runtime
	_ "github.com/agui-coder/simple-admin-product-rpc/ent/runtime"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
