package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/entx"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePropertyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePropertyLogic {
	return &DeletePropertyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePropertyLogic) DeleteProperty(in *product.IDReq) (*product.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		count, err := l.svcCtx.DB.PropertyValue.Query().Where(propertyvalue.PropertyIDEQ(in.Id)).Count(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		if count > 0 {
			return errorx.NewInvalidArgumentError("PROPERTY_DELETE_FAIL_VALUE_EXISTS")
		}
		err = l.svcCtx.DB.Property.DeleteOneID(in.Id).Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		_, err = l.svcCtx.DB.PropertyValue.Delete().Where(propertyvalue.PropertyIDEQ(in.Id)).Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &product.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
