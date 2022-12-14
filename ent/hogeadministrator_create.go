// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/bug/ent/hoge"
	"entgo.io/bug/ent/hogeadministrator"
	"entgo.io/bug/ulid"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HogeAdministratorCreate is the builder for creating a HogeAdministrator entity.
type HogeAdministratorCreate struct {
	config
	mutation *HogeAdministratorMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (hac *HogeAdministratorCreate) SetCreatedAt(t time.Time) *HogeAdministratorCreate {
	hac.mutation.SetCreatedAt(t)
	return hac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hac *HogeAdministratorCreate) SetNillableCreatedAt(t *time.Time) *HogeAdministratorCreate {
	if t != nil {
		hac.SetCreatedAt(*t)
	}
	return hac
}

// SetUpdatedAt sets the "updated_at" field.
func (hac *HogeAdministratorCreate) SetUpdatedAt(t time.Time) *HogeAdministratorCreate {
	hac.mutation.SetUpdatedAt(t)
	return hac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hac *HogeAdministratorCreate) SetNillableUpdatedAt(t *time.Time) *HogeAdministratorCreate {
	if t != nil {
		hac.SetUpdatedAt(*t)
	}
	return hac
}

// SetFirstName sets the "first_name" field.
func (hac *HogeAdministratorCreate) SetFirstName(s string) *HogeAdministratorCreate {
	hac.mutation.SetFirstName(s)
	return hac
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (hac *HogeAdministratorCreate) SetNillableFirstName(s *string) *HogeAdministratorCreate {
	if s != nil {
		hac.SetFirstName(*s)
	}
	return hac
}

// SetLastName sets the "last_name" field.
func (hac *HogeAdministratorCreate) SetLastName(s string) *HogeAdministratorCreate {
	hac.mutation.SetLastName(s)
	return hac
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (hac *HogeAdministratorCreate) SetNillableLastName(s *string) *HogeAdministratorCreate {
	if s != nil {
		hac.SetLastName(*s)
	}
	return hac
}

// SetEmail sets the "email" field.
func (hac *HogeAdministratorCreate) SetEmail(s string) *HogeAdministratorCreate {
	hac.mutation.SetEmail(s)
	return hac
}

// SetIsActive sets the "is_active" field.
func (hac *HogeAdministratorCreate) SetIsActive(b bool) *HogeAdministratorCreate {
	hac.mutation.SetIsActive(b)
	return hac
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (hac *HogeAdministratorCreate) SetNillableIsActive(b *bool) *HogeAdministratorCreate {
	if b != nil {
		hac.SetIsActive(*b)
	}
	return hac
}

// SetID sets the "id" field.
func (hac *HogeAdministratorCreate) SetID(u ulid.ID) *HogeAdministratorCreate {
	hac.mutation.SetID(u)
	return hac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hac *HogeAdministratorCreate) SetNillableID(u *ulid.ID) *HogeAdministratorCreate {
	if u != nil {
		hac.SetID(*u)
	}
	return hac
}

// SetHogeID sets the "hoge" edge to the Hoge entity by ID.
func (hac *HogeAdministratorCreate) SetHogeID(id ulid.ID) *HogeAdministratorCreate {
	hac.mutation.SetHogeID(id)
	return hac
}

// SetNillableHogeID sets the "hoge" edge to the Hoge entity by ID if the given value is not nil.
func (hac *HogeAdministratorCreate) SetNillableHogeID(id *ulid.ID) *HogeAdministratorCreate {
	if id != nil {
		hac = hac.SetHogeID(*id)
	}
	return hac
}

// SetHoge sets the "hoge" edge to the Hoge entity.
func (hac *HogeAdministratorCreate) SetHoge(h *Hoge) *HogeAdministratorCreate {
	return hac.SetHogeID(h.ID)
}

// Mutation returns the HogeAdministratorMutation object of the builder.
func (hac *HogeAdministratorCreate) Mutation() *HogeAdministratorMutation {
	return hac.mutation
}

// Save creates the HogeAdministrator in the database.
func (hac *HogeAdministratorCreate) Save(ctx context.Context) (*HogeAdministrator, error) {
	var (
		err  error
		node *HogeAdministrator
	)
	hac.defaults()
	if len(hac.hooks) == 0 {
		if err = hac.check(); err != nil {
			return nil, err
		}
		node, err = hac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HogeAdministratorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hac.check(); err != nil {
				return nil, err
			}
			hac.mutation = mutation
			if node, err = hac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hac.hooks) - 1; i >= 0; i-- {
			if hac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*HogeAdministrator)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HogeAdministratorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hac *HogeAdministratorCreate) SaveX(ctx context.Context) *HogeAdministrator {
	v, err := hac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hac *HogeAdministratorCreate) Exec(ctx context.Context) error {
	_, err := hac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hac *HogeAdministratorCreate) ExecX(ctx context.Context) {
	if err := hac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hac *HogeAdministratorCreate) defaults() {
	if _, ok := hac.mutation.CreatedAt(); !ok {
		v := hogeadministrator.DefaultCreatedAt()
		hac.mutation.SetCreatedAt(v)
	}
	if _, ok := hac.mutation.UpdatedAt(); !ok {
		v := hogeadministrator.DefaultUpdatedAt()
		hac.mutation.SetUpdatedAt(v)
	}
	if _, ok := hac.mutation.IsActive(); !ok {
		v := hogeadministrator.DefaultIsActive
		hac.mutation.SetIsActive(v)
	}
	if _, ok := hac.mutation.ID(); !ok {
		v := hogeadministrator.DefaultID()
		hac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hac *HogeAdministratorCreate) check() error {
	if _, ok := hac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HogeAdministrator.created_at"`)}
	}
	if _, ok := hac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "HogeAdministrator.updated_at"`)}
	}
	if v, ok := hac.mutation.FirstName(); ok {
		if err := hogeadministrator.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "HogeAdministrator.first_name": %w`, err)}
		}
	}
	if v, ok := hac.mutation.LastName(); ok {
		if err := hogeadministrator.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "HogeAdministrator.last_name": %w`, err)}
		}
	}
	if _, ok := hac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "HogeAdministrator.email"`)}
	}
	if v, ok := hac.mutation.Email(); ok {
		if err := hogeadministrator.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "HogeAdministrator.email": %w`, err)}
		}
	}
	if _, ok := hac.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "HogeAdministrator.is_active"`)}
	}
	return nil
}

