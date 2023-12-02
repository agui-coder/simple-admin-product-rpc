package category

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/category"
	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEnableCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEnableCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnableCategoryListLogic {
	return &GetEnableCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEnableCategoryListLogic) GetEnableCategoryList(in *product.CategoryListReq) (*product.CategoryListResp, error) {
	var predicate []predicate.Category
	if in.Name != nil {
		predicate = append(predicate, category.NameContains(*in.Name))
	}
	if in.Status != nil {
		predicate = append(predicate, category.StatusEQ(*pointy.GetStatusPointer(in.Status)))
	}
	if in.ParentId != nil {
		predicate = append(predicate, category.ParentIDEQ(*in.ParentId))
	}
	categories, err := l.svcCtx.DB.Category.Query().Where(predicate...).All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	categoryInfos := make([]*product.CategoryInfo, len(categories))
	for i, c := range categories {
		status := uint32(c.Status)
		createdAt := c.CreatedAt.UnixMilli()
		updatedAt := c.UpdatedAt.UnixMilli()
		categoryInfos[i] = &product.CategoryInfo{
			Id:        &c.ID,
			Status:    &status,
			ParentId:  &c.ParentID,
			Name:      &c.Name,
			PicUrl:    &c.PicURL,
			BigPicUrl: &c.BigPicURL,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		}
	}
	return &product.CategoryListResp{Data: categoryInfos}, nil
}
