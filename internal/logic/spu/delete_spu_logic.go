package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	"github.com/agui-coder/simple-admin-product-rpc/ent/sku"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/entx"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSpuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSpuLogic {
	return &DeleteSpuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSpuLogic) DeleteSpu(in *product.IDReq) (*product.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		err := l.svcCtx.DB.Spu.DeleteOneID(in.Id).Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		_, err = l.svcCtx.DB.Sku.Delete().Where(sku.SpuIDEQ(in.Id)).Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in.Id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &product.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