func (hac *HogeAdministratorCreate) sqlSave(ctx context.Context) (*HogeAdministrator, error) {
	_node, _spec := hac.createSpec()
	if err := sqlgraph.CreateNode(ctx, hac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(ulid.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected HogeAdministrator.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (hac *HogeAdministratorCreate) createSpec() (*HogeAdministrator, *sqlgraph.CreateSpec) {
	var (
		_node = &HogeAdministrator{config: hac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hogeadministrator.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: hogeadministrator.FieldID,
			},
		}
	)
	if id, ok := hac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hogeadministrator.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := hac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: hogeadministrator.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := hac.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hogeadministrator.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := hac.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hogeadministrator.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := hac.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hogeadministrator.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := hac.mutation.IsActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: hogeadministrator.FieldIsActive,
		})
		_node.IsActive = value
	}
	if nodes := hac.mutation.HogeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hogeadministrator.HogeTable,
			Columns: []string{hogeadministrator.HogeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: hoge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.hoge_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HogeAdministratorCreateBulk is the builder for creating many HogeAdministrator entities in bulk.
type HogeAdministratorCreateBulk struct {
	config
	builders []*HogeAdministratorCreate
}

// Save creates the HogeAdministrator entities in the database.
func (hacb *HogeAdministratorCreateBulk) Save(ctx context.Context) ([]*HogeAdministrator, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hacb.builders))
	nodes := make([]*HogeAdministrator, len(hacb.builders))
	mutators := make([]Mutator, len(hacb.builders))
	for i := range hacb.builders {
		func(i int, root context.Context) {
			builder := hacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HogeAdministratorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, hacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hacb *HogeAdministratorCreateBulk) SaveX(ctx context.Context) []*HogeAdministrator {
	v, err := hacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hacb *HogeAdministratorCreateBulk) Exec(ctx context.Context) error {
	_, err := hacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hacb *HogeAdministratorCreateBulk) ExecX(ctx context.Context) {
	if err := hacb.Exec(ctx); err != nil {
		panic(err)
	}
}
