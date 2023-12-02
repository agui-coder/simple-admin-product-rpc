package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/agui-coder/simple-admin-product-rpc/ent/schema/enttype"
	mixins2 "github.com/agui-coder/simple-admin-product-rpc/ent/schema/localmixin"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Sku struct {
	ent.Schema
}

func (Sku) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("spu_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("spu编号"),
		field.JSON("properties", []entType.SkuProperty{}).Optional().
			Annotations(entsql.WithComments(true)).
			Comment("属性数组，JSON 格式 [{propertId: , valueId: }, {propertId: , valueId: }]"),
		field.Int32("price").
			Annotations(entsql.WithComments(true)).
			Comment("商品价格，单位：分"),
		field.Int32("market_price").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("市场价，单位：分"),
		field.Int32("cost_price").
			Annotations(entsql.WithComments(true)).
			Comment("成本价，单位： 分"),
		field.String("bar_code").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("SKU 的条形码"),
		field.String("pic_url").
			Annotations(entsql.WithComments(true)).
			Comment("图片地址"),
		field.Int32("stock").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("库存"),
		field.Float("weight").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品重量，单位：kg 千克"),
		field.Float("volume").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品体积，单位：m^3 平米"),
		field.Int32("first_brokerage_price").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("一级分销的佣金，单位：分"),
		field.Float("second_brokerage_price").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("二级分销的佣金，单位：分"),
		field.Int32("sales_count").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品销量")}
}

func (Sku) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Sku) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spus", Spu.Type).Ref("skus").Unique().Field("spu_id"), // 多个Sku关联一个Spu
		edge.To("comments", Comment.Type),                                //一个SKU可以有多个评论
	}
}
func (Sku) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "product_sku"}}
}
