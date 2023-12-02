package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	mixins2 "github.com/agui-coder/simple-admin-product-rpc/ent/schema/localmixin"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Brand struct {
	ent.Schema
}

func (Brand) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("品牌名称"),
		field.String("pic_url").
			Annotations(entsql.WithComments(true)).
			Comment("品牌图片"),
		field.String("description").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("品牌描述")}
}

func (Brand) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Brand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("spus", Spu.Type), //一个 Brand 可以拥有多个 Spu。
	}
}

func (Brand) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}

func (Brand) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "product_brand"}}
}
