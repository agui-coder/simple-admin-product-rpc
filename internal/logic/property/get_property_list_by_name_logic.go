package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyListByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyListByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyListByNameLogic {
	return &GetPropertyListByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyListByNameLogic) GetPropertyListByName(in *product.PropertyListByNameReq) (*product.PropertyListResp, error) {
	var predicate []predicate.Property
	if in.Name != nil {
		predicate = append(predicate, property.NameContains(*in.Name))
	}
	properties, err := l.svcCtx.DB.Property.Query().Where(predicate...).All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
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
