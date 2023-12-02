// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	entType "github.com/agui-coder/simple-admin-product-rpc/ent/schema/enttype"
	"github.com/agui-coder/simple-admin-product-rpc/ent/sku"
	"github.com/agui-coder/simple-admin-product-rpc/ent/spu"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Sku is the model entity for the Sku schema.
type Sku struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Delete Time | 删除日期
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// spu编号
	SpuID uint64 `json:"spu_id,omitempty"`
	// 属性数组，JSON 格式 [{propertId: , valueId: }, {propertId: , valueId: }]
	Properties []entType.SkuProperty `json:"properties,omitempty"`
	// 商品价格，单位：分
	Price int32 `json:"price,omitempty"`
	// 市场价，单位：分
	MarketPrice int32 `json:"market_price,omitempty"`
	// 成本价，单位： 分
	CostPrice int32 `json:"cost_price,omitempty"`
	// SKU 的条形码
	BarCode string `json:"bar_code,omitempty"`
	// 图片地址
	PicURL string `json:"pic_url,omitempty"`
	// 库存
	Stock int32 `json:"stock,omitempty"`
	// 商品重量，单位：kg 千克
	Weight float64 `json:"weight,omitempty"`
	// 商品体积，单位：m^3 平米
	Volume float64 `json:"volume,omitempty"`
	// 一级分销的佣金，单位：分
	FirstBrokeragePrice int32 `json:"first_brokerage_price,omitempty"`
	// 二级分销的佣金，单位：分
	SecondBrokeragePrice float64 `json:"second_brokerage_price,omitempty"`
	// 商品销量
	SalesCount int32 `json:"sales_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkuQuery when eager-loading is set.
	Edges        SkuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SkuEdges holds the relations/edges for other nodes in the graph.
type SkuEdges struct {
	// Spus holds the value of the spus edge.
	Spus *Spu `json:"spus,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SpusOrErr returns the Spus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SkuEdges) SpusOrErr() (*Spu, error) {
	if e.loadedTypes[0] {
		if e.Spus == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: spu.Label}
		}
		return e.Spus, nil
	}
	return nil, &NotLoadedError{edge: "spus"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e SkuEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sku) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sku.FieldProperties:
			values[i] = new([]byte)
		case sku.FieldWeight, sku.FieldVolume, sku.FieldSecondBrokeragePrice:
			values[i] = new(sql.NullFloat64)
		case sku.FieldID, sku.FieldSpuID, sku.FieldPrice, sku.FieldMarketPrice, sku.FieldCostPrice, sku.FieldStock, sku.FieldFirstBrokeragePrice, sku.FieldSalesCount:
			values[i] = new(sql.NullInt64)
		case sku.FieldBarCode, sku.FieldPicURL:
			values[i] = new(sql.NullString)
		case sku.FieldCreatedAt, sku.FieldUpdatedAt, sku.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sku fields.
func (s *Sku) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sku.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case sku.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case sku.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case sku.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = value.Time
			}
		case sku.FieldSpuID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field spu_id", values[i])
			} else if value.Valid {
				s.SpuID = uint64(value.Int64)
			}
		case sku.FieldProperties:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field properties", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Properties); err != nil {
					return fmt.Errorf("unmarshal field properties: %w", err)
				}
			}
		case sku.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				s.Price = int32(value.Int64)
			}
		case sku.FieldMarketPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field market_price", values[i])
			} else if value.Valid {
				s.MarketPrice = int32(value.Int64)
			}
		case sku.FieldCostPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cost_price", values[i])
			} else if value.Valid {
				s.CostPrice = int32(value.Int64)
			}
		case sku.FieldBarCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bar_code", values[i])
			} else if value.Valid {
				s.BarCode = value.String
			}
		case sku.FieldPicURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic_url", values[i])
			} else if value.Valid {
				s.PicURL = value.String
			}
		case sku.FieldStock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stock", values[i])
			} else if value.Valid {
				s.Stock = int32(value.Int64)
			}
		case sku.FieldWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				s.Weight = value.Float64
			}
		case sku.FieldVolume:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field volume", values[i])
			} else if value.Valid {
				s.Volume = value.Float64
			}
		case sku.FieldFirstBrokeragePrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field first_brokerage_price", values[i])
			} else if value.Valid {
				s.FirstBrokeragePrice = int32(value.Int64)
			}
		case sku.FieldSecondBrokeragePrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field second_brokerage_price", values[i])
			} else if value.Valid {
				s.SecondBrokeragePrice = value.Float64
			}
		case sku.FieldSalesCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sales_count", values[i])
			} else if value.Valid {
				s.SalesCount = int32(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Sku.
// This includes values selected through modifiers, order, etc.
func (s *Sku) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QuerySpus queries the "spus" edge of the Sku entity.
func (s *Sku) QuerySpus() *SpuQuery {
	return NewSkuClient(s.config).QuerySpus(s)
}

// QueryComments queries the "comments" edge of the Sku entity.
func (s *Sku) QueryComments() *CommentQuery {
	return NewSkuClient(s.config).QueryComments(s)
}

// Update returns a builder for updating this Sku.
// Note that you need to call Sku.Unwrap() before calling this method if this Sku
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sku) Update() *SkuUpdateOne {
	return NewSkuClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Sku entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sku) Unwrap() *Sku {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sku is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sku) String() string {
	var builder strings.Builder
	builder.WriteString("Sku(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(s.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("spu_id=")
	builder.WriteString(fmt.Sprintf("%v", s.SpuID))
	builder.WriteString(", ")
	builder.WriteString("properties=")
	builder.WriteString(fmt.Sprintf("%v", s.Properties))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", s.Price))
	builder.WriteString(", ")
	builder.WriteString("market_price=")
	builder.WriteString(fmt.Sprintf("%v", s.MarketPrice))
	builder.WriteString(", ")
	builder.WriteString("cost_price=")
	builder.WriteString(fmt.Sprintf("%v", s.CostPrice))
	builder.WriteString(", ")
	builder.WriteString("bar_code=")
	builder.WriteString(s.BarCode)
	builder.WriteString(", ")
	builder.WriteString("pic_url=")
	builder.WriteString(s.PicURL)
	builder.WriteString(", ")
	builder.WriteString("stock=")
	builder.WriteString(fmt.Sprintf("%v", s.Stock))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", s.Weight))
	builder.WriteString(", ")
	builder.WriteString("volume=")
	builder.WriteString(fmt.Sprintf("%v", s.Volume))
	builder.WriteString(", ")
	builder.WriteString("first_brokerage_price=")
	builder.WriteString(fmt.Sprintf("%v", s.FirstBrokeragePrice))
	builder.WriteString(", ")
	builder.WriteString("second_brokerage_price=")
	builder.WriteString(fmt.Sprintf("%v", s.SecondBrokeragePrice))
	builder.WriteString(", ")
	builder.WriteString("sales_count=")
	builder.WriteString(fmt.Sprintf("%v", s.SalesCount))
	builder.WriteByte(')')
	return builder.String()
}

// Skus is a parsable slice of Sku.
type Skus []*Sku
