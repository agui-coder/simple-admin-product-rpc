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

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("user_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("评价人的用户编号，关联 MemberUserDO 的 id 编号"),
		field.String("user_nickname").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("评价人名称"),
		field.String("user_avatar").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("评价人头像"),
		field.Bool("anonymous").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("是否匿名"),
		field.Uint64("order_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("交易订单编号，关联 TradeOrderDO 的 id 编号"),
		field.Uint64("order_item_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("交易订单项编号，关联 TradeOrderItemDO 的 id 编号"),
		field.Uint64("spu_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品 SPU 编号，关联 ProductSpuDO 的 id"),
		field.String("spu_name").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品 SPU 名称"),
		field.Uint64("sku_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品 SKU 编号，关联 ProductSkuDO 的 id 编号"),
		field.String("sku_pic_url").
			Annotations(entsql.WithComments(true)).
			Comment("图片地址"),
		field.JSON("sku_properties", []entType.SkuProperty{}).Optional().
			Annotations(entsql.WithComments(true)).
			Comment("属性数组，JSON 格式 [{propertId: , valueId: }, {propertId: , valueId: }]"),
		field.Bool("visible").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("是否可见，true:显示false:隐藏"),
		field.Int8("scores").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("评分星级1-5分"),
		field.Int8("description_scores").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("描述星级 1-5 星"),
		field.Int8("benefit_scores").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("服务星级 1-5 星"),
		field.String("content").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("评论内容"),
		field.String("pic_urls").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("评论图片地址数组"),
		field.Bool("reply_status").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商家是否回复"),
		field.Int("reply_user_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("回复管理员编号，关联 AdminUserDO 的 id 编号"),
		field.String("reply_content").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商家回复内容"),
		field.Time("reply_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商家回复时间")}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("skus", Sku.Type).
			Ref("comments").
			Unique().Field("sku_id"), //一个SKU可以有多个评论
	}
}
func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "product_comment"}}
}
