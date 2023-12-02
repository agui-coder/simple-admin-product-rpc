package sku

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkuByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkuByIdLogic {
	return &GetSkuByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkuByIdLogic) GetSkuById(in *product.IDReq) (*product.SkuInfo, error) {
	sku, err := l.svcCtx.DB.Sku.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	skuPropertyList := make([]*product.SkuProperty, len(sku.Properties))
	for i, property := range sku.Properties {
		skuPropertyList[i] = &product.SkuProperty{
			PropertyId:   property.PropertyId,
			PropertyName: property.PropertyName,
			ValueId:      property.ValueId,
			ValueName:    property.ValueName,
		}
	}
	return &product.SkuInfo{
		Id:          &sku.ID,
		SpuId:       &sku.SpuID,
		Properties:  skuPropertyList,
		Price:       &sku.Price,
		MarketPrice: &sku.MarketPrice,
		CostPrice:   &sku.CostPrice,
		BarCode:     &sku.BarCode,
		PicUrl:      &sku.PicURL,
		Stock:       &sku.Stock,
		Weight:      &sku.Weight,
		Volume:      &sku.Volume,
		SalesCount:  &sku.SalesCount,
	}, nil
}
