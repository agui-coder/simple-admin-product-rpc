// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package productclient

import (
	"context"

	"github.com/agui-coder/simple-admin-product-rpc/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp              = product.BaseIDResp
	BaseMsg                 = product.BaseMsg
	BaseResp                = product.BaseResp
	BaseUUIDResp            = product.BaseUUIDResp
	BrandCreateReq          = product.BrandCreateReq
	BrandInfo               = product.BrandInfo
	BrandListReq            = product.BrandListReq
	BrandListResp           = product.BrandListResp
	BrandPageReq            = product.BrandPageReq
	BrandResp               = product.BrandResp
	BrandUpdateReq          = product.BrandUpdateReq
	CategoryBase            = product.CategoryBase
	CategoryCreateReq       = product.CategoryCreateReq
	CategoryInfo            = product.CategoryInfo
	CategoryLevelResp       = product.CategoryLevelResp
	CategoryListReq         = product.CategoryListReq
	CategoryListResp        = product.CategoryListResp
	CategoryUpdateReq       = product.CategoryUpdateReq
	CommentInfo             = product.CommentInfo
	CommentListReq          = product.CommentListReq
	CommentListResp         = product.CommentListResp
	Empty                   = product.Empty
	GiveCouponTemplate      = product.GiveCouponTemplate
	IDReq                   = product.IDReq
	IDsReq                  = product.IDsReq
	Ids                     = product.Ids
	PageInfoReq             = product.PageInfoReq
	PropertyCreateReq       = product.PropertyCreateReq
	PropertyInfo            = product.PropertyInfo
	PropertyListByIdsReq    = product.PropertyListByIdsReq
	PropertyListByNameReq   = product.PropertyListByNameReq
	PropertyListResp        = product.PropertyListResp
	PropertyPageReq         = product.PropertyPageReq
	PropertyUpdateReq       = product.PropertyUpdateReq
	PropertyValueCreateReq  = product.PropertyValueCreateReq
	PropertyValueDetailResp = product.PropertyValueDetailResp
	PropertyValueInfo       = product.PropertyValueInfo
	PropertyValueListReq    = product.PropertyValueListReq
	PropertyValueListResp   = product.PropertyValueListResp
	PropertyValuePageReq    = product.PropertyValuePageReq
	PropertyValueUpdateReq  = product.PropertyValueUpdateReq
	SkuCreateOrUpdateReq    = product.SkuCreateOrUpdateReq
	SkuInfo                 = product.SkuInfo
	SkuListResp             = product.SkuListResp
	SkuProperty             = product.SkuProperty
	SpuCreateReq            = product.SpuCreateReq
	SpuInfo                 = product.SpuInfo
	SpuListReq              = product.SpuListReq
	SpuListResp             = product.SpuListResp
	SpuUpdateReqVO          = product.SpuUpdateReqVO
	SpuUpdateStatusReq      = product.SpuUpdateStatusReq
	UUIDReq                 = product.UUIDReq
	UUIDsReq                = product.UUIDsReq

	Product interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Brand management
		CreateBrand(ctx context.Context, in *BrandCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateBrand(ctx context.Context, in *BrandUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetBrandList(ctx context.Context, in *BrandListReq, opts ...grpc.CallOption) (*BrandListResp, error)
		GetBrandById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BrandInfo, error)
		DeleteBrand(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetBrandPage(ctx context.Context, in *BrandPageReq, opts ...grpc.CallOption) (*BrandListResp, error)
		// Category management
		CreateCategory(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateCategory(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteCategory(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetCategoryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CategoryInfo, error)
		GetCategoryLevel(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CategoryLevelResp, error)
		GetEnableCategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
		// Comment management
		CreateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
		GetCommentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CommentInfo, error)
		DeleteComment(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Property management
		CreateProperty(ctx context.Context, in *PropertyCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateProperty(ctx context.Context, in *PropertyUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteProperty(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetPropertyListByName(ctx context.Context, in *PropertyListByNameReq, opts ...grpc.CallOption) (*PropertyListResp, error)
		GetPropertyPage(ctx context.Context, in *PropertyPageReq, opts ...grpc.CallOption) (*PropertyListResp, error)
		GetPropertyById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PropertyInfo, error)
		GetPropertyListByIds(ctx context.Context, in *PropertyListByIdsReq, opts ...grpc.CallOption) (*PropertyListResp, error)
		// PropertyValue management
		CreatePropertyValue(ctx context.Context, in *PropertyValueCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		DeletePropertyValue(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetPropertyValueById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PropertyValueInfo, error)
		GetPropertyValueDetailList(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*PropertyValueDetailResp, error)
		GetPropertyValueListByPropertyId(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*PropertyValueListResp, error)
		GetPropertyValuePage(ctx context.Context, in *PropertyValuePageReq, opts ...grpc.CallOption) (*PropertyValueListResp, error)
		UpdatePropertyValue(ctx context.Context, in *PropertyValueUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Sku management
		DeleteSku(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetSkuList(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*SkuListResp, error)
		GetSkuById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SkuInfo, error)
		GetSkuListBySpuId(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SkuListResp, error)
		// Spu management
		CreateSpu(ctx context.Context, in *SpuCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateSpu(ctx context.Context, in *SpuUpdateReqVO, opts ...grpc.CallOption) (*BaseResp, error)
		UpdateStatus(ctx context.Context, in *SpuUpdateStatusReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetSpuList(ctx context.Context, in *SpuListReq, opts ...grpc.CallOption) (*SpuListResp, error)
		GetSpuById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SpuInfo, error)
		DeleteSpu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		ValidateSpuList(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*SpuListResp, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Brand management
func (m *defaultProduct) CreateBrand(ctx context.Context, in *BrandCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateBrand(ctx, in, opts...)
}

func (m *defaultProduct) UpdateBrand(ctx context.Context, in *BrandUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateBrand(ctx, in, opts...)
}

func (m *defaultProduct) GetBrandList(ctx context.Context, in *BrandListReq, opts ...grpc.CallOption) (*BrandListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetBrandList(ctx, in, opts...)
}

func (m *defaultProduct) GetBrandById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BrandInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetBrandById(ctx, in, opts...)
}

func (m *defaultProduct) DeleteBrand(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeleteBrand(ctx, in, opts...)
}

func (m *defaultProduct) GetBrandPage(ctx context.Context, in *BrandPageReq, opts ...grpc.CallOption) (*BrandListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetBrandPage(ctx, in, opts...)
}

// Category management
func (m *defaultProduct) CreateCategory(ctx context.Context, in *CategoryCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateCategory(ctx, in, opts...)
}

func (m *defaultProduct) UpdateCategory(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateCategory(ctx, in, opts...)
}

func (m *defaultProduct) DeleteCategory(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeleteCategory(ctx, in, opts...)
}

func (m *defaultProduct) GetCategoryById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CategoryInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetCategoryById(ctx, in, opts...)
}

func (m *defaultProduct) GetCategoryLevel(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CategoryLevelResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetCategoryLevel(ctx, in, opts...)
}

func (m *defaultProduct) GetEnableCategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetEnableCategoryList(ctx, in, opts...)
}

// Comment management
func (m *defaultProduct) CreateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateComment(ctx, in, opts...)
}

func (m *defaultProduct) UpdateComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateComment(ctx, in, opts...)
}

func (m *defaultProduct) GetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}

func (m *defaultProduct) GetCommentById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*CommentInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetCommentById(ctx, in, opts...)
}

func (m *defaultProduct) DeleteComment(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

// Property management
func (m *defaultProduct) CreateProperty(ctx context.Context, in *PropertyCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateProperty(ctx, in, opts...)
}

func (m *defaultProduct) UpdateProperty(ctx context.Context, in *PropertyUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateProperty(ctx, in, opts...)
}

func (m *defaultProduct) DeleteProperty(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeleteProperty(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyListByName(ctx context.Context, in *PropertyListByNameReq, opts ...grpc.CallOption) (*PropertyListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyListByName(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyPage(ctx context.Context, in *PropertyPageReq, opts ...grpc.CallOption) (*PropertyListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyPage(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PropertyInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyById(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyListByIds(ctx context.Context, in *PropertyListByIdsReq, opts ...grpc.CallOption) (*PropertyListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyListByIds(ctx, in, opts...)
}

// PropertyValue management
func (m *defaultProduct) CreatePropertyValue(ctx context.Context, in *PropertyValueCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreatePropertyValue(ctx, in, opts...)
}

func (m *defaultProduct) DeletePropertyValue(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeletePropertyValue(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyValueById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*PropertyValueInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyValueById(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyValueDetailList(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*PropertyValueDetailResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyValueDetailList(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyValueListByPropertyId(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*PropertyValueListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyValueListByPropertyId(ctx, in, opts...)
}

func (m *defaultProduct) GetPropertyValuePage(ctx context.Context, in *PropertyValuePageReq, opts ...grpc.CallOption) (*PropertyValueListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetPropertyValuePage(ctx, in, opts...)
}

func (m *defaultProduct) UpdatePropertyValue(ctx context.Context, in *PropertyValueUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdatePropertyValue(ctx, in, opts...)
}

// Sku management
func (m *defaultProduct) DeleteSku(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeleteSku(ctx, in, opts...)
}

func (m *defaultProduct) GetSkuList(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*SkuListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetSkuList(ctx, in, opts...)
}

func (m *defaultProduct) GetSkuById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SkuInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetSkuById(ctx, in, opts...)
}

func (m *defaultProduct) GetSkuListBySpuId(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SkuListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetSkuListBySpuId(ctx, in, opts...)
}

// Spu management
func (m *defaultProduct) CreateSpu(ctx context.Context, in *SpuCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateSpu(ctx, in, opts...)
}

func (m *defaultProduct) UpdateSpu(ctx context.Context, in *SpuUpdateReqVO, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateSpu(ctx, in, opts...)
}

func (m *defaultProduct) UpdateStatus(ctx context.Context, in *SpuUpdateStatusReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateStatus(ctx, in, opts...)
}

func (m *defaultProduct) GetSpuList(ctx context.Context, in *SpuListReq, opts ...grpc.CallOption) (*SpuListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetSpuList(ctx, in, opts...)
}

func (m *defaultProduct) GetSpuById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SpuInfo, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetSpuById(ctx, in, opts...)
}

func (m *defaultProduct) DeleteSpu(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeleteSpu(ctx, in, opts...)
}

func (m *defaultProduct) ValidateSpuList(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*SpuListResp, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.ValidateSpuList(ctx, in, opts...)
}
