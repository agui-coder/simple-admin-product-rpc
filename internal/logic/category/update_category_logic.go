package category

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/common/consts"
	"github.com/agui-coder/simple-admin-product-rpc/ent/category"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *product.CategoryUpdateReq) (*product.BaseResp, error) {
	if in.ParentId != consts.ParentIdNull {
		category, err := l.svcCtx.DB.Category.Query().Where(category.ParentIDEQ(in.ParentId)).First(l.ctx)
		if err != nil {
			return nil, errorhandler.DefaultEntError(l.Logger, err, in)
		}
		if category.ParentID != consts.ParentIdNull {
			return nil, errorx.NewInvalidArgumentError("CATEGORY_PARENT_NOT_FIRST_LEVEL")
		}
	}
	err := l.svcCtx.DB.Category.UpdateOneID(in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(&in.Status)).
		SetParentID(in.ParentId).SetName(in.Name).
		SetPicURL(in.PicUrl).
		SetNotNilBigPicURL(in.BigPicUrl).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
