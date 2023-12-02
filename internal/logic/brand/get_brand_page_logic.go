package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/brand"
	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/agui-coder/simple-admin-product-rpc/product"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type GetBrandPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBrandPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandPageLogic {
	return &GetBrandPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBrandPageLogic) GetBrandPage(in *product.BrandPageReq) (*product.BrandListResp, error) {
	var predicate []predicate.Brand
	if in.Name != nil {
		predicate = append(predicate, brand.NameContains(*in.Name))
	}
	if in.Status != nil {
		predicate = append(predicate, brand.StatusEQ(*pointy.GetStatusPointer(in.Status)))
	}
	if len(in.CreatedAt) > 0 {
		times := make([]time.Time, len(in.CreatedAt))
		for i, c := range in.CreatedAt {
			times[i] = time.UnixMilli(c)
		}
		predicate = append(predicate, brand.CreatedAtIn(times...))
	}
	brandPage, err := l.svcCtx.DB.Brand.Query().Where(predicate...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	brandList := make([]*product.BrandInfo, len(brandPage.List))
	for i, b := range brandPage.List {
		status := uint32(b.Status)
		createdAt := b.CreatedAt.UnixMilli()
		updatedAt := b.UpdatedAt.UnixMilli()
		brandList[i] = &product.BrandInfo{
			Id:          &b.ID,
			Status:      &status,
			Name:        &b.Name,
			PicUrl:      &b.PicURL,
			Sort:        &b.Sort,
			Description: &b.Description,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		}
	}

	return &product.BrandListResp{
		Total: brandPage.PageDetails.Total,
		Data:  brandList,
	}, nil
}
