package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/agui-coder/simple-admin-product-rpc/product"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBrandLogic {
	return &CreateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateBrand  management
func (l *CreateBrandLogic) CreateBrand(in *product.BrandCreateReq) (*product.BaseIDResp, error) {
	brand, err := l.svcCtx.DB.Brand.Create().SetName(in.Name).
		SetSort(in.Sort).
		SetPicURL(in.PicUrl).
		SetNotNilDescription(in.Description).
		SetStatus(*pointy.GetStatusPointer(&in.Status)).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.BaseIDResp{
		Id:  brand.ID,
		Msg: i18n.CreateSuccess,
	}, nil
}
