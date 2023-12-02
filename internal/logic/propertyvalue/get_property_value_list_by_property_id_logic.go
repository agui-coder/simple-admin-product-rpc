package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyValueListByPropertyIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyValueListByPropertyIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyValueListByPropertyIdLogic {
	return &GetPropertyValueListByPropertyIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyValueListByPropertyIdLogic) GetPropertyValueListByPropertyId(in *product.IDsReq) (*product.PropertyValueListResp, error) {
	propertyValues, err := l.svcCtx.DB.PropertyValue.Query().Where(propertyvalue.PropertyIDIn(in.Ids...)).All(l.ctx)
	if err != nil {
		if err != nil {
			return nil, errorhandler.DefaultEntError(l.Logger, err, in)
		}
	}
	propertyValueInfos := make([]*product.PropertyValueInfo, len(propertyValues))
	for i, propertyValue := range propertyValues {
		propertyValueInfos[i] = &product.PropertyValueInfo{
			Id:         &propertyValue.ID,
			PropertyId: &propertyValue.PropertyID,
			Name:       &propertyValue.Name,
			Remark:     &propertyValue.Remark,
		}
	}
	return &product.PropertyValueListResp{Data: propertyValueInfos}, nil
}
