package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/brand"
	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandListLogic {
	return &GetBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBrandListLogic) GetBrandList(in *product.BrandListReq) (*product.BrandListResp, error) {
	var predicate []predicate.Brand
	if in.Name != nil {
		predicate = append(predicate, brand.NameContains(*in.Name))
	}
	data, err := l.svcCtx.DB.Brand.Query().Where(predicate...).All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	brandList := make([]*product.BrandInfo, len(data))
	for i, b := range data {
		status := uint32(b.Status)
		createdAt := b.CreatedAt.UnixMilli()
		brandList[i] = &product.BrandInfo{
			Id:          &b.ID,
			Status:      &status,
			Name:        &b.Name,
			PicUrl:      &b.PicURL,
			Sort:        &b.Sort,
			Description: &b.Description,
			CreatedAt:   &createdAt,
		}
	}
	return &product.BrandListResp{
		Data: brandList,
	}, nil
}
