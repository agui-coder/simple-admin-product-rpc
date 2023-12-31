// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PropertyValue is the model entity for the PropertyValue schema.
type PropertyValue struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Delete Time | 删除日期
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 属性项的编号
	PropertyID uint64 `json:"property_id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PropertyValueQuery when eager-loading is set.
	Edges        PropertyValueEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PropertyValueEdges holds the relations/edges for other nodes in the graph.
type PropertyValueEdges struct {
	// Propertys holds the value of the propertys edge.
	Propertys *Property `json:"propertys,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PropertysOrErr returns the Propertys value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PropertyValueEdges) PropertysOrErr() (*Property, error) {
	if e.loadedTypes[0] {
		if e.Propertys == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: property.Label}
		}
		return e.Propertys, nil
	}
	return nil, &NotLoadedError{edge: "propertys"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PropertyValue) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case propertyvalue.FieldID, propertyvalue.FieldPropertyID:
			values[i] = new(sql.NullInt64)
		case propertyvalue.FieldName, propertyvalue.FieldRemark:
			values[i] = new(sql.NullString)
		case propertyvalue.FieldCreatedAt, propertyvalue.FieldUpdatedAt, propertyvalue.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PropertyValue fields.
func (pv *PropertyValue) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case propertyvalue.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pv.ID = uint64(value.Int64)
		case propertyvalue.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pv.CreatedAt = value.Time
			}
		case propertyvalue.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pv.UpdatedAt = value.Time
			}
		case propertyvalue.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pv.DeletedAt = value.Time
			}
		case propertyvalue.FieldPropertyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field property_id", values[i])
			} else if value.Valid {
				pv.PropertyID = uint64(value.Int64)
			}
		case propertyvalue.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pv.Name = value.String
			}
		case propertyvalue.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				pv.Remark = value.String
			}
		default:
			pv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PropertyValue.
// This includes values selected through modifiers, order, etc.
func (pv *PropertyValue) Value(name string) (ent.Value, error) {
	return pv.selectValues.Get(name)
}

// QueryPropertys queries the "propertys" edge of the PropertyValue entity.
func (pv *PropertyValue) QueryPropertys() *PropertyQuery {
	return NewPropertyValueClient(pv.config).QueryPropertys(pv)
}

// Update returns a builder for updating this PropertyValue.
// Note that you need to call PropertyValue.Unwrap() before calling this method if this PropertyValue
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *PropertyValue) Update() *PropertyValueUpdateOne {
	return NewPropertyValueClient(pv.config).UpdateOne(pv)
}

// Unwrap unwraps the PropertyValue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pv *PropertyValue) Unwrap() *PropertyValue {
	_tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("ent: PropertyValue is not a transactional entity")
	}
	pv.config.driver = _tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *PropertyValue) String() string {
	var builder strings.Builder
	builder.WriteString("PropertyValue(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pv.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pv.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pv.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pv.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("property_id=")
	builder.WriteString(fmt.Sprintf("%v", pv.PropertyID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pv.Name)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(pv.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// PropertyValues is a parsable slice of PropertyValue.
type PropertyValues []*PropertyValue
