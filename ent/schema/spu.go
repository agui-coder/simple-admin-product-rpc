package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	mixins2 "github.com/agui-coder/simple-admin-product-rpc/ent/schema/localmixin"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Spu struct {
	ent.Schema
}

func (Spu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("商品名称"),
		field.String("keyword").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("关键字"),
		field.String("introduction").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品简介"),
		field.Text("description").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品详情"),
		field.String("bar_code").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("条形码"),
		field.Uint64("category_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品分类编号"),
		field.Uint64("brand_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品品牌编号"),
		field.String("pic_url").
			Annotations(entsql.WithComments(true)).
			Comment("商品封面图"),
		field.JSON("slider_pic_urls", []string{}).Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品轮播图地址数组，以逗号分隔最多上传15张"),
		field.String("video_url").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品视频"),
		field.Uint8("unit").
			Annotations(entsql.WithComments(true)).
			Comment("单位"),
		field.Bool("spec_type").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("规格类型：0 单规格 1 多规格"),
		field.Int32("price").
			Annotations(entsql.WithComments(true)).
			Comment("商品价格，单位使用：分"),
		field.Int32("market_price").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("市场价，单位使用：分"),
		field.Int32("cost_price").
			Annotations(entsql.WithComments(true)).
			Comment("成本价，单位： 分"),
		field.Int32("stock").
			Annotations(entsql.WithComments(true)).
			Comment("库存"),
		field.Uint64("deliveryTemplate_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("物流配置模板编号"),
		field.Bool("recommend_hot").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("是否热卖推荐"),
		field.Bool("recommend_benefit").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("是否优惠推荐"),
		field.Bool("recommend_best").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("是否精品推荐"),
		field.Bool("recommend_new").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("是否新品推荐"),
		field.Bool("recommend_good").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("是否优品推荐"),
		field.Int32("give_integral").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("赠送积分"),
		field.JSON("give_coupon_template_ids", []uint64{}).Optional().
			Annotations(entsql.WithComments(true)).
			Comment("赠送的优惠劵编号的数组"),
		field.Bool("sub_commission_type").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("分销类型"),
		field.JSON("activity_orders", []int32{}).Optional().
			Annotations(entsql.WithComments(true)).
			Comment("活动展示顺序"),
		field.Int32("sales_count").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品销量"),
		field.Int32("virtual_sales_count").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("虚拟销量"),
		field.Int32("browse_count").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("商品点击量")}
}

func (Spu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Spu) Edges() []ent.Edge {
	return []ent.Edge{
		// SPU与SKU的一对多关系
		edge.To("skus", Sku.Type),
		// 产品品牌与SPU的一对多关系
		edge.From("brands", Brand.Type).Ref("spus").Unique().Field("brand_id"),
		// 产品分类与SPU的多对多关系
		edge.From("categorys", Category.Type).Ref("spus").Unique().Field("category_id"),
	}
}

func (Spu) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Spu) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "product_spu"}}
}
