package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePropertyValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePropertyValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePropertyValueLogic {
	return &DeletePropertyValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePropertyValueLogic) DeletePropertyValue(in *product.IDReq) (*product.BaseResp, error) {
	err := l.svcCtx.DB.PropertyValue.DeleteOneID(in.Id).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
