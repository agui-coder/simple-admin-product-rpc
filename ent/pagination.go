// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github.com/agui-coder/simple-admin-product-rpc/ent/brand"
	"github.com/agui-coder/simple-admin-product-rpc/ent/category"
	"github.com/agui-coder/simple-admin-product-rpc/ent/comment"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"github.com/agui-coder/simple-admin-product-rpc/ent/sku"
	"github.com/agui-coder/simple-admin-product-rpc/ent/spu"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type BrandPager struct {
	Order  brand.OrderOption
	Filter func(*BrandQuery) (*BrandQuery, error)
}

// BrandPaginateOption enables pagination customization.
type BrandPaginateOption func(*BrandPager)

// DefaultBrandOrder is the default ordering of Brand.
var DefaultBrandOrder = Desc(brand.FieldID)

func newBrandPager(opts []BrandPaginateOption) (*BrandPager, error) {
	pager := &BrandPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultBrandOrder
	}
	return pager, nil
}

func (p *BrandPager) ApplyFilter(query *BrandQuery) (*BrandQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// BrandPageList is Brand PageList result.
type BrandPageList struct {
	List        []*Brand     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (b *BrandQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...BrandPaginateOption,
) (*BrandPageList, error) {

	pager, err := newBrandPager(opts)
	if err != nil {
		return nil, err
	}

	if b, err = pager.ApplyFilter(b); err != nil {
		return nil, err
	}

	ret := &BrandPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := b.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		b = b.Order(pager.Order)
	} else {
		b = b.Order(DefaultBrandOrder)
	}

	b = b.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := b.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type CategoryPager struct {
	Order  category.OrderOption
	Filter func(*CategoryQuery) (*CategoryQuery, error)
}

// CategoryPaginateOption enables pagination customization.
type CategoryPaginateOption func(*CategoryPager)

// DefaultCategoryOrder is the default ordering of Category.
var DefaultCategoryOrder = Desc(category.FieldID)

func newCategoryPager(opts []CategoryPaginateOption) (*CategoryPager, error) {
	pager := &CategoryPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultCategoryOrder
	}
	return pager, nil
}

func (p *CategoryPager) ApplyFilter(query *CategoryQuery) (*CategoryQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// CategoryPageList is Category PageList result.
type CategoryPageList struct {
	List        []*Category  `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (c *CategoryQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...CategoryPaginateOption,
) (*CategoryPageList, error) {

	pager, err := newCategoryPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.ApplyFilter(c); err != nil {
		return nil, err
	}

	ret := &CategoryPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := c.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		c = c.Order(pager.Order)
	} else {
		c = c.Order(DefaultCategoryOrder)
	}

	c = c.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type CommentPager struct {
	Order  comment.OrderOption
	Filter func(*CommentQuery) (*CommentQuery, error)
}

// CommentPaginateOption enables pagination customization.
type CommentPaginateOption func(*CommentPager)

// DefaultCommentOrder is the default ordering of Comment.
var DefaultCommentOrder = Desc(comment.FieldID)

func newCommentPager(opts []CommentPaginateOption) (*CommentPager, error) {
	pager := &CommentPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultCommentOrder
	}
	return pager, nil
}

func (p *CommentPager) ApplyFilter(query *CommentQuery) (*CommentQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// CommentPageList is Comment PageList result.
type CommentPageList struct {
	List        []*Comment   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (c *CommentQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...CommentPaginateOption,
) (*CommentPageList, error) {

	pager, err := newCommentPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.ApplyFilter(c); err != nil {
		return nil, err
	}

	ret := &CommentPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := c.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		c = c.Order(pager.Order)
	} else {
		c = c.Order(DefaultCommentOrder)
	}

	c = c.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type PropertyPager struct {
	Order  property.OrderOption
	Filter func(*PropertyQuery) (*PropertyQuery, error)
}

// PropertyPaginateOption enables pagination customization.
type PropertyPaginateOption func(*PropertyPager)

// DefaultPropertyOrder is the default ordering of Property.
var DefaultPropertyOrder = Desc(property.FieldID)

func newPropertyPager(opts []PropertyPaginateOption) (*PropertyPager, error) {
	pager := &PropertyPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultPropertyOrder
	}
	return pager, nil
}

func (p *PropertyPager) ApplyFilter(query *PropertyQuery) (*PropertyQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// PropertyPageList is Property PageList result.
type PropertyPageList struct {
	List        []*Property  `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (pr *PropertyQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...PropertyPaginateOption,
) (*PropertyPageList, error) {

	pager, err := newPropertyPager(opts)
	if err != nil {
		return nil, err
	}

	if pr, err = pager.ApplyFilter(pr); err != nil {
		return nil, err
	}

	ret := &PropertyPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := pr.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		pr = pr.Order(pager.Order)
	} else {
		pr = pr.Order(DefaultPropertyOrder)
	}

	pr = pr.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := pr.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type PropertyValuePager struct {
	Order  propertyvalue.OrderOption
	Filter func(*PropertyValueQuery) (*PropertyValueQuery, error)
}

// PropertyValuePaginateOption enables pagination customization.
type PropertyValuePaginateOption func(*PropertyValuePager)

// DefaultPropertyValueOrder is the default ordering of PropertyValue.
var DefaultPropertyValueOrder = Desc(propertyvalue.FieldID)

func newPropertyValuePager(opts []PropertyValuePaginateOption) (*PropertyValuePager, error) {
	pager := &PropertyValuePager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultPropertyValueOrder
	}
	return pager, nil
}

func (p *PropertyValuePager) ApplyFilter(query *PropertyValueQuery) (*PropertyValueQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// PropertyValuePageList is PropertyValue PageList result.
type PropertyValuePageList struct {
	List        []*PropertyValue `json:"list"`
	PageDetails *PageDetails     `json:"pageDetails"`
}

func (pv *PropertyValueQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...PropertyValuePaginateOption,
) (*PropertyValuePageList, error) {

	pager, err := newPropertyValuePager(opts)
	if err != nil {
		return nil, err
	}

	if pv, err = pager.ApplyFilter(pv); err != nil {
		return nil, err
	}

	ret := &PropertyValuePageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := pv.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		pv = pv.Order(pager.Order)
	} else {
		pv = pv.Order(DefaultPropertyValueOrder)
	}

	pv = pv.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := pv.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type SkuPager struct {
	Order  sku.OrderOption
	Filter func(*SkuQuery) (*SkuQuery, error)
}

// SkuPaginateOption enables pagination customization.
type SkuPaginateOption func(*SkuPager)

// DefaultSkuOrder is the default ordering of Sku.
var DefaultSkuOrder = Desc(sku.FieldID)

func newSkuPager(opts []SkuPaginateOption) (*SkuPager, error) {
	pager := &SkuPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultSkuOrder
	}
	return pager, nil
}

func (p *SkuPager) ApplyFilter(query *SkuQuery) (*SkuQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// SkuPageList is Sku PageList result.
type SkuPageList struct {
	List        []*Sku       `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *SkuQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...SkuPaginateOption,
) (*SkuPageList, error) {

	pager, err := newSkuPager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &SkuPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := s.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultSkuOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type SpuPager struct {
	Order  spu.OrderOption
	Filter func(*SpuQuery) (*SpuQuery, error)
}

// SpuPaginateOption enables pagination customization.
type SpuPaginateOption func(*SpuPager)

// DefaultSpuOrder is the default ordering of Spu.
var DefaultSpuOrder = Desc(spu.FieldID)

func newSpuPager(opts []SpuPaginateOption) (*SpuPager, error) {
	pager := &SpuPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultSpuOrder
	}
	return pager, nil
}

func (p *SpuPager) ApplyFilter(query *SpuQuery) (*SpuQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// SpuPageList is Spu PageList result.
type SpuPageList struct {
	List        []*Spu       `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *SpuQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...SpuPaginateOption,
) (*SpuPageList, error) {

	pager, err := newSpuPager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &SpuPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := s.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultSpuOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
