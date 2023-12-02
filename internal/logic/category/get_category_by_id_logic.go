package category

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(in *product.IDReq) (*product.CategoryInfo, error) {
	c, err := l.svcCtx.DB.Category.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	status := uint32(c.Status)
	createdAt := c.CreatedAt.UnixMilli()
	updatedAt := c.UpdatedAt.UnixMilli()
	return &product.CategoryInfo{
		Id:        &c.ID,
		Status:    &status,
		ParentId:  &c.ParentID,
		Name:      &c.Name,
		PicUrl:    &c.PicURL,
		BigPicUrl: &c.BigPicURL,
		Sort:      &c.Sort,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}, nil
}
