package property

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

type UpdatePropertyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePropertyLogic {
	return &UpdatePropertyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePropertyLogic) UpdateProperty(in *product.PropertyUpdateReq) (*product.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		err := l.svcCtx.DB.Property.UpdateOneID(in.Id).SetName(in.Name).SetNotNilRemark(in.Remark).Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, in)
		}
		err = l.updateSkuProperty(in.Id, in.Name)
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

func (l *UpdatePropertyLogic) updateSkuProperty(propertyValueId uint64, propertyValueName string) error {
	//TODO  @Kevin 不会有性能问题吗
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
