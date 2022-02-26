// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appuserpurchaseamountsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
)

// AppUserPurchaseAmountSettingDelete is the builder for deleting a AppUserPurchaseAmountSetting entity.
type AppUserPurchaseAmountSettingDelete struct {
	config
	hooks    []Hook
	mutation *AppUserPurchaseAmountSettingMutation
}

// Where appends a list predicates to the AppUserPurchaseAmountSettingDelete builder.
func (aupasd *AppUserPurchaseAmountSettingDelete) Where(ps ...predicate.AppUserPurchaseAmountSetting) *AppUserPurchaseAmountSettingDelete {
	aupasd.mutation.Where(ps...)
	return aupasd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aupasd *AppUserPurchaseAmountSettingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aupasd.hooks) == 0 {
		affected, err = aupasd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserPurchaseAmountSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aupasd.mutation = mutation
			affected, err = aupasd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aupasd.hooks) - 1; i >= 0; i-- {
			if aupasd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aupasd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aupasd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aupasd *AppUserPurchaseAmountSettingDelete) ExecX(ctx context.Context) int {
	n, err := aupasd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aupasd *AppUserPurchaseAmountSettingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appuserpurchaseamountsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserpurchaseamountsetting.FieldID,
			},
		},
	}
	if ps := aupasd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aupasd.driver, _spec)
}

// AppUserPurchaseAmountSettingDeleteOne is the builder for deleting a single AppUserPurchaseAmountSetting entity.
type AppUserPurchaseAmountSettingDeleteOne struct {
	aupasd *AppUserPurchaseAmountSettingDelete
}

// Exec executes the deletion query.
func (aupasdo *AppUserPurchaseAmountSettingDeleteOne) Exec(ctx context.Context) error {
	n, err := aupasdo.aupasd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appuserpurchaseamountsetting.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aupasdo *AppUserPurchaseAmountSettingDeleteOne) ExecX(ctx context.Context) {
	aupasdo.aupasd.ExecX(ctx)
}