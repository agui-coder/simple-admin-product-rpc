package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyValuePageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyValuePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyValuePageLogic {
	return &GetPropertyValuePageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyValuePageLogic) GetPropertyValuePage(in *product.PropertyValuePageReq) (*product.PropertyValueListResp, error) {
	propertyValuePage, err := l.svcCtx.DB.PropertyValue.Query().Where().Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	propertyValueInfos := make([]*product.PropertyValueInfo, len(propertyValuePage.List))
	for i, propertyValue := range propertyValuePage.List {
		propertyValueInfos[i] = &product.PropertyValueInfo{
			Id:         &propertyValue.ID,
			PropertyId: &propertyValue.PropertyID,
			Name:       &propertyValue.Name,
			Remark:     &propertyValue.Remark,
		}
	}
	return &product.PropertyValueListResp{Total: propertyValuePage.PageDetails.Total, Data: propertyValueInfos}, nil
}
