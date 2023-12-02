package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/common/convert"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyValueDetailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyValueDetailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyValueDetailListLogic {
	return &GetPropertyValueDetailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyValueDetailListLogic) GetPropertyValueDetailList(in *product.IDsReq) (*product.PropertyValueDetailResp, error) {
	if len(in.Ids) < 0 {
		properties := make([]*product.SkuProperty, 0)
		return &product.PropertyValueDetailResp{SkuProperty: properties}, nil
	}
	propertyValues, err := l.svcCtx.DB.PropertyValue.Query().Where(propertyvalue.IDIn(in.Ids...)).All(l.ctx)
	if err != nil {
		if err != nil {
			return nil, errorhandler.DefaultEntError(l.Logger, err, in)
		}
	}
	if len(propertyValues) < 0 {
		properties := make([]*product.SkuProperty, 0)
		return &product.PropertyValueDetailResp{SkuProperty: properties}, nil
	}
	ids := make([]uint64, len(propertyValues))
	for i, value := range propertyValues {
		ids[i] = value.PropertyID
	}
	properties, err := l.svcCtx.DB.Property.Query().Where(property.IDIn(ids...)).All(l.ctx)
	if err != nil {
		if err != nil {
			return nil, errorhandler.DefaultEntError(l.Logger, err, in)
		}
	}

	keyMap := convert.Map(properties, func(t *ent.Property) uint64 {
		return t.ID
	})
	result := convert.SliceUpdateElement(propertyValues, func(value *ent.PropertyValue, _ int) *product.SkuProperty {
		skuProperty := &product.SkuProperty{
			ValueId:   value.ID,
			ValueName: value.Name,
		}
		// 设置属性项
		if key, found := keyMap[value.PropertyID]; found {
			skuProperty.PropertyId = key.ID
			skuProperty.PropertyName = key.Name
		}
		return skuProperty
	})
	return &product.PropertyValueDetailResp{SkuProperty: result}, nil
}
