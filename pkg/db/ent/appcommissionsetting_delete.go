// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcommissionsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
)

// AppCommissionSettingDelete is the builder for deleting a AppCommissionSetting entity.
type AppCommissionSettingDelete struct {
	config
	hooks    []Hook
	mutation *AppCommissionSettingMutation
}

// Where appends a list predicates to the AppCommissionSettingDelete builder.
func (acsd *AppCommissionSettingDelete) Where(ps ...predicate.AppCommissionSetting) *AppCommissionSettingDelete {
	acsd.mutation.Where(ps...)
	return acsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acsd *AppCommissionSettingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(acsd.hooks) == 0 {
		affected, err = acsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppCommissionSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acsd.mutation = mutation
			affected, err = acsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acsd.hooks) - 1; i >= 0; i-- {
			if acsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (acsd *AppCommissionSettingDelete) ExecX(ctx context.Context) int {
	n, err := acsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acsd *AppCommissionSettingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appcommissionsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcommissionsetting.FieldID,
			},
		},
	}
	if ps := acsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppCommissionSettingDeleteOne is the builder for deleting a single AppCommissionSetting entity.
type AppCommissionSettingDeleteOne struct {
	acsd *AppCommissionSettingDelete
}

// Exec executes the deletion query.
func (acsdo *AppCommissionSettingDeleteOne) Exec(ctx context.Context) error {
	n, err := acsdo.acsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appcommissionsetting.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acsdo *AppCommissionSettingDeleteOne) ExecX(ctx context.Context) {
	acsdo.acsd.ExecX(ctx)
}
