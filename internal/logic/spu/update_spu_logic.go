package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/common/convert"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	entType "github.com/agui-coder/simple-admin-product-rpc/ent/schema/enttype"
	"github.com/agui-coder/simple-admin-product-rpc/ent/sku"
	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/entx"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/agui-coder/simple-admin-product-rpc/product"
	"github.com/suyuan32/simple-admin-common/i18n"
	"sort"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpuLogic {
	return &UpdateSpuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSpuLogic) UpdateSpu(in *product.SpuUpdateReqVO) (*product.BaseResp, error) {

	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		_, err := l.svcCtx.DB.Spu.UpdateOneID(in.Id).
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
			SetStatus(uint8(1)).
			SetSalesCount(int32(0)).
			SetBrowseCount(int32(0)).Save(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}

		err = l.updateSkuList(in.Id, in.Skus)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &product.BaseResp{Msg: i18n.UpdateSuccess}, nil
}

func (l *UpdateSpuLogic) updateSkuList(spuId uint64, skus []*product.SkuCreateOrUpdateReq) error {
	existsSku, err := l.svcCtx.DB.Sku.Query().Where(sku.SpuIDEQ(spuId)).All(l.ctx)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, spuId)
	}
	existsSkuMap := convert.MapValues(existsSku, buildPropertyKey, func(t *ent.Sku) uint64 {
		return t.ID
	})
	var createBulks []*ent.SkuCreate
	var updateBulks []*ent.SkuUpdate
	allUpdateSkus := convertList(skus, spuId)
	for _, s := range allUpdateSkus {
		propertiesKey := buildPropertyKey(s)
		// 找得到的，进行更新
		existsSkuId, exists := existsSkuMap[propertiesKey]
		delete(existsSkuMap, propertiesKey)
		if exists {
			s.ID = existsSkuId
			updateBulks = append(updateBulks, l.svcCtx.DB.Sku.Update().Where(sku.IDEQ(existsSkuId)).
				SetPrice(s.Price).
				SetNotNilMarketPrice(&s.MarketPrice).
				SetNotNilCostPrice(&s.CostPrice).
				SetNotNilBarCode(&s.BarCode).
				SetPicURL(s.PicURL).
				SetStock(s.Stock).SetNotNilWeight(&s.Weight).
				SetNotNilVolume(&s.Volume).
				SetNotNilProperties(&s.Properties))
			break
		}
		// 找不到，进行插入
		s.SpuID = spuId
		createBulks = append(createBulks, l.svcCtx.DB.Sku.Create().
			SetPrice(s.Price).
			SetNotNilMarketPrice(&s.MarketPrice).
			SetNotNilCostPrice(&s.CostPrice).
			SetNotNilBarCode(&s.BarCode).
			SetPicURL(s.PicURL).
			SetStock(s.Stock).SetNotNilWeight(&s.Weight).
			SetNotNilVolume(&s.Volume).
			SetNotNilProperties(&s.Properties))
	}

	_, err = l.svcCtx.DB.Sku.CreateBulk(createBulks...).Save(l.ctx)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, createBulks)
	}

	for _, updateBulk := range updateBulks {
		err = updateBulk.Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, updateBulk)
		}
	}
	var skuIds []uint64
	for _, id := range existsSkuMap {
		skuIds = append(skuIds, id)
	}
	_, err = l.svcCtx.DB.Sku.Delete().Where(sku.IDIn(skuIds...)).Exec(l.ctx)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, skuIds)
	}
	return nil
}

func buildPropertyKey(t *ent.Sku) string {
	properties := t.Properties
	if len(properties) == 0 {
		return ""
	}

	// 将属性列表按照 ValueId 排序
	sort.Slice(properties, func(i, j int) bool {
		return properties[i].ValueId < properties[j].ValueId
	})

	// 将排序后的 ValueId 拼接成字符串
	var valueIds []string
	for _, p := range properties {
		valueIds = append(valueIds, strconv.FormatUint(p.ValueId, 10))
	}

	// 使用逗号将字符串拼接起来
	return strings.Join(valueIds, ",")
}

func convertList(list []*product.SkuCreateOrUpdateReq, spuId uint64) []*ent.Sku {
	var result []*ent.Sku
	for _, reqVO := range list {
		properties := make([]entType.SkuProperty, len(reqVO.Properties))
		for i, property := range reqVO.Properties {
			skuProperty := entType.SkuProperty{
				PropertyId:   property.PropertyId,
				PropertyName: property.PropertyName,
				ValueId:      property.ValueId,
				ValueName:    property.ValueName,
			}
			properties[i] = skuProperty
		}

		sku := &ent.Sku{
			SpuID:       spuId,
			Properties:  properties,
			Price:       reqVO.Price,
			MarketPrice: *reqVO.MarketPrice,
			CostPrice:   *reqVO.CostPrice,
			BarCode:     *reqVO.BarCode,
			PicURL:      reqVO.PicUrl,
			Stock:       reqVO.Stock,
			Weight:      *reqVO.Weight,
			Volume:      *reqVO.Volume,
		}

		result = append(result, sku)
	}

	return result
}
