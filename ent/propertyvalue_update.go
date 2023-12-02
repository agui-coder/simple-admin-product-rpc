// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-product-rpc/ent/property"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PropertyValueUpdate is the builder for updating PropertyValue entities.
type PropertyValueUpdate struct {
	config
	hooks    []Hook
	mutation *PropertyValueMutation
}

// Where appends a list predicates to the PropertyValueUpdate builder.
func (pvu *PropertyValueUpdate) Where(ps ...predicate.PropertyValue) *PropertyValueUpdate {
	pvu.mutation.Where(ps...)
	return pvu
}

// SetUpdatedAt sets the "updated_at" field.
func (pvu *PropertyValueUpdate) SetUpdatedAt(t time.Time) *PropertyValueUpdate {
	pvu.mutation.SetUpdatedAt(t)
	return pvu
}

// SetDeletedAt sets the "deleted_at" field.
func (pvu *PropertyValueUpdate) SetDeletedAt(t time.Time) *PropertyValueUpdate {
	pvu.mutation.SetDeletedAt(t)
	return pvu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pvu *PropertyValueUpdate) SetNillableDeletedAt(t *time.Time) *PropertyValueUpdate {
	if t != nil {
		pvu.SetDeletedAt(*t)
	}
	return pvu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pvu *PropertyValueUpdate) ClearDeletedAt() *PropertyValueUpdate {
	pvu.mutation.ClearDeletedAt()
	return pvu
}

// SetPropertyID sets the "property_id" field.
func (pvu *PropertyValueUpdate) SetPropertyID(u uint64) *PropertyValueUpdate {
	pvu.mutation.SetPropertyID(u)
	return pvu
}

// SetNillablePropertyID sets the "property_id" field if the given value is not nil.
func (pvu *PropertyValueUpdate) SetNillablePropertyID(u *uint64) *PropertyValueUpdate {
	if u != nil {
		pvu.SetPropertyID(*u)
	}
	return pvu
}

// ClearPropertyID clears the value of the "property_id" field.
func (pvu *PropertyValueUpdate) ClearPropertyID() *PropertyValueUpdate {
	pvu.mutation.ClearPropertyID()
	return pvu
}

// SetName sets the "name" field.
func (pvu *PropertyValueUpdate) SetName(s string) *PropertyValueUpdate {
	pvu.mutation.SetName(s)
	return pvu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pvu *PropertyValueUpdate) SetNillableName(s *string) *PropertyValueUpdate {
	if s != nil {
		pvu.SetName(*s)
	}
	return pvu
}

// ClearName clears the value of the "name" field.
func (pvu *PropertyValueUpdate) ClearName() *PropertyValueUpdate {
	pvu.mutation.ClearName()
	return pvu
}

// SetRemark sets the "remark" field.
func (pvu *PropertyValueUpdate) SetRemark(s string) *PropertyValueUpdate {
	pvu.mutation.SetRemark(s)
	return pvu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pvu *PropertyValueUpdate) SetNillableRemark(s *string) *PropertyValueUpdate {
	if s != nil {
		pvu.SetRemark(*s)
	}
	return pvu
}

// ClearRemark clears the value of the "remark" field.
func (pvu *PropertyValueUpdate) ClearRemark() *PropertyValueUpdate {
	pvu.mutation.ClearRemark()
	return pvu
}

// SetPropertysID sets the "propertys" edge to the Property entity by ID.
func (pvu *PropertyValueUpdate) SetPropertysID(id uint64) *PropertyValueUpdate {
	pvu.mutation.SetPropertysID(id)
	return pvu
}

// SetNillablePropertysID sets the "propertys" edge to the Property entity by ID if the given value is not nil.
func (pvu *PropertyValueUpdate) SetNillablePropertysID(id *uint64) *PropertyValueUpdate {
	if id != nil {
		pvu = pvu.SetPropertysID(*id)
	}
	return pvu
}

// SetPropertys sets the "propertys" edge to the Property entity.
func (pvu *PropertyValueUpdate) SetPropertys(p *Property) *PropertyValueUpdate {
	return pvu.SetPropertysID(p.ID)
}

// Mutation returns the PropertyValueMutation object of the builder.
func (pvu *PropertyValueUpdate) Mutation() *PropertyValueMutation {
	return pvu.mutation
}

