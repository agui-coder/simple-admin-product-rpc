package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/spu"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSpuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuListLogic {
	return &GetSpuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSpuListLogic) GetSpuList(in *product.SpuListReq) (*product.SpuListResp, error) {
	var predicate []predicate.Spu
	if in.Name != nil {
		predicate = append(predicate, spu.NameContains(*in.Name))
	}
	if in.Keyword != nil {
		predicate = append(predicate, spu.KeywordContains(*in.Keyword))
	}
	if in.Introduction != nil {
		predicate = append(predicate, spu.IntroductionContains(*in.Introduction))
	}

	spuPageList, err := l.svcCtx.DB.Spu.Query().Where(predicate...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	spuList := make([]*product.SpuInfo, len(spuPageList.List))
	for i, s := range spuPageList.List {
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
	return &product.SpuListResp{
		Total: spuPageList.PageDetails.Total,
		Data:  spuList,
	}, nil
}
