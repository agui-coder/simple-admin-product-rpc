package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandLogic {
	return &UpdateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBrandLogic) UpdateBrand(in *product.BrandUpdateReq) (*product.BaseResp, error) {
	err := l.svcCtx.DB.Brand.UpdateOneID(in.Id).
		SetName(in.Name).
		SetPicURL(in.PicUrl).
		SetSort(in.Sort).
		SetNotNilDescription(in.Description).
		Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
