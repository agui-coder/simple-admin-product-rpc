// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/agui-coder/simple-admin-product-rpc/ent/category"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Sort Number | 排序编号
	Sort uint32 `json:"sort,omitempty"`
	// Delete Time | 删除日期
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 父分类编号
	ParentID uint64 `json:"parent_id,omitempty"`
	// 分类名称
	Name string `json:"name,omitempty"`
	// 移动端分类图
	PicURL string `json:"pic_url,omitempty"`
	// PC 端分类图
	BigPicURL string `json:"big_pic_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CategoryQuery when eager-loading is set.
	Edges        CategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CategoryEdges holds the relations/edges for other nodes in the graph.
type CategoryEdges struct {
	// Spus holds the value of the spus edge.
	Spus []*Spu `json:"spus,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Category `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Category `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SpusOrErr returns the Spus value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) SpusOrErr() ([]*Spu, error) {
	if e.loadedTypes[0] {
		return e.Spus, nil
	}
	return nil, &NotLoadedError{edge: "spus"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CategoryEdges) ParentOrErr() (*Category, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) ChildrenOrErr() ([]*Category, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case category.FieldID, category.FieldStatus, category.FieldSort, category.FieldParentID:
			values[i] = new(sql.NullInt64)
		case category.FieldName, category.FieldPicURL, category.FieldBigPicURL:
			values[i] = new(sql.NullString)
		case category.FieldCreatedAt, category.FieldUpdatedAt, category.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case category.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case category.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case category.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = uint8(value.Int64)
			}
		case category.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				c.Sort = uint32(value.Int64)
			}
		case category.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case category.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				c.ParentID = uint64(value.Int64)
			}
		case category.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case category.FieldPicURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic_url", values[i])
			} else if value.Valid {
				c.PicURL = value.String
			}
		case category.FieldBigPicURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field big_pic_url", values[i])
			} else if value.Valid {
				c.BigPicURL = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Category.
// This includes values selected through modifiers, order, etc.
func (c *Category) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QuerySpus queries the "spus" edge of the Category entity.
func (c *Category) QuerySpus() *SpuQuery {
	return NewCategoryClient(c.config).QuerySpus(c)
}

// QueryParent queries the "parent" edge of the Category entity.
func (c *Category) QueryParent() *CategoryQuery {
	return NewCategoryClient(c.config).QueryParent(c)
}

// QueryChildren queries the "children" edge of the Category entity.
func (c *Category) QueryChildren() *CategoryQuery {
	return NewCategoryClient(c.config).QueryChildren(c)
}

// Update returns a builder for updating this Category.
// Note that you need to call Category.Unwrap() before calling this method if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return NewCategoryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", c.Sort))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ParentID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("pic_url=")
	builder.WriteString(c.PicURL)
	builder.WriteString(", ")
	builder.WriteString("big_pic_url=")
	builder.WriteString(c.BigPicURL)
	builder.WriteByte(')')
	return builder.String()
}

// Categories is a parsable slice of Category.
type Categories []*Category
