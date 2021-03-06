// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/activity"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ActivityUpdate is the builder for updating Activity entities.
type ActivityUpdate struct {
	config
	hooks    []Hook
	mutation *ActivityMutation
}

// Where appends a list predicates to the ActivityUpdate builder.
func (au *ActivityUpdate) Where(ps ...predicate.Activity) *ActivityUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAppID sets the "app_id" field.
func (au *ActivityUpdate) SetAppID(u uuid.UUID) *ActivityUpdate {
	au.mutation.SetAppID(u)
	return au
}

// SetCreatedBy sets the "created_by" field.
func (au *ActivityUpdate) SetCreatedBy(u uuid.UUID) *ActivityUpdate {
	au.mutation.SetCreatedBy(u)
	return au
}

// SetName sets the "name" field.
func (au *ActivityUpdate) SetName(s string) *ActivityUpdate {
	au.mutation.SetName(s)
	return au
}

// SetStart sets the "start" field.
func (au *ActivityUpdate) SetStart(u uint32) *ActivityUpdate {
	au.mutation.ResetStart()
	au.mutation.SetStart(u)
	return au
}

// AddStart adds u to the "start" field.
func (au *ActivityUpdate) AddStart(u int32) *ActivityUpdate {
	au.mutation.AddStart(u)
	return au
}

// SetEnd sets the "end" field.
func (au *ActivityUpdate) SetEnd(u uint32) *ActivityUpdate {
	au.mutation.ResetEnd()
	au.mutation.SetEnd(u)
	return au
}

// AddEnd adds u to the "end" field.
func (au *ActivityUpdate) AddEnd(u int32) *ActivityUpdate {
	au.mutation.AddEnd(u)
	return au
}

// SetSystemActivity sets the "system_activity" field.
func (au *ActivityUpdate) SetSystemActivity(b bool) *ActivityUpdate {
	au.mutation.SetSystemActivity(b)
	return au
}

// SetCreateAt sets the "create_at" field.
func (au *ActivityUpdate) SetCreateAt(u uint32) *ActivityUpdate {
	au.mutation.ResetCreateAt()
	au.mutation.SetCreateAt(u)
	return au
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (au *ActivityUpdate) SetNillableCreateAt(u *uint32) *ActivityUpdate {
	if u != nil {
		au.SetCreateAt(*u)
	}
	return au
}

// AddCreateAt adds u to the "create_at" field.
func (au *ActivityUpdate) AddCreateAt(u int32) *ActivityUpdate {
	au.mutation.AddCreateAt(u)
	return au
}

// SetUpdateAt sets the "update_at" field.
func (au *ActivityUpdate) SetUpdateAt(u uint32) *ActivityUpdate {
	au.mutation.ResetUpdateAt()
	au.mutation.SetUpdateAt(u)
	return au
}

// AddUpdateAt adds u to the "update_at" field.
func (au *ActivityUpdate) AddUpdateAt(u int32) *ActivityUpdate {
	au.mutation.AddUpdateAt(u)
	return au
}

// SetDeleteAt sets the "delete_at" field.
func (au *ActivityUpdate) SetDeleteAt(u uint32) *ActivityUpdate {
	au.mutation.ResetDeleteAt()
	au.mutation.SetDeleteAt(u)
	return au
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (au *ActivityUpdate) SetNillableDeleteAt(u *uint32) *ActivityUpdate {
	if u != nil {
		au.SetDeleteAt(*u)
	}
	return au
}

// AddDeleteAt adds u to the "delete_at" field.
func (au *ActivityUpdate) AddDeleteAt(u int32) *ActivityUpdate {
	au.mutation.AddDeleteAt(u)
	return au
}

// Mutation returns the ActivityMutation object of the builder.
func (au *ActivityUpdate) Mutation() *ActivityMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ActivityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActivityUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActivityUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ActivityUpdate) defaults() {
	if _, ok := au.mutation.UpdateAt(); !ok {
		v := activity.UpdateDefaultUpdateAt()
		au.mutation.SetUpdateAt(v)
	}
}

func (au *ActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: activity.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: activity.FieldAppID,
		})
	}
	if value, ok := au.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: activity.FieldCreatedBy,
		})
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldName,
		})
	}
	if value, ok := au.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldStart,
		})
	}
	if value, ok := au.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldStart,
		})
	}
	if value, ok := au.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldEnd,
		})
	}
	if value, ok := au.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldEnd,
		})
	}
	if value, ok := au.mutation.SystemActivity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: activity.FieldSystemActivity,
		})
	}
	if value, ok := au.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldCreateAt,
		})
	}
	if value, ok := au.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldCreateAt,
		})
	}
	if value, ok := au.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldUpdateAt,
		})
	}
	if value, ok := au.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldUpdateAt,
		})
	}
	if value, ok := au.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldDeleteAt,
		})
	}
	if value, ok := au.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ActivityUpdateOne is the builder for updating a single Activity entity.
type ActivityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActivityMutation
}

