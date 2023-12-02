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

type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("parent_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("父分类编号"),
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("分类名称"),
		field.String("pic_url").
			Annotations(entsql.WithComments(true)).
			Comment("移动端分类图"),
		field.String("big_pic_url").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("PC 端分类图")}
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("spus", Spu.Type), //一个 Category 可以拥有多个 Spu。
		edge.To("children", Category.Type).From("parent").Unique().Field("parent_id"),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "product_category"}}
}
