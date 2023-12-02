package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyListByIdsLogic {
	return &GetPropertyListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyListByIdsLogic) GetPropertyListByIds(in *product.PropertyListByIdsReq) (*product.PropertyListResp, error) {
	properties, err := l.svcCtx.DB.Property.Query().Where(property.IDIn(in.Id...)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	propertyList := make([]*product.PropertyInfo, len(properties))
	for i, property := range properties {
		status := uint32(property.Status)
		propertyList[i] = &product.PropertyInfo{
			Id:     &property.ID,
			Status: &status,
			Name:   &property.Name,
			Remark: &property.Remark,
		}
	}
	return &product.PropertyListResp{
		Data: propertyList,
	}, nil
}
