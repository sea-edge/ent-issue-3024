// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/bug/ent/hogeadministrator"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HogeAdministratorDelete is the builder for deleting a HogeAdministrator entity.
type HogeAdministratorDelete struct {
	config
	hooks    []Hook
	mutation *HogeAdministratorMutation
}

// Where appends a list predicates to the HogeAdministratorDelete builder.
func (had *HogeAdministratorDelete) Where(ps ...predicate.HogeAdministrator) *HogeAdministratorDelete {
	had.mutation.Where(ps...)
	return had
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (had *HogeAdministratorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(had.hooks) == 0 {
		affected, err = had.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HogeAdministratorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			had.mutation = mutation
			affected, err = had.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(had.hooks) - 1; i >= 0; i-- {
			if had.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = had.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, had.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (had *HogeAdministratorDelete) ExecX(ctx context.Context) int {
	n, err := had.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (had *HogeAdministratorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: hogeadministrator.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: hogeadministrator.FieldID,
			},
		},
	}
	if ps := had.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, had.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// HogeAdministratorDeleteOne is the builder for deleting a single HogeAdministrator entity.
type HogeAdministratorDeleteOne struct {
	had *HogeAdministratorDelete
}

// Exec executes the deletion query.
func (hado *HogeAdministratorDeleteOne) Exec(ctx context.Context) error {
	n, err := hado.had.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hogeadministrator.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hado *HogeAdministratorDeleteOne) ExecX(ctx context.Context) {
	hado.had.ExecX(ctx)
}
