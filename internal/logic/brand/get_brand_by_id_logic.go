package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBrandByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandByIdLogic {
	return &GetBrandByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBrandByIdLogic) GetBrandById(in *product.IDReq) (*product.BrandInfo, error) {
	brand, err := l.svcCtx.DB.Brand.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	status := uint32(brand.Status)
	createdAt := brand.CreatedAt.UnixMilli()
	updatedAt := brand.UpdatedAt.UnixMilli()
	return &product.BrandInfo{
		Id:          &brand.ID,
		Status:      &status,
		Name:        &brand.Name,
		PicUrl:      &brand.PicURL,
		Sort:        &brand.Sort,
		Description: &brand.Description,
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
	}, nil
}
