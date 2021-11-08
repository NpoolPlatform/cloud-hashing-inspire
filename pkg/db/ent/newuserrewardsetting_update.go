// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
)

// NewUserRewardSettingUpdate is the builder for updating NewUserRewardSetting entities.
type NewUserRewardSettingUpdate struct {
	config
	hooks    []Hook
	mutation *NewUserRewardSettingMutation
}

// Where appends a list predicates to the NewUserRewardSettingUpdate builder.
func (nursu *NewUserRewardSettingUpdate) Where(ps ...predicate.NewUserRewardSetting) *NewUserRewardSettingUpdate {
	nursu.mutation.Where(ps...)
	return nursu
}

// Mutation returns the NewUserRewardSettingMutation object of the builder.
func (nursu *NewUserRewardSettingUpdate) Mutation() *NewUserRewardSettingMutation {
	return nursu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nursu *NewUserRewardSettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nursu.hooks) == 0 {
		affected, err = nursu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewUserRewardSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nursu.mutation = mutation
			affected, err = nursu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nursu.hooks) - 1; i >= 0; i-- {
			if nursu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nursu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nursu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nursu *NewUserRewardSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := nursu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nursu *NewUserRewardSettingUpdate) Exec(ctx context.Context) error {
	_, err := nursu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nursu *NewUserRewardSettingUpdate) ExecX(ctx context.Context) {
	if err := nursu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nursu *NewUserRewardSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newuserrewardsetting.Table,
			Columns: newuserrewardsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newuserrewardsetting.FieldID,
			},
		},
	}
	if ps := nursu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nursu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newuserrewardsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NewUserRewardSettingUpdateOne is the builder for updating a single NewUserRewardSetting entity.
type NewUserRewardSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewUserRewardSettingMutation
}

// Mutation returns the NewUserRewardSettingMutation object of the builder.
func (nursuo *NewUserRewardSettingUpdateOne) Mutation() *NewUserRewardSettingMutation {
	return nursuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nursuo *NewUserRewardSettingUpdateOne) Select(field string, fields ...string) *NewUserRewardSettingUpdateOne {
	nursuo.fields = append([]string{field}, fields...)
	return nursuo
}

// Save executes the query and returns the updated NewUserRewardSetting entity.
func (nursuo *NewUserRewardSettingUpdateOne) Save(ctx context.Context) (*NewUserRewardSetting, error) {
	var (
		err  error
		node *NewUserRewardSetting
	)
	if len(nursuo.hooks) == 0 {
		node, err = nursuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewUserRewardSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nursuo.mutation = mutation
			node, err = nursuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nursuo.hooks) - 1; i >= 0; i-- {
			if nursuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nursuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nursuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nursuo *NewUserRewardSettingUpdateOne) SaveX(ctx context.Context) *NewUserRewardSetting {
	node, err := nursuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nursuo *NewUserRewardSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := nursuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nursuo *NewUserRewardSettingUpdateOne) ExecX(ctx context.Context) {
	if err := nursuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nursuo *NewUserRewardSettingUpdateOne) sqlSave(ctx context.Context) (_node *NewUserRewardSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newuserrewardsetting.Table,
			Columns: newuserrewardsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: newuserrewardsetting.FieldID,
			},
		},
	}
	id, ok := nursuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing NewUserRewardSetting.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := nursuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, newuserrewardsetting.FieldID)
		for _, f := range fields {
			if !newuserrewardsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != newuserrewardsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nursuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &NewUserRewardSetting{config: nursuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nursuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newuserrewardsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}