package category

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"math"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLevelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLevelLogic {
	return &GetCategoryLevelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryLevelLogic) GetCategoryLevel(in *product.IDReq) (*product.CategoryLevelResp, error) {
	id := in.Id
	level := int64(0)
	if in.Id == 0 {
		return &product.CategoryLevelResp{Level: &level}, nil
	}
	level = 1
	for i := 0; i < math.MaxInt8; i++ {
		category, err := l.svcCtx.DB.Category.Get(l.ctx, id)
		if err != nil && ent.IsNotFound(err) {
			return nil, errorhandler.DefaultEntError(l.Logger, err, id)
		}
		if !ent.IsNotFound(err) && category.ParentID == 0 {
			break
		}
		level++
		id = category.ParentID
	}
	return &product.CategoryLevelResp{Level: &level}, nil
}
