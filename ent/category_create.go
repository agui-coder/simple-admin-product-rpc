// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/agui-coder/simple-admin-product-rpc/ent/category"
	"github.com/agui-coder/simple-admin-product-rpc/ent/spu"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CategoryCreate) SetCreatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableCreatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CategoryCreate) SetUpdatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableUpdatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CategoryCreate) SetStatus(u uint8) *CategoryCreate {
	cc.mutation.SetStatus(u)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableStatus(u *uint8) *CategoryCreate {
	if u != nil {
		cc.SetStatus(*u)
	}
	return cc
}

// SetSort sets the "sort" field.
func (cc *CategoryCreate) SetSort(u uint32) *CategoryCreate {
	cc.mutation.SetSort(u)
	return cc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableSort(u *uint32) *CategoryCreate {
	if u != nil {
		cc.SetSort(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CategoryCreate) SetDeletedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableDeletedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetParentID sets the "parent_id" field.
func (cc *CategoryCreate) SetParentID(u uint64) *CategoryCreate {
	cc.mutation.SetParentID(u)
	return cc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableParentID(u *uint64) *CategoryCreate {
	if u != nil {
		cc.SetParentID(*u)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetPicURL sets the "pic_url" field.
func (cc *CategoryCreate) SetPicURL(s string) *CategoryCreate {
	cc.mutation.SetPicURL(s)
	return cc
}

// SetBigPicURL sets the "big_pic_url" field.
func (cc *CategoryCreate) SetBigPicURL(s string) *CategoryCreate {
	cc.mutation.SetBigPicURL(s)
	return cc
}

// SetNillableBigPicURL sets the "big_pic_url" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableBigPicURL(s *string) *CategoryCreate {
	if s != nil {
		cc.SetBigPicURL(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CategoryCreate) SetID(u uint64) *CategoryCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddSpuIDs adds the "spus" edge to the Spu entity by IDs.
func (cc *CategoryCreate) AddSpuIDs(ids ...uint64) *CategoryCreate {
	cc.mutation.AddSpuIDs(ids...)
	return cc
}

// AddSpus adds the "spus" edges to the Spu entity.
func (cc *CategoryCreate) AddSpus(s ...*Spu) *CategoryCreate {
	ids := make([]uint64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSpuIDs(ids...)
}

// SetParent sets the "parent" edge to the Category entity.
func (cc *CategoryCreate) SetParent(c *Category) *CategoryCreate {
	return cc.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Category entity by IDs.
func (cc *CategoryCreate) AddChildIDs(ids ...uint64) *CategoryCreate {
	cc.mutation.AddChildIDs(ids...)
	return cc
}

// AddChildren adds the "children" edges to the Category entity.
func (cc *CategoryCreate) AddChildren(c ...*Category) *CategoryCreate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddChildIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CategoryCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if category.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized category.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := category.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if category.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized category.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := category.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := category.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.Sort(); !ok {
		v := category.DefaultSort
		cc.mutation.SetSort(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Category.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Category.updated_at"`)}
	}
	if _, ok := cc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Category.sort"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Category.name"`)}
	}
	if _, ok := cc.mutation.PicURL(); !ok {
		return &ValidationError{Name: "pic_url", err: errors.New(`ent: missing required field "Category.pic_url"`)}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(category.Table, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(category.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(category.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.Sort(); ok {
		_spec.SetField(category.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(category.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.PicURL(); ok {
		_spec.SetField(category.FieldPicURL, field.TypeString, value)
		_node.PicURL = value
	}
	if value, ok := cc.mutation.BigPicURL(); ok {
		_spec.SetField(category.FieldBigPicURL, field.TypeString, value)
		_node.BigPicURL = value
	}
	if nodes := cc.mutation.SpusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.SpusTable,
			Columns: []string{category.SpusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spu.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentTable,
			Columns: []string{category.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildrenTable,
			Columns: []string{category.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	err      error
	builders []*CategoryCreate
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
