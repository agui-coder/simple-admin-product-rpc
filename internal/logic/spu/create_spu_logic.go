package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/common/consts"
	"github.com/agui-coder/simple-admin-product-rpc/common/convert"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	entType "github.com/agui-coder/simple-admin-product-rpc/ent/schema/enttype"
	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/entx"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/agui-coder/simple-admin-product-rpc/product"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSpuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSpuLogic {
	return &CreateSpuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Spu management
func (l *CreateSpuLogic) CreateSpu(in *product.SpuCreateReq) (*product.BaseIDResp, error) {
	var spuId uint64
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		giveCouponTemplateIds := convert.SliceUpdateElement(in.GiveCouponTemplates, func(t *product.GiveCouponTemplate, _ int) uint64 {
			return t.Id
		})
		giveCouponTemplateIds = convert.SliceUniq(giveCouponTemplateIds, func(t uint64) uint64 {
			return t
		})
		spu, err := l.svcCtx.DB.Spu.Create().
			SetName(in.Name).
			SetKeyword(in.Keyword).
			SetIntroduction(in.Introduction).
			SetDescription(in.Description).
			SetCategoryID(in.CategoryId).
			SetBrandID(in.BrandId).
			SetPicURL(in.PicUrl).
			SetNotNilSliderPicUrls(&in.SliderPicUrls).
			SetNotNilVideoURL(in.VideoUrl).
			SetUnit(uint8(in.Unit)).
			SetSort(in.Sort).
			SetDeliveryTemplateID(in.DeliveryTemplateId).
			SetRecommendHot(in.RecommendHot).
			SetRecommendBenefit(in.RecommendBenefit).
			SetRecommendBest(in.RecommendBest).
			SetRecommendNew(in.RecommendNew).
			SetRecommendGood(in.RecommendGood).
			SetGiveIntegral(in.GiveIntegral).
			SetGiveCouponTemplateIds(giveCouponTemplateIds).
			SetSubCommissionType(in.SubCommissionType).
			SetActivityOrders(in.ActivityOrders).
			SetNotNilPrice(convert.GetMinValue(in.Skus, func(t *product.SkuCreateOrUpdateReq) *int32 {
				return &t.Price
			})).
			SetNotNilMarketPrice(convert.GetMinValue(in.Skus, func(t *product.SkuCreateOrUpdateReq) *int32 {
				return t.MarketPrice
			})).SetNotNilCostPrice(convert.GetMinValue(in.Skus, func(t *product.SkuCreateOrUpdateReq) *int32 {
			return t.CostPrice
		})).
			SetStock(convert.GetSumValue(in.Skus, func(t *product.SkuCreateOrUpdateReq) int32 {
				return t.Stock
			}, convert.Int32Sum)).
			SetNotNilBarCode(in.BarCode).
			SetStatus(consts.Disable).
			SetSalesCount(int32(0)).
			SetBrowseCount(int32(0)).Save(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		bulk := make([]*ent.SkuCreate, len(in.Skus))
		for i, skus := range in.Skus {
			properties := skus.Properties
			skuProperties := make([]entType.SkuProperty, len(properties))
			for i2, property := range properties {
				skuProperty := entType.SkuProperty{
					PropertyId:   property.PropertyId,
					PropertyName: property.PropertyName,
					ValueId:      property.ValueId,
					ValueName:    property.ValueName,
				}
				skuProperties[i2] = skuProperty
			}
			bulk[i] = l.svcCtx.DB.Sku.Create().
				SetSpuID(spu.ID).
				SetPrice(skus.Price).
				SetNotNilMarketPrice(skus.MarketPrice).
				SetNotNilCostPrice(skus.CostPrice).
				SetNotNilBarCode(skus.BarCode).
				SetPicURL(skus.PicUrl).
				SetStock(skus.Stock).SetNotNilWeight(skus.Weight).
				SetNotNilVolume(skus.Volume).SetNotNilProperties(&skuProperties)

		}
		_, err = l.svcCtx.DB.Sku.CreateBulk(bulk...).Save(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in.Skus)
		}
		spuId = spu.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &product.BaseIDResp{
		Id:  spuId,
		Msg: i18n.CreateSuccess,
	}, nil
}
