// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userspecialreduction"
	"github.com/google/uuid"
)

// UserSpecialReductionUpdate is the builder for updating UserSpecialReduction entities.
type UserSpecialReductionUpdate struct {
	config
	hooks    []Hook
	mutation *UserSpecialReductionMutation
}

// Where appends a list predicates to the UserSpecialReductionUpdate builder.
func (usru *UserSpecialReductionUpdate) Where(ps ...predicate.UserSpecialReduction) *UserSpecialReductionUpdate {
	usru.mutation.Where(ps...)
	return usru
}

// SetAppID sets the "app_id" field.
func (usru *UserSpecialReductionUpdate) SetAppID(u uuid.UUID) *UserSpecialReductionUpdate {
	usru.mutation.SetAppID(u)
	return usru
}

// SetUserID sets the "user_id" field.
func (usru *UserSpecialReductionUpdate) SetUserID(u uuid.UUID) *UserSpecialReductionUpdate {
	usru.mutation.SetUserID(u)
	return usru
}

// SetAmount sets the "amount" field.
func (usru *UserSpecialReductionUpdate) SetAmount(u uint64) *UserSpecialReductionUpdate {
	usru.mutation.ResetAmount()
	usru.mutation.SetAmount(u)
	return usru
}

// AddAmount adds u to the "amount" field.
func (usru *UserSpecialReductionUpdate) AddAmount(u uint64) *UserSpecialReductionUpdate {
	usru.mutation.AddAmount(u)
	return usru
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (usru *UserSpecialReductionUpdate) SetReleaseByUserID(u uuid.UUID) *UserSpecialReductionUpdate {
	usru.mutation.SetReleaseByUserID(u)
	return usru
}

// SetStart sets the "start" field.
func (usru *UserSpecialReductionUpdate) SetStart(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.ResetStart()
	usru.mutation.SetStart(u)
	return usru
}

// AddStart adds u to the "start" field.
func (usru *UserSpecialReductionUpdate) AddStart(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.AddStart(u)
	return usru
}

// SetDurationDays sets the "duration_days" field.
func (usru *UserSpecialReductionUpdate) SetDurationDays(i int32) *UserSpecialReductionUpdate {
	usru.mutation.ResetDurationDays()
	usru.mutation.SetDurationDays(i)
	return usru
}

// AddDurationDays adds i to the "duration_days" field.
func (usru *UserSpecialReductionUpdate) AddDurationDays(i int32) *UserSpecialReductionUpdate {
	usru.mutation.AddDurationDays(i)
	return usru
}

// SetMessage sets the "message" field.
func (usru *UserSpecialReductionUpdate) SetMessage(s string) *UserSpecialReductionUpdate {
	usru.mutation.SetMessage(s)
	return usru
}

// SetCreateAt sets the "create_at" field.
func (usru *UserSpecialReductionUpdate) SetCreateAt(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.ResetCreateAt()
	usru.mutation.SetCreateAt(u)
	return usru
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (usru *UserSpecialReductionUpdate) SetNillableCreateAt(u *uint32) *UserSpecialReductionUpdate {
	if u != nil {
		usru.SetCreateAt(*u)
	}
	return usru
}

// AddCreateAt adds u to the "create_at" field.
func (usru *UserSpecialReductionUpdate) AddCreateAt(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.AddCreateAt(u)
	return usru
}

// SetUpdateAt sets the "update_at" field.
func (usru *UserSpecialReductionUpdate) SetUpdateAt(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.ResetUpdateAt()
	usru.mutation.SetUpdateAt(u)
	return usru
}

// AddUpdateAt adds u to the "update_at" field.
func (usru *UserSpecialReductionUpdate) AddUpdateAt(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.AddUpdateAt(u)
	return usru
}

// SetDeleteAt sets the "delete_at" field.
func (usru *UserSpecialReductionUpdate) SetDeleteAt(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.ResetDeleteAt()
	usru.mutation.SetDeleteAt(u)
	return usru
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (usru *UserSpecialReductionUpdate) SetNillableDeleteAt(u *uint32) *UserSpecialReductionUpdate {
	if u != nil {
		usru.SetDeleteAt(*u)
	}
	return usru
}

// AddDeleteAt adds u to the "delete_at" field.
func (usru *UserSpecialReductionUpdate) AddDeleteAt(u uint32) *UserSpecialReductionUpdate {
	usru.mutation.AddDeleteAt(u)
	return usru
}

// Mutation returns the UserSpecialReductionMutation object of the builder.
func (usru *UserSpecialReductionUpdate) Mutation() *UserSpecialReductionMutation {
	return usru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usru *UserSpecialReductionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	usru.defaults()
	if len(usru.hooks) == 0 {
		if err = usru.check(); err != nil {
			return 0, err
		}
		affected, err = usru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSpecialReductionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usru.check(); err != nil {
				return 0, err
			}
			usru.mutation = mutation
			affected, err = usru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(usru.hooks) - 1; i >= 0; i-- {
			if usru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (usru *UserSpecialReductionUpdate) SaveX(ctx context.Context) int {
	affected, err := usru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usru *UserSpecialReductionUpdate) Exec(ctx context.Context) error {
	_, err := usru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usru *UserSpecialReductionUpdate) ExecX(ctx context.Context) {
	if err := usru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usru *UserSpecialReductionUpdate) defaults() {
	if _, ok := usru.mutation.UpdateAt(); !ok {
		v := userspecialreduction.UpdateDefaultUpdateAt()
		usru.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usru *UserSpecialReductionUpdate) check() error {
	if v, ok := usru.mutation.Message(); ok {
		if err := userspecialreduction.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf("ent: validator failed for field \"message\": %w", err)}
		}
	}
	return nil
}

func (usru *UserSpecialReductionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userspecialreduction.Table,
			Columns: userspecialreduction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userspecialreduction.FieldID,
			},
		},
	}
	if ps := usru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usru.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldAppID,
		})
	}
	if value, ok := usru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldUserID,
		})
	}
	if value, ok := usru.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userspecialreduction.FieldAmount,
		})
	}
	if value, ok := usru.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userspecialreduction.FieldAmount,
		})
	}
	if value, ok := usru.mutation.ReleaseByUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldReleaseByUserID,
		})
	}
	if value, ok := usru.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldStart,
		})
	}
	if value, ok := usru.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldStart,
		})
	}
	if value, ok := usru.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userspecialreduction.FieldDurationDays,
		})
	}
	if value, ok := usru.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userspecialreduction.FieldDurationDays,
		})
	}
	if value, ok := usru.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userspecialreduction.FieldMessage,
		})
	}
	if value, ok := usru.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldCreateAt,
		})
	}
	if value, ok := usru.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldCreateAt,
		})
	}
	if value, ok := usru.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldUpdateAt,
		})
	}
	if value, ok := usru.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldUpdateAt,
		})
	}
	if value, ok := usru.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldDeleteAt,
		})
	}
	if value, ok := usru.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userspecialreduction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserSpecialReductionUpdateOne is the builder for updating a single UserSpecialReduction entity.
type UserSpecialReductionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserSpecialReductionMutation
}

// SetAppID sets the "app_id" field.
func (usruo *UserSpecialReductionUpdateOne) SetAppID(u uuid.UUID) *UserSpecialReductionUpdateOne {
	usruo.mutation.SetAppID(u)
	return usruo
}

// SetUserID sets the "user_id" field.
func (usruo *UserSpecialReductionUpdateOne) SetUserID(u uuid.UUID) *UserSpecialReductionUpdateOne {
	usruo.mutation.SetUserID(u)
	return usruo
}

// SetAmount sets the "amount" field.
func (usruo *UserSpecialReductionUpdateOne) SetAmount(u uint64) *UserSpecialReductionUpdateOne {
	usruo.mutation.ResetAmount()
	usruo.mutation.SetAmount(u)
	return usruo
}

// AddAmount adds u to the "amount" field.
func (usruo *UserSpecialReductionUpdateOne) AddAmount(u uint64) *UserSpecialReductionUpdateOne {
	usruo.mutation.AddAmount(u)
	return usruo
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (usruo *UserSpecialReductionUpdateOne) SetReleaseByUserID(u uuid.UUID) *UserSpecialReductionUpdateOne {
	usruo.mutation.SetReleaseByUserID(u)
	return usruo
}

// SetStart sets the "start" field.
func (usruo *UserSpecialReductionUpdateOne) SetStart(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.ResetStart()
	usruo.mutation.SetStart(u)
	return usruo
}

// AddStart adds u to the "start" field.
func (usruo *UserSpecialReductionUpdateOne) AddStart(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.AddStart(u)
	return usruo
}

// SetDurationDays sets the "duration_days" field.
func (usruo *UserSpecialReductionUpdateOne) SetDurationDays(i int32) *UserSpecialReductionUpdateOne {
	usruo.mutation.ResetDurationDays()
	usruo.mutation.SetDurationDays(i)
	return usruo
}

// AddDurationDays adds i to the "duration_days" field.
func (usruo *UserSpecialReductionUpdateOne) AddDurationDays(i int32) *UserSpecialReductionUpdateOne {
	usruo.mutation.AddDurationDays(i)
	return usruo
}

// SetMessage sets the "message" field.
func (usruo *UserSpecialReductionUpdateOne) SetMessage(s string) *UserSpecialReductionUpdateOne {
	usruo.mutation.SetMessage(s)
	return usruo
}

// SetCreateAt sets the "create_at" field.
func (usruo *UserSpecialReductionUpdateOne) SetCreateAt(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.ResetCreateAt()
	usruo.mutation.SetCreateAt(u)
	return usruo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (usruo *UserSpecialReductionUpdateOne) SetNillableCreateAt(u *uint32) *UserSpecialReductionUpdateOne {
	if u != nil {
		usruo.SetCreateAt(*u)
	}
	return usruo
}

// AddCreateAt adds u to the "create_at" field.
func (usruo *UserSpecialReductionUpdateOne) AddCreateAt(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.AddCreateAt(u)
	return usruo
}

// SetUpdateAt sets the "update_at" field.
func (usruo *UserSpecialReductionUpdateOne) SetUpdateAt(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.ResetUpdateAt()
	usruo.mutation.SetUpdateAt(u)
	return usruo
}

// AddUpdateAt adds u to the "update_at" field.
func (usruo *UserSpecialReductionUpdateOne) AddUpdateAt(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.AddUpdateAt(u)
	return usruo
}

// SetDeleteAt sets the "delete_at" field.
func (usruo *UserSpecialReductionUpdateOne) SetDeleteAt(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.ResetDeleteAt()
	usruo.mutation.SetDeleteAt(u)
	return usruo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (usruo *UserSpecialReductionUpdateOne) SetNillableDeleteAt(u *uint32) *UserSpecialReductionUpdateOne {
	if u != nil {
		usruo.SetDeleteAt(*u)
	}
	return usruo
}

// AddDeleteAt adds u to the "delete_at" field.
func (usruo *UserSpecialReductionUpdateOne) AddDeleteAt(u uint32) *UserSpecialReductionUpdateOne {
	usruo.mutation.AddDeleteAt(u)
	return usruo
}

// Mutation returns the UserSpecialReductionMutation object of the builder.
func (usruo *UserSpecialReductionUpdateOne) Mutation() *UserSpecialReductionMutation {
	return usruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usruo *UserSpecialReductionUpdateOne) Select(field string, fields ...string) *UserSpecialReductionUpdateOne {
	usruo.fields = append([]string{field}, fields...)
	return usruo
}

// Save executes the query and returns the updated UserSpecialReduction entity.
func (usruo *UserSpecialReductionUpdateOne) Save(ctx context.Context) (*UserSpecialReduction, error) {
	var (
		err  error
		node *UserSpecialReduction
	)
	usruo.defaults()
	if len(usruo.hooks) == 0 {
		if err = usruo.check(); err != nil {
			return nil, err
		}
		node, err = usruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSpecialReductionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usruo.check(); err != nil {
				return nil, err
			}
			usruo.mutation = mutation
			node, err = usruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(usruo.hooks) - 1; i >= 0; i-- {
			if usruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (usruo *UserSpecialReductionUpdateOne) SaveX(ctx context.Context) *UserSpecialReduction {
	node, err := usruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usruo *UserSpecialReductionUpdateOne) Exec(ctx context.Context) error {
	_, err := usruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usruo *UserSpecialReductionUpdateOne) ExecX(ctx context.Context) {
	if err := usruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usruo *UserSpecialReductionUpdateOne) defaults() {
	if _, ok := usruo.mutation.UpdateAt(); !ok {
		v := userspecialreduction.UpdateDefaultUpdateAt()
		usruo.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usruo *UserSpecialReductionUpdateOne) check() error {
	if v, ok := usruo.mutation.Message(); ok {
		if err := userspecialreduction.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf("ent: validator failed for field \"message\": %w", err)}
		}
	}
	return nil
}

func (usruo *UserSpecialReductionUpdateOne) sqlSave(ctx context.Context) (_node *UserSpecialReduction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userspecialreduction.Table,
			Columns: userspecialreduction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userspecialreduction.FieldID,
			},
		},
	}
	id, ok := usruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserSpecialReduction.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := usruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userspecialreduction.FieldID)
		for _, f := range fields {
			if !userspecialreduction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userspecialreduction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usruo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldAppID,
		})
	}
	if value, ok := usruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldUserID,
		})
	}
	if value, ok := usruo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userspecialreduction.FieldAmount,
		})
	}
	if value, ok := usruo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userspecialreduction.FieldAmount,
		})
	}
	if value, ok := usruo.mutation.ReleaseByUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldReleaseByUserID,
		})
	}
	if value, ok := usruo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldStart,
		})
	}
	if value, ok := usruo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldStart,
		})
	}
	if value, ok := usruo.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userspecialreduction.FieldDurationDays,
		})
	}
	if value, ok := usruo.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userspecialreduction.FieldDurationDays,
		})
	}
	if value, ok := usruo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userspecialreduction.FieldMessage,
		})
	}
	if value, ok := usruo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldCreateAt,
		})
	}
	if value, ok := usruo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldCreateAt,
		})
	}
	if value, ok := usruo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldUpdateAt,
		})
	}
	if value, ok := usruo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldUpdateAt,
		})
	}
	if value, ok := usruo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldDeleteAt,
		})
	}
	if value, ok := usruo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldDeleteAt,
		})
	}
	_node = &UserSpecialReduction{config: usruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userspecialreduction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
