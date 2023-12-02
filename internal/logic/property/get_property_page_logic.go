package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"
	"github.com/agui-coder/simple-admin-product-rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"time"

	"github.com/agui-coder/simple-admin-product-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPropertyPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyPageLogic {
	return &GetPropertyPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPropertyPageLogic) GetPropertyPage(in *product.PropertyPageReq) (*product.PropertyListResp, error) {
	var predicate []predicate.Property
	if in.Name != nil {
		predicate = append(predicate, property.NameContains(*in.Name))
	}
	if in.Status != nil {
		predicate = append(predicate, property.StatusEQ(*pointy.GetStatusPointer(in.Status)))
	}
	if len(in.CreatedAt) > 0 {
		times := make([]time.Time, len(in.CreatedAt))
		for i, c := range in.CreatedAt {
			times[i] = time.UnixMilli(c)
		}
		predicate = append(predicate, property.CreatedAtIn(times...))
	}
	propertyPage, err := l.svcCtx.DB.Property.Query().Where(predicate...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	propertyList := make([]*product.PropertyInfo, len(propertyPage.List))
	for i, property := range propertyPage.List {
		status := uint32(property.Status)
		propertyList[i] = &product.PropertyInfo{
			Id:     &property.ID,
			Status: &status,
			Name:   &property.Name,
			Remark: &property.Remark,
		}
	}
	return &product.PropertyListResp{
		Total: propertyPage.PageDetails.Total,
		Data:  propertyList,
	}, nil
}
