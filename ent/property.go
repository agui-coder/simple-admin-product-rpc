// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Property is the model entity for the Property schema.
type Property struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Delete Time | 删除日期
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PropertyQuery when eager-loading is set.
	Edges        PropertyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PropertyEdges holds the relations/edges for other nodes in the graph.
type PropertyEdges struct {
	// PropertyValues holds the value of the property_values edge.
	PropertyValues []*PropertyValue `json:"property_values,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PropertyValuesOrErr returns the PropertyValues value or an error if the edge
// was not loaded in eager-loading.
func (e PropertyEdges) PropertyValuesOrErr() ([]*PropertyValue, error) {
	if e.loadedTypes[0] {
		return e.PropertyValues, nil
	}
	return nil, &NotLoadedError{edge: "property_values"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Property) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case property.FieldID, property.FieldStatus:
			values[i] = new(sql.NullInt64)
		case property.FieldName, property.FieldRemark:
			values[i] = new(sql.NullString)
		case property.FieldCreatedAt, property.FieldUpdatedAt, property.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Property fields.
func (pr *Property) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case property.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint64(value.Int64)
		case property.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case property.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case property.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = uint8(value.Int64)
			}
		case property.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = value.Time
			}
		case property.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case property.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				pr.Remark = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Property.
// This includes values selected through modifiers, order, etc.
func (pr *Property) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryPropertyValues queries the "property_values" edge of the Property entity.
func (pr *Property) QueryPropertyValues() *PropertyValueQuery {
	return NewPropertyClient(pr.config).QueryPropertyValues(pr)
}

// Update returns a builder for updating this Property.
// Note that you need to call Property.Unwrap() before calling this method if this Property
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Property) Update() *PropertyUpdateOne {
	return NewPropertyClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Property entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Property) Unwrap() *Property {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Property is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Property) String() string {
	var builder strings.Builder
	builder.WriteString("Property(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(pr.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Properties is a parsable slice of Property.
type Properties []*Property
