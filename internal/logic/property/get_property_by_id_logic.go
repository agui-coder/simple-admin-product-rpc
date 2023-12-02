package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyByIdLogic {
	return &GetPropertyByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyByIdLogic) GetPropertyById(in *product.IDReq) (*product.PropertyInfo, error) {
	property, err := l.svcCtx.DB.Property.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	status := uint32(property.Status)
	return &product.PropertyInfo{
		Id:     &property.ID,
		Status: &status,
		Name:   &property.Name,
		Remark: &property.Remark,
	}, nil
}
