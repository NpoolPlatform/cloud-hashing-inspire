// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/commissioncoinsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CommissionCoinSettingUpdate is the builder for updating CommissionCoinSetting entities.
type CommissionCoinSettingUpdate struct {
	config
	hooks    []Hook
	mutation *CommissionCoinSettingMutation
}

// Where appends a list predicates to the CommissionCoinSettingUpdate builder.
func (ccsu *CommissionCoinSettingUpdate) Where(ps ...predicate.CommissionCoinSetting) *CommissionCoinSettingUpdate {
	ccsu.mutation.Where(ps...)
	return ccsu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ccsu *CommissionCoinSettingUpdate) SetCoinTypeID(u uuid.UUID) *CommissionCoinSettingUpdate {
	ccsu.mutation.SetCoinTypeID(u)
	return ccsu
}

// SetUsing sets the "using" field.
func (ccsu *CommissionCoinSettingUpdate) SetUsing(b bool) *CommissionCoinSettingUpdate {
	ccsu.mutation.SetUsing(b)
	return ccsu
}

// SetNillableUsing sets the "using" field if the given value is not nil.
func (ccsu *CommissionCoinSettingUpdate) SetNillableUsing(b *bool) *CommissionCoinSettingUpdate {
	if b != nil {
		ccsu.SetUsing(*b)
	}
	return ccsu
}

// SetCreateAt sets the "create_at" field.
func (ccsu *CommissionCoinSettingUpdate) SetCreateAt(u uint32) *CommissionCoinSettingUpdate {
	ccsu.mutation.ResetCreateAt()
	ccsu.mutation.SetCreateAt(u)
	return ccsu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ccsu *CommissionCoinSettingUpdate) SetNillableCreateAt(u *uint32) *CommissionCoinSettingUpdate {
	if u != nil {
		ccsu.SetCreateAt(*u)
	}
	return ccsu
}

// AddCreateAt adds u to the "create_at" field.
func (ccsu *CommissionCoinSettingUpdate) AddCreateAt(u int32) *CommissionCoinSettingUpdate {
	ccsu.mutation.AddCreateAt(u)
	return ccsu
}

// SetUpdateAt sets the "update_at" field.
func (ccsu *CommissionCoinSettingUpdate) SetUpdateAt(u uint32) *CommissionCoinSettingUpdate {
	ccsu.mutation.ResetUpdateAt()
	ccsu.mutation.SetUpdateAt(u)
	return ccsu
}

// AddUpdateAt adds u to the "update_at" field.
func (ccsu *CommissionCoinSettingUpdate) AddUpdateAt(u int32) *CommissionCoinSettingUpdate {
	ccsu.mutation.AddUpdateAt(u)
	return ccsu
}

// SetDeleteAt sets the "delete_at" field.
func (ccsu *CommissionCoinSettingUpdate) SetDeleteAt(u uint32) *CommissionCoinSettingUpdate {
	ccsu.mutation.ResetDeleteAt()
	ccsu.mutation.SetDeleteAt(u)
	return ccsu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ccsu *CommissionCoinSettingUpdate) SetNillableDeleteAt(u *uint32) *CommissionCoinSettingUpdate {
	if u != nil {
		ccsu.SetDeleteAt(*u)
	}
	return ccsu
}

// AddDeleteAt adds u to the "delete_at" field.
func (ccsu *CommissionCoinSettingUpdate) AddDeleteAt(u int32) *CommissionCoinSettingUpdate {
	ccsu.mutation.AddDeleteAt(u)
	return ccsu
}

// Mutation returns the CommissionCoinSettingMutation object of the builder.
func (ccsu *CommissionCoinSettingUpdate) Mutation() *CommissionCoinSettingMutation {
	return ccsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccsu *CommissionCoinSettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ccsu.defaults()
	if len(ccsu.hooks) == 0 {
		affected, err = ccsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommissionCoinSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccsu.mutation = mutation
			affected, err = ccsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccsu.hooks) - 1; i >= 0; i-- {
			if ccsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccsu *CommissionCoinSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := ccsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccsu *CommissionCoinSettingUpdate) Exec(ctx context.Context) error {
	_, err := ccsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccsu *CommissionCoinSettingUpdate) ExecX(ctx context.Context) {
	if err := ccsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccsu *CommissionCoinSettingUpdate) defaults() {
	if _, ok := ccsu.mutation.UpdateAt(); !ok {
		v := commissioncoinsetting.UpdateDefaultUpdateAt()
		ccsu.mutation.SetUpdateAt(v)
	}
}

func (ccsu *CommissionCoinSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commissioncoinsetting.Table,
			Columns: commissioncoinsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: commissioncoinsetting.FieldID,
			},
		},
	}
	if ps := ccsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccsu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commissioncoinsetting.FieldCoinTypeID,
		})
	}
	if value, ok := ccsu.mutation.Using(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: commissioncoinsetting.FieldUsing,
		})
	}
	if value, ok := ccsu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldCreateAt,
		})
	}
	if value, ok := ccsu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldCreateAt,
		})
	}
	if value, ok := ccsu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldUpdateAt,
		})
	}
	if value, ok := ccsu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldUpdateAt,
		})
	}
	if value, ok := ccsu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldDeleteAt,
		})
	}
	if value, ok := ccsu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissioncoinsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CommissionCoinSettingUpdateOne is the builder for updating a single CommissionCoinSetting entity.
type CommissionCoinSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommissionCoinSettingMutation
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ccsuo *CommissionCoinSettingUpdateOne) SetCoinTypeID(u uuid.UUID) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.SetCoinTypeID(u)
	return ccsuo
}

// SetUsing sets the "using" field.
func (ccsuo *CommissionCoinSettingUpdateOne) SetUsing(b bool) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.SetUsing(b)
	return ccsuo
}

// SetNillableUsing sets the "using" field if the given value is not nil.
func (ccsuo *CommissionCoinSettingUpdateOne) SetNillableUsing(b *bool) *CommissionCoinSettingUpdateOne {
	if b != nil {
		ccsuo.SetUsing(*b)
	}
	return ccsuo
}

// SetCreateAt sets the "create_at" field.
func (ccsuo *CommissionCoinSettingUpdateOne) SetCreateAt(u uint32) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.ResetCreateAt()
	ccsuo.mutation.SetCreateAt(u)
	return ccsuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ccsuo *CommissionCoinSettingUpdateOne) SetNillableCreateAt(u *uint32) *CommissionCoinSettingUpdateOne {
	if u != nil {
		ccsuo.SetCreateAt(*u)
	}
	return ccsuo
}

// AddCreateAt adds u to the "create_at" field.
func (ccsuo *CommissionCoinSettingUpdateOne) AddCreateAt(u int32) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.AddCreateAt(u)
	return ccsuo
}

// SetUpdateAt sets the "update_at" field.
func (ccsuo *CommissionCoinSettingUpdateOne) SetUpdateAt(u uint32) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.ResetUpdateAt()
	ccsuo.mutation.SetUpdateAt(u)
	return ccsuo
}

// AddUpdateAt adds u to the "update_at" field.
func (ccsuo *CommissionCoinSettingUpdateOne) AddUpdateAt(u int32) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.AddUpdateAt(u)
	return ccsuo
}

// SetDeleteAt sets the "delete_at" field.
func (ccsuo *CommissionCoinSettingUpdateOne) SetDeleteAt(u uint32) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.ResetDeleteAt()
	ccsuo.mutation.SetDeleteAt(u)
	return ccsuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ccsuo *CommissionCoinSettingUpdateOne) SetNillableDeleteAt(u *uint32) *CommissionCoinSettingUpdateOne {
	if u != nil {
		ccsuo.SetDeleteAt(*u)
	}
	return ccsuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (ccsuo *CommissionCoinSettingUpdateOne) AddDeleteAt(u int32) *CommissionCoinSettingUpdateOne {
	ccsuo.mutation.AddDeleteAt(u)
	return ccsuo
}

// Mutation returns the CommissionCoinSettingMutation object of the builder.
func (ccsuo *CommissionCoinSettingUpdateOne) Mutation() *CommissionCoinSettingMutation {
	return ccsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccsuo *CommissionCoinSettingUpdateOne) Select(field string, fields ...string) *CommissionCoinSettingUpdateOne {
	ccsuo.fields = append([]string{field}, fields...)
	return ccsuo
}

// Save executes the query and returns the updated CommissionCoinSetting entity.
func (ccsuo *CommissionCoinSettingUpdateOne) Save(ctx context.Context) (*CommissionCoinSetting, error) {
	var (
		err  error
		node *CommissionCoinSetting
	)
	ccsuo.defaults()
	if len(ccsuo.hooks) == 0 {
		node, err = ccsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommissionCoinSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccsuo.mutation = mutation
			node, err = ccsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ccsuo.hooks) - 1; i >= 0; i-- {
			if ccsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CommissionCoinSetting)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CommissionCoinSettingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccsuo *CommissionCoinSettingUpdateOne) SaveX(ctx context.Context) *CommissionCoinSetting {
	node, err := ccsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccsuo *CommissionCoinSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := ccsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccsuo *CommissionCoinSettingUpdateOne) ExecX(ctx context.Context) {
	if err := ccsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccsuo *CommissionCoinSettingUpdateOne) defaults() {
	if _, ok := ccsuo.mutation.UpdateAt(); !ok {
		v := commissioncoinsetting.UpdateDefaultUpdateAt()
		ccsuo.mutation.SetUpdateAt(v)
	}
}

func (ccsuo *CommissionCoinSettingUpdateOne) sqlSave(ctx context.Context) (_node *CommissionCoinSetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   commissioncoinsetting.Table,
			Columns: commissioncoinsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: commissioncoinsetting.FieldID,
			},
		},
	}
	id, ok := ccsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CommissionCoinSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commissioncoinsetting.FieldID)
		for _, f := range fields {
			if !commissioncoinsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commissioncoinsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccsuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commissioncoinsetting.FieldCoinTypeID,
		})
	}
	if value, ok := ccsuo.mutation.Using(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: commissioncoinsetting.FieldUsing,
		})
	}
	if value, ok := ccsuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldCreateAt,
		})
	}
	if value, ok := ccsuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldCreateAt,
		})
	}
	if value, ok := ccsuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldUpdateAt,
		})
	}
	if value, ok := ccsuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldUpdateAt,
		})
	}
	if value, ok := ccsuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldDeleteAt,
		})
	}
	if value, ok := ccsuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldDeleteAt,
		})
	}
	_node = &CommissionCoinSetting{config: ccsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commissioncoinsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