// SetAppID sets the "app_id" field.
func (auo *ActivityUpdateOne) SetAppID(u uuid.UUID) *ActivityUpdateOne {
	auo.mutation.SetAppID(u)
	return auo
}

// SetCreatedBy sets the "created_by" field.
func (auo *ActivityUpdateOne) SetCreatedBy(u uuid.UUID) *ActivityUpdateOne {
	auo.mutation.SetCreatedBy(u)
	return auo
}

// SetName sets the "name" field.
func (auo *ActivityUpdateOne) SetName(s string) *ActivityUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetStart sets the "start" field.
func (auo *ActivityUpdateOne) SetStart(u uint32) *ActivityUpdateOne {
	auo.mutation.ResetStart()
	auo.mutation.SetStart(u)
	return auo
}

// AddStart adds u to the "start" field.
func (auo *ActivityUpdateOne) AddStart(u int32) *ActivityUpdateOne {
	auo.mutation.AddStart(u)
	return auo
}

// SetEnd sets the "end" field.
func (auo *ActivityUpdateOne) SetEnd(u uint32) *ActivityUpdateOne {
	auo.mutation.ResetEnd()
	auo.mutation.SetEnd(u)
	return auo
}

// AddEnd adds u to the "end" field.
func (auo *ActivityUpdateOne) AddEnd(u int32) *ActivityUpdateOne {
	auo.mutation.AddEnd(u)
	return auo
}

// SetSystemActivity sets the "system_activity" field.
func (auo *ActivityUpdateOne) SetSystemActivity(b bool) *ActivityUpdateOne {
	auo.mutation.SetSystemActivity(b)
	return auo
}

// SetCreateAt sets the "create_at" field.
func (auo *ActivityUpdateOne) SetCreateAt(u uint32) *ActivityUpdateOne {
	auo.mutation.ResetCreateAt()
	auo.mutation.SetCreateAt(u)
	return auo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableCreateAt(u *uint32) *ActivityUpdateOne {
	if u != nil {
		auo.SetCreateAt(*u)
	}
	return auo
}

// AddCreateAt adds u to the "create_at" field.
func (auo *ActivityUpdateOne) AddCreateAt(u int32) *ActivityUpdateOne {
	auo.mutation.AddCreateAt(u)
	return auo
}

// SetUpdateAt sets the "update_at" field.
func (auo *ActivityUpdateOne) SetUpdateAt(u uint32) *ActivityUpdateOne {
	auo.mutation.ResetUpdateAt()
	auo.mutation.SetUpdateAt(u)
	return auo
}

// AddUpdateAt adds u to the "update_at" field.
func (auo *ActivityUpdateOne) AddUpdateAt(u int32) *ActivityUpdateOne {
	auo.mutation.AddUpdateAt(u)
	return auo
}

// SetDeleteAt sets the "delete_at" field.
func (auo *ActivityUpdateOne) SetDeleteAt(u uint32) *ActivityUpdateOne {
	auo.mutation.ResetDeleteAt()
	auo.mutation.SetDeleteAt(u)
	return auo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableDeleteAt(u *uint32) *ActivityUpdateOne {
	if u != nil {
		auo.SetDeleteAt(*u)
	}
	return auo
}

// AddDeleteAt adds u to the "delete_at" field.
func (auo *ActivityUpdateOne) AddDeleteAt(u int32) *ActivityUpdateOne {
	auo.mutation.AddDeleteAt(u)
	return auo
}

// Mutation returns the ActivityMutation object of the builder.
func (auo *ActivityUpdateOne) Mutation() *ActivityMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ActivityUpdateOne) Select(field string, fields ...string) *ActivityUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Activity entity.
func (auo *ActivityUpdateOne) Save(ctx context.Context) (*Activity, error) {
	var (
		err  error
		node *Activity
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Activity)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ActivityMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ActivityUpdateOne) SaveX(ctx context.Context) *Activity {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActivityUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ActivityUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdateAt(); !ok {
		v := activity.UpdateDefaultUpdateAt()
		auo.mutation.SetUpdateAt(v)
	}
}

func (auo *ActivityUpdateOne) sqlSave(ctx context.Context) (_node *Activity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: activity.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Activity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activity.FieldID)
		for _, f := range fields {
			if !activity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != activity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: activity.FieldAppID,
		})
	}
	if value, ok := auo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: activity.FieldCreatedBy,
		})
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldName,
		})
	}
	if value, ok := auo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldStart,
		})
	}
	if value, ok := auo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldStart,
		})
	}
	if value, ok := auo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldEnd,
		})
	}
	if value, ok := auo.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldEnd,
		})
	}
	if value, ok := auo.mutation.SystemActivity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: activity.FieldSystemActivity,
		})
	}
	if value, ok := auo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldCreateAt,
		})
	}
	if value, ok := auo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldCreateAt,
		})
	}
	if value, ok := auo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldUpdateAt,
		})
	}
	if value, ok := auo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldUpdateAt,
		})
	}
	if value, ok := auo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldDeleteAt,
		})
	}
	if value, ok := auo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: activity.FieldDeleteAt,
		})
	}
	_node = &Activity{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
