package sku

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/sku"
	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSkuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSkuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSkuListLogic {
	return &GetSkuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSkuListLogic) GetSkuList(in *product.IDsReq) (*product.SkuListResp, error) {
	skus, err := l.svcCtx.DB.Sku.Query().Where(sku.IDIn(in.Ids...)).All(l.ctx)
	if err != nil {
		errorhandler.DefaultEntError(l.Logger, err, in)
	}
	skuInfos := make([]*product.SkuInfo, len(skus))
	for i, s := range skus {
		skuPropertyList := make([]*product.SkuProperty, len(s.Properties))
		for i, property := range s.Properties {
			skuPropertyList[i] = &product.SkuProperty{
				PropertyId:   property.PropertyId,
				PropertyName: property.PropertyName,
				ValueId:      property.ValueId,
				ValueName:    property.ValueName,
			}
		}
		createdAt := s.CreatedAt.UnixMilli()
		updatedAt := s.UpdatedAt.UnixMilli()
		skuInfos[i] = &product.SkuInfo{
			Id:          &s.ID,
			SpuId:       &s.SpuID,
			Properties:  skuPropertyList,
			Price:       &s.Price,
			MarketPrice: &s.MarketPrice,
			CostPrice:   &s.CostPrice,
			BarCode:     &s.BarCode,
			PicUrl:      &s.PicURL,
			Stock:       &s.Stock,
			Weight:      &s.Weight,
			Volume:      &s.Volume,
			SalesCount:  &s.SalesCount,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		}
	}
	return &product.SkuListResp{
		Data: skuInfos,
	}, nil
}
