package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePropertyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropertyLogic {
	return &CreatePropertyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Property management
func (l *CreatePropertyLogic) CreateProperty(in *product.PropertyCreateReq) (*product.BaseIDResp, error) {
	property, err := l.svcCtx.DB.Property.Create().SetName(in.Name).SetNotNilRemark(in.Remark).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.BaseIDResp{
		Id:  property.ID,
		Msg: i18n.CreateSuccess,
	}, nil
}
