package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyValueByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyValueByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyValueByIdLogic {
	return &GetPropertyValueByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyValueByIdLogic) GetPropertyValueById(in *product.IDReq) (*product.PropertyValueInfo, error) {
	propertyValue, err := l.svcCtx.DB.PropertyValue.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.PropertyValueInfo{
		Id:         &propertyValue.ID,
		PropertyId: &propertyValue.PropertyID,
		Name:       &propertyValue.Name,
		Remark:     &propertyValue.Remark,
	}, nil
}
