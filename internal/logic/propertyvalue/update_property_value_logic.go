package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"github.com/agui-coder/simple-admin-product-rpc/ent/sku"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/entx"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePropertyValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePropertyValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePropertyValueLogic {
	return &UpdatePropertyValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePropertyValueLogic) UpdatePropertyValue(in *product.PropertyValueUpdateReq) (*product.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		dbValue, err := l.svcCtx.DB.PropertyValue.Query().Where(propertyvalue.PropertyIDEQ(in.PropertyId), propertyvalue.NameEQ(in.Name)).Only(l.ctx)
		if err != nil && !ent.IsNotFound(err) {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		if dbValue != nil && dbValue.ID != in.Id {
			return errorx.NewInvalidArgumentError("PROPERTY_VALUE_EXISTS")
		}
		err = l.svcCtx.DB.PropertyValue.UpdateOneID(in.Id).
			SetPropertyID(in.PropertyId).
			SetName(in.Name).
			SetNotNilRemark(in.Remark).
			Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		err = l.updateSkuPropertyValue(in.Id, in.Name)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &product.BaseResp{Msg: i18n.UpdateSuccess}, nil
}

func (l *UpdatePropertyValueLogic) updateSkuPropertyValue(propertyValueId uint64, propertyValueName string) error {
	//TODO  不会有性能问题吗
	skus, err := l.svcCtx.DB.Sku.Query().All(l.ctx)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, propertyValueId)
	}
	if len(skus) == 0 {
		return nil
	}
	var updateBulks []*ent.SkuUpdate
	for _, s := range skus {
		if s.Properties != nil {
			for idx, property := range s.Properties {
				if property.PropertyId == propertyValueId {
					s.Properties[idx].PropertyName = propertyValueName
					updateBulks = append(updateBulks, l.svcCtx.DB.Sku.Update().Where(sku.IDEQ(s.ID)).
						SetNotNilProperties(&s.Properties))
					break
				}
			}
		}
	}
	if len(updateBulks) == 0 {
		return nil
	}
	for _, updateBulk := range updateBulks {
		err = updateBulk.Exec(l.ctx)
		if err != nil {
			errorhandler.DefaultEntError(l.Logger, err, updateBulk)
		}
	}
	return nil
}
