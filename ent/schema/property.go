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

type Property struct {
	ent.Schema
}

func (Property) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("名称"),
		field.String("remark").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("备注")}
}

func (Property) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("property_values", PropertyValue.Type), // 一个 Property 可以有多个 PropertyValue
	}
}
func (Property) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "product_property"}}
}
func (Property) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name")}
}
