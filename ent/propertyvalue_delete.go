// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-product-rpc/ent/propertyvalue"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PropertyValueDelete is the builder for deleting a PropertyValue entity.
type PropertyValueDelete struct {
	config
	hooks    []Hook
	mutation *PropertyValueMutation
}

// Where appends a list predicates to the PropertyValueDelete builder.
func (pvd *PropertyValueDelete) Where(ps ...predicate.PropertyValue) *PropertyValueDelete {
	pvd.mutation.Where(ps...)
	return pvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pvd *PropertyValueDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pvd.sqlExec, pvd.mutation, pvd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pvd *PropertyValueDelete) ExecX(ctx context.Context) int {
	n, err := pvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pvd *PropertyValueDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(propertyvalue.Table, sqlgraph.NewFieldSpec(propertyvalue.FieldID, field.TypeUint64))
	if ps := pvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pvd.mutation.done = true
	return affected, err
}

// PropertyValueDeleteOne is the builder for deleting a single PropertyValue entity.
type PropertyValueDeleteOne struct {
	pvd *PropertyValueDelete
}

// Where appends a list predicates to the PropertyValueDelete builder.
func (pvdo *PropertyValueDeleteOne) Where(ps ...predicate.PropertyValue) *PropertyValueDeleteOne {
	pvdo.pvd.mutation.Where(ps...)
	return pvdo
}

// Exec executes the deletion query.
func (pvdo *PropertyValueDeleteOne) Exec(ctx context.Context) error {
	n, err := pvdo.pvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{propertyvalue.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pvdo *PropertyValueDeleteOne) ExecX(ctx context.Context) {
	if err := pvdo.Exec(ctx); err != nil {
		panic(err)
	}
}