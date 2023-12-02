package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/common/convert"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	"github.com/agui-coder/simple-admin-product-rpc/ent/spu"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateSpuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateSpuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateSpuListLogic {
	return &ValidateSpuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateSpuListLogic) ValidateSpuList(in *product.IDsReq) (*product.SpuListResp, error) {
	if len(in.Ids) == 0 {
		return &product.SpuListResp{}, nil
	}
	list, err := l.svcCtx.DB.Spu.Query().Where(spu.IDIn(in.Ids...)).All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	spuMap := convert.Map(list, func(t *ent.Spu) uint64 {
		return t.ID
	})
	for _, id := range in.Ids {
		spu, exist := spuMap[id]
		if !exist {
			return nil, err
		}
		if spu.Status != 1 {
			return nil, err
		}
	}
	spuList := make([]*product.SpuInfo, len(list))
	for i, s := range list {
		status := uint32(s.Status)
		unit := uint32(s.Unit)
		spuList[i] = &product.SpuInfo{
			Id:                    &s.ID,
			Status:                &status,
			Sort:                  &s.Sort,
			Name:                  &s.Name,
			Keyword:               &s.Keyword,
			Introduction:          &s.Introduction,
			Description:           &s.Description,
			BarCode:               &s.BarCode,
			CategoryId:            &s.CategoryID,
			BrandId:               &s.BrandID,
			PicUrl:                &s.PicURL,
			SliderPicUrls:         s.SliderPicUrls,
			VideoUrl:              &s.VideoURL,
			Unit:                  &unit,
			SpecType:              &s.SpecType,
			Price:                 &s.Price,
			MarketPrice:           &s.MarketPrice,
			CostPrice:             &s.CostPrice,
			Stock:                 &s.Stock,
			DeliveryTemplateId:    &s.DeliveryTemplateID,
			RecommendHot:          &s.RecommendHot,
			RecommendBenefit:      &s.RecommendBenefit,
			RecommendBest:         &s.RecommendBest,
			RecommendNew:          &s.RecommendNew,
			RecommendGood:         &s.RecommendGood,
			GiveIntegral:          &s.GiveIntegral,
			GiveCouponTemplateIds: s.GiveCouponTemplateIds,
			SubCommissionType:     &s.SubCommissionType,
			ActivityOrders:        s.ActivityOrders,
			SalesCount:            &s.SalesCount,
			VirtualSalesCount:     &s.VirtualSalesCount,
			BrowseCount:           &s.BrowseCount,
		}
	}
	return &product.SpuListResp{Data: spuList}, nil
}