// ClearPropertys clears the "propertys" edge to the Property entity.
func (pvu *PropertyValueUpdate) ClearPropertys() *PropertyValueUpdate {
	pvu.mutation.ClearPropertys()
	return pvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pvu *PropertyValueUpdate) Save(ctx context.Context) (int, error) {
	if err := pvu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, pvu.sqlSave, pvu.mutation, pvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvu *PropertyValueUpdate) SaveX(ctx context.Context) int {
	affected, err := pvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pvu *PropertyValueUpdate) Exec(ctx context.Context) error {
	_, err := pvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvu *PropertyValueUpdate) ExecX(ctx context.Context) {
	if err := pvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pvu *PropertyValueUpdate) defaults() error {
	if _, ok := pvu.mutation.UpdatedAt(); !ok {
		if propertyvalue.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized propertyvalue.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := propertyvalue.UpdateDefaultUpdatedAt()
		pvu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pvu *PropertyValueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(propertyvalue.Table, propertyvalue.Columns, sqlgraph.NewFieldSpec(propertyvalue.FieldID, field.TypeUint64))
	if ps := pvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvu.mutation.UpdatedAt(); ok {
		_spec.SetField(propertyvalue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pvu.mutation.DeletedAt(); ok {
		_spec.SetField(propertyvalue.FieldDeletedAt, field.TypeTime, value)
	}
	if pvu.mutation.DeletedAtCleared() {
		_spec.ClearField(propertyvalue.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pvu.mutation.Name(); ok {
		_spec.SetField(propertyvalue.FieldName, field.TypeString, value)
	}
	if pvu.mutation.NameCleared() {
		_spec.ClearField(propertyvalue.FieldName, field.TypeString)
	}
	if value, ok := pvu.mutation.Remark(); ok {
		_spec.SetField(propertyvalue.FieldRemark, field.TypeString, value)
	}
	if pvu.mutation.RemarkCleared() {
		_spec.ClearField(propertyvalue.FieldRemark, field.TypeString)
	}
	if pvu.mutation.PropertysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   propertyvalue.PropertysTable,
			Columns: []string{propertyvalue.PropertysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.PropertysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   propertyvalue.PropertysTable,
			Columns: []string{propertyvalue.PropertysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{propertyvalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pvu.mutation.done = true
	return n, nil
}

// PropertyValueUpdateOne is the builder for updating a single PropertyValue entity.
type PropertyValueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PropertyValueMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pvuo *PropertyValueUpdateOne) SetUpdatedAt(t time.Time) *PropertyValueUpdateOne {
	pvuo.mutation.SetUpdatedAt(t)
	return pvuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pvuo *PropertyValueUpdateOne) SetDeletedAt(t time.Time) *PropertyValueUpdateOne {
	pvuo.mutation.SetDeletedAt(t)
	return pvuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pvuo *PropertyValueUpdateOne) SetNillableDeletedAt(t *time.Time) *PropertyValueUpdateOne {
	if t != nil {
		pvuo.SetDeletedAt(*t)
	}
	return pvuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pvuo *PropertyValueUpdateOne) ClearDeletedAt() *PropertyValueUpdateOne {
	pvuo.mutation.ClearDeletedAt()
	return pvuo
}

// SetPropertyID sets the "property_id" field.
func (pvuo *PropertyValueUpdateOne) SetPropertyID(u uint64) *PropertyValueUpdateOne {
	pvuo.mutation.SetPropertyID(u)
	return pvuo
}

// SetNillablePropertyID sets the "property_id" field if the given value is not nil.
func (pvuo *PropertyValueUpdateOne) SetNillablePropertyID(u *uint64) *PropertyValueUpdateOne {
	if u != nil {
		pvuo.SetPropertyID(*u)
	}
	return pvuo
}

// ClearPropertyID clears the value of the "property_id" field.
func (pvuo *PropertyValueUpdateOne) ClearPropertyID() *PropertyValueUpdateOne {
	pvuo.mutation.ClearPropertyID()
	return pvuo
}

// SetName sets the "name" field.
func (pvuo *PropertyValueUpdateOne) SetName(s string) *PropertyValueUpdateOne {
	pvuo.mutation.SetName(s)
	return pvuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pvuo *PropertyValueUpdateOne) SetNillableName(s *string) *PropertyValueUpdateOne {
	if s != nil {
		pvuo.SetName(*s)
	}
	return pvuo
}

// ClearName clears the value of the "name" field.
func (pvuo *PropertyValueUpdateOne) ClearName() *PropertyValueUpdateOne {
	pvuo.mutation.ClearName()
	return pvuo
}

// SetRemark sets the "remark" field.
func (pvuo *PropertyValueUpdateOne) SetRemark(s string) *PropertyValueUpdateOne {
	pvuo.mutation.SetRemark(s)
	return pvuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pvuo *PropertyValueUpdateOne) SetNillableRemark(s *string) *PropertyValueUpdateOne {
	if s != nil {
		pvuo.SetRemark(*s)
	}
	return pvuo
}

// ClearRemark clears the value of the "remark" field.
func (pvuo *PropertyValueUpdateOne) ClearRemark() *PropertyValueUpdateOne {
	pvuo.mutation.ClearRemark()
	return pvuo
}

// SetPropertysID sets the "propertys" edge to the Property entity by ID.
func (pvuo *PropertyValueUpdateOne) SetPropertysID(id uint64) *PropertyValueUpdateOne {
	pvuo.mutation.SetPropertysID(id)
	return pvuo
}

// SetNillablePropertysID sets the "propertys" edge to the Property entity by ID if the given value is not nil.
func (pvuo *PropertyValueUpdateOne) SetNillablePropertysID(id *uint64) *PropertyValueUpdateOne {
	if id != nil {
		pvuo = pvuo.SetPropertysID(*id)
	}
	return pvuo
}

// SetPropertys sets the "propertys" edge to the Property entity.
func (pvuo *PropertyValueUpdateOne) SetPropertys(p *Property) *PropertyValueUpdateOne {
	return pvuo.SetPropertysID(p.ID)
}

// Mutation returns the PropertyValueMutation object of the builder.
func (pvuo *PropertyValueUpdateOne) Mutation() *PropertyValueMutation {
	return pvuo.mutation
}

// ClearPropertys clears the "propertys" edge to the Property entity.
func (pvuo *PropertyValueUpdateOne) ClearPropertys() *PropertyValueUpdateOne {
	pvuo.mutation.ClearPropertys()
	return pvuo
}

// Where appends a list predicates to the PropertyValueUpdate builder.
func (pvuo *PropertyValueUpdateOne) Where(ps ...predicate.PropertyValue) *PropertyValueUpdateOne {
	pvuo.mutation.Where(ps...)
	return pvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pvuo *PropertyValueUpdateOne) Select(field string, fields ...string) *PropertyValueUpdateOne {
	pvuo.fields = append([]string{field}, fields...)
	return pvuo
}

// Save executes the query and returns the updated PropertyValue entity.
func (pvuo *PropertyValueUpdateOne) Save(ctx context.Context) (*PropertyValue, error) {
	if err := pvuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pvuo.sqlSave, pvuo.mutation, pvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pvuo *PropertyValueUpdateOne) SaveX(ctx context.Context) *PropertyValue {
	node, err := pvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pvuo *PropertyValueUpdateOne) Exec(ctx context.Context) error {
	_, err := pvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvuo *PropertyValueUpdateOne) ExecX(ctx context.Context) {
	if err := pvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pvuo *PropertyValueUpdateOne) defaults() error {
	if _, ok := pvuo.mutation.UpdatedAt(); !ok {
		if propertyvalue.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized propertyvalue.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := propertyvalue.UpdateDefaultUpdatedAt()
		pvuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (pvuo *PropertyValueUpdateOne) sqlSave(ctx context.Context) (_node *PropertyValue, err error) {
	_spec := sqlgraph.NewUpdateSpec(propertyvalue.Table, propertyvalue.Columns, sqlgraph.NewFieldSpec(propertyvalue.FieldID, field.TypeUint64))
	id, ok := pvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PropertyValue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, propertyvalue.FieldID)
		for _, f := range fields {
			if !propertyvalue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != propertyvalue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvuo.mutation.UpdatedAt(); ok {
		_spec.SetField(propertyvalue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pvuo.mutation.DeletedAt(); ok {
		_spec.SetField(propertyvalue.FieldDeletedAt, field.TypeTime, value)
	}
	if pvuo.mutation.DeletedAtCleared() {
		_spec.ClearField(propertyvalue.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pvuo.mutation.Name(); ok {
		_spec.SetField(propertyvalue.FieldName, field.TypeString, value)
	}
	if pvuo.mutation.NameCleared() {
		_spec.ClearField(propertyvalue.FieldName, field.TypeString)
	}
	if value, ok := pvuo.mutation.Remark(); ok {
		_spec.SetField(propertyvalue.FieldRemark, field.TypeString, value)
	}
	if pvuo.mutation.RemarkCleared() {
		_spec.ClearField(propertyvalue.FieldRemark, field.TypeString)
	}
	if pvuo.mutation.PropertysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   propertyvalue.PropertysTable,
			Columns: []string{propertyvalue.PropertysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.PropertysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   propertyvalue.PropertysTable,
			Columns: []string{propertyvalue.PropertysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PropertyValue{config: pvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{propertyvalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pvuo.mutation.done = true
	return _node, nil
}