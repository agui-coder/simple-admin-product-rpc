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

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Category management
func (l *CreateCategoryLogic) CreateCategory(in *product.CategoryCreateReq) (*product.BaseIDResp, error) {
	if in.ParentId != consts.ParentIdNull {
		category, err := l.svcCtx.DB.Category.Query().Where(category.ParentIDEQ(in.ParentId)).First(l.ctx)
		if err != nil {
			return nil, errorhandler.DefaultEntError(l.Logger, err, in)
		}
		if category.ParentID != consts.ParentIdNull {
			return nil, errorx.NewInvalidArgumentError("CATEGORY_PARENT_NOT_FIRST_LEVEL")
		}
	}
	category, err := l.svcCtx.DB.Category.Create().
		SetParentID(in.ParentId).
		SetName(in.Name).
		SetPicURL(in.PicUrl).
		SetNotNilSort(in.Sort).
		SetNotNilStatus(pointy.GetStatusPointer(&in.Status)).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &product.BaseIDResp{
		Id:  category.ID,
		Msg: i18n.CreateSuccess,
	}, nil
}
