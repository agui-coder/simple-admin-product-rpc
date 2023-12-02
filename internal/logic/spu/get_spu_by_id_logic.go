package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuByIdLogic {
	return &GetSpuByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSpuByIdLogic) GetSpuById(in *product.IDReq) (*product.SpuInfo, error) {
	spu, err := l.svcCtx.DB.Spu.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	status := uint32(spu.Status)
	unit := uint32(spu.Unit)
	createdAt := spu.CreatedAt.UnixMilli()
	updatedAt := spu.UpdatedAt.UnixMilli()
	return &product.SpuInfo{
		Id:                    &spu.ID,
		Status:                &status,
		Sort:                  &spu.Sort,
		Name:                  &spu.Name,
		Keyword:               &spu.Keyword,
		Introduction:          &spu.Introduction,
		Description:           &spu.Description,
		BarCode:               &spu.BarCode,
		CategoryId:            &spu.CategoryID,
		BrandId:               &spu.BrandID,
		PicUrl:                &spu.PicURL,
		SliderPicUrls:         spu.SliderPicUrls,
		VideoUrl:              &spu.VideoURL,
		Unit:                  &unit,
		SpecType:              &spu.SpecType,
		Price:                 &spu.Price,
		MarketPrice:           &spu.MarketPrice,
		CostPrice:             &spu.CostPrice,
		Stock:                 &spu.Stock,
		DeliveryTemplateId:    &spu.DeliveryTemplateID,
		RecommendHot:          &spu.RecommendHot,
		RecommendBenefit:      &spu.RecommendBenefit,
		RecommendBest:         &spu.RecommendBest,
		RecommendNew:          &spu.RecommendNew,
		RecommendGood:         &spu.RecommendGood,
		GiveIntegral:          &spu.GiveIntegral,
		GiveCouponTemplateIds: spu.GiveCouponTemplateIds,
		SubCommissionType:     &spu.SubCommissionType,
		ActivityOrders:        spu.ActivityOrders,
		SalesCount:            &spu.SalesCount,
		VirtualSalesCount:     &spu.VirtualSalesCount,
		BrowseCount:           &spu.BrowseCount,
		CreatedAt:             &createdAt,
		UpdatedAt:             &updatedAt,
	}, nil
}
