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

type PropertyValue struct {
	ent.Schema
}

func (PropertyValue) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("property_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("属性项的编号"),
		field.String("name").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("名称"),
		field.String("remark").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("备注"),
	}
}

func (PropertyValue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (PropertyValue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("propertys", Property.Type).Ref("property_values").Unique().Field("property_id"), //多个属性值属于同一个属性
	}
}
func (PropertyValue) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "product_property_value"}}
}
