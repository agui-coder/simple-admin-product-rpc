package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePropertyValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePropertyValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropertyValueLogic {
	return &CreatePropertyValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PropertyValue management
func (l *CreatePropertyValueLogic) CreatePropertyValue(in *product.PropertyValueCreateReq) (*product.BaseIDResp, error) {
	dbValue := l.svcCtx.DB.PropertyValue.Query().Where(propertyvalue.PropertyIDEQ(in.PropertyId), propertyvalue.NameEQ(in.Name)).FirstX(l.ctx)
	if dbValue != nil {
		return &product.BaseIDResp{Id: dbValue.ID, Msg: "属性值已存在"}, nil
	}
	propertyValue, err := l.svcCtx.DB.PropertyValue.Create().
		SetPropertyID(in.PropertyId).
		SetName(in.Name).
		SetNotNilRemark(in.Remark).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.BaseIDResp{
		Id:  propertyValue.ID,
		Msg: i18n.CreateSuccess,
	}, nil
}
