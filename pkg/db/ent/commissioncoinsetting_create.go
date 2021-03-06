// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/commissioncoinsetting"
	"github.com/google/uuid"
)

// CommissionCoinSettingCreate is the builder for creating a CommissionCoinSetting entity.
type CommissionCoinSettingCreate struct {
	config
	mutation *CommissionCoinSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ccsc *CommissionCoinSettingCreate) SetCoinTypeID(u uuid.UUID) *CommissionCoinSettingCreate {
	ccsc.mutation.SetCoinTypeID(u)
	return ccsc
}

// SetUsing sets the "using" field.
func (ccsc *CommissionCoinSettingCreate) SetUsing(b bool) *CommissionCoinSettingCreate {
	ccsc.mutation.SetUsing(b)
	return ccsc
}

// SetNillableUsing sets the "using" field if the given value is not nil.
func (ccsc *CommissionCoinSettingCreate) SetNillableUsing(b *bool) *CommissionCoinSettingCreate {
	if b != nil {
		ccsc.SetUsing(*b)
	}
	return ccsc
}

// SetCreateAt sets the "create_at" field.
func (ccsc *CommissionCoinSettingCreate) SetCreateAt(u uint32) *CommissionCoinSettingCreate {
	ccsc.mutation.SetCreateAt(u)
	return ccsc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ccsc *CommissionCoinSettingCreate) SetNillableCreateAt(u *uint32) *CommissionCoinSettingCreate {
	if u != nil {
		ccsc.SetCreateAt(*u)
	}
	return ccsc
}

// SetUpdateAt sets the "update_at" field.
func (ccsc *CommissionCoinSettingCreate) SetUpdateAt(u uint32) *CommissionCoinSettingCreate {
	ccsc.mutation.SetUpdateAt(u)
	return ccsc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (ccsc *CommissionCoinSettingCreate) SetNillableUpdateAt(u *uint32) *CommissionCoinSettingCreate {
	if u != nil {
		ccsc.SetUpdateAt(*u)
	}
	return ccsc
}

// SetDeleteAt sets the "delete_at" field.
func (ccsc *CommissionCoinSettingCreate) SetDeleteAt(u uint32) *CommissionCoinSettingCreate {
	ccsc.mutation.SetDeleteAt(u)
	return ccsc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ccsc *CommissionCoinSettingCreate) SetNillableDeleteAt(u *uint32) *CommissionCoinSettingCreate {
	if u != nil {
		ccsc.SetDeleteAt(*u)
	}
	return ccsc
}

// SetID sets the "id" field.
func (ccsc *CommissionCoinSettingCreate) SetID(u uuid.UUID) *CommissionCoinSettingCreate {
	ccsc.mutation.SetID(u)
	return ccsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ccsc *CommissionCoinSettingCreate) SetNillableID(u *uuid.UUID) *CommissionCoinSettingCreate {
	if u != nil {
		ccsc.SetID(*u)
	}
	return ccsc
}

// Mutation returns the CommissionCoinSettingMutation object of the builder.
func (ccsc *CommissionCoinSettingCreate) Mutation() *CommissionCoinSettingMutation {
	return ccsc.mutation
}

// Save creates the CommissionCoinSetting in the database.
func (ccsc *CommissionCoinSettingCreate) Save(ctx context.Context) (*CommissionCoinSetting, error) {
	var (
		err  error
		node *CommissionCoinSetting
	)
	ccsc.defaults()
	if len(ccsc.hooks) == 0 {
		if err = ccsc.check(); err != nil {
			return nil, err
		}
		node, err = ccsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommissionCoinSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccsc.check(); err != nil {
				return nil, err
			}
			ccsc.mutation = mutation
			if node, err = ccsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ccsc.hooks) - 1; i >= 0; i-- {
			if ccsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccsc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ccsc *CommissionCoinSettingCreate) SaveX(ctx context.Context) *CommissionCoinSetting {
	v, err := ccsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccsc *CommissionCoinSettingCreate) Exec(ctx context.Context) error {
	_, err := ccsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccsc *CommissionCoinSettingCreate) ExecX(ctx context.Context) {
	if err := ccsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccsc *CommissionCoinSettingCreate) defaults() {
	if _, ok := ccsc.mutation.Using(); !ok {
		v := commissioncoinsetting.DefaultUsing
		ccsc.mutation.SetUsing(v)
	}
	if _, ok := ccsc.mutation.CreateAt(); !ok {
		v := commissioncoinsetting.DefaultCreateAt()
		ccsc.mutation.SetCreateAt(v)
	}
	if _, ok := ccsc.mutation.UpdateAt(); !ok {
		v := commissioncoinsetting.DefaultUpdateAt()
		ccsc.mutation.SetUpdateAt(v)
	}
	if _, ok := ccsc.mutation.DeleteAt(); !ok {
		v := commissioncoinsetting.DefaultDeleteAt()
		ccsc.mutation.SetDeleteAt(v)
	}
	if _, ok := ccsc.mutation.ID(); !ok {
		v := commissioncoinsetting.DefaultID()
		ccsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccsc *CommissionCoinSettingCreate) check() error {
	if _, ok := ccsc.mutation.CoinTypeID(); !ok {
		return &ValidationError{Name: "coin_type_id", err: errors.New(`ent: missing required field "CommissionCoinSetting.coin_type_id"`)}
	}
	if _, ok := ccsc.mutation.Using(); !ok {
		return &ValidationError{Name: "using", err: errors.New(`ent: missing required field "CommissionCoinSetting.using"`)}
	}
	if _, ok := ccsc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "CommissionCoinSetting.create_at"`)}
	}
	if _, ok := ccsc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "CommissionCoinSetting.update_at"`)}
	}
	if _, ok := ccsc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "CommissionCoinSetting.delete_at"`)}
	}
	return nil
}

func (ccsc *CommissionCoinSettingCreate) sqlSave(ctx context.Context) (*CommissionCoinSetting, error) {
	_node, _spec := ccsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ccsc *CommissionCoinSettingCreate) createSpec() (*CommissionCoinSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &CommissionCoinSetting{config: ccsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: commissioncoinsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: commissioncoinsetting.FieldID,
			},
		}
	)
	_spec.OnConflict = ccsc.conflict
	if id, ok := ccsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ccsc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: commissioncoinsetting.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := ccsc.mutation.Using(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: commissioncoinsetting.FieldUsing,
		})
		_node.Using = value
	}
	if value, ok := ccsc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := ccsc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := ccsc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: commissioncoinsetting.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CommissionCoinSetting.Create().
//		SetCoinTypeID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CommissionCoinSettingUpsert) {
//			SetCoinTypeID(v+v).
//		}).
//		Exec(ctx)
//
func (ccsc *CommissionCoinSettingCreate) OnConflict(opts ...sql.ConflictOption) *CommissionCoinSettingUpsertOne {
	ccsc.conflict = opts
	return &CommissionCoinSettingUpsertOne{
		create: ccsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CommissionCoinSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccsc *CommissionCoinSettingCreate) OnConflictColumns(columns ...string) *CommissionCoinSettingUpsertOne {
	ccsc.conflict = append(ccsc.conflict, sql.ConflictColumns(columns...))
	return &CommissionCoinSettingUpsertOne{
		create: ccsc,
	}
}

type (
	// CommissionCoinSettingUpsertOne is the builder for "upsert"-ing
	//  one CommissionCoinSetting node.
	CommissionCoinSettingUpsertOne struct {
		create *CommissionCoinSettingCreate
	}

	// CommissionCoinSettingUpsert is the "OnConflict" setter.
	CommissionCoinSettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CommissionCoinSettingUpsert) SetCoinTypeID(v uuid.UUID) *CommissionCoinSettingUpsert {
	u.Set(commissioncoinsetting.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsert) UpdateCoinTypeID() *CommissionCoinSettingUpsert {
	u.SetExcluded(commissioncoinsetting.FieldCoinTypeID)
	return u
}

// SetUsing sets the "using" field.
func (u *CommissionCoinSettingUpsert) SetUsing(v bool) *CommissionCoinSettingUpsert {
	u.Set(commissioncoinsetting.FieldUsing, v)
	return u
}

// UpdateUsing sets the "using" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsert) UpdateUsing() *CommissionCoinSettingUpsert {
	u.SetExcluded(commissioncoinsetting.FieldUsing)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *CommissionCoinSettingUpsert) SetCreateAt(v uint32) *CommissionCoinSettingUpsert {
	u.Set(commissioncoinsetting.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsert) UpdateCreateAt() *CommissionCoinSettingUpsert {
	u.SetExcluded(commissioncoinsetting.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *CommissionCoinSettingUpsert) AddCreateAt(v uint32) *CommissionCoinSettingUpsert {
	u.Add(commissioncoinsetting.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *CommissionCoinSettingUpsert) SetUpdateAt(v uint32) *CommissionCoinSettingUpsert {
	u.Set(commissioncoinsetting.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsert) UpdateUpdateAt() *CommissionCoinSettingUpsert {
	u.SetExcluded(commissioncoinsetting.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *CommissionCoinSettingUpsert) AddUpdateAt(v uint32) *CommissionCoinSettingUpsert {
	u.Add(commissioncoinsetting.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *CommissionCoinSettingUpsert) SetDeleteAt(v uint32) *CommissionCoinSettingUpsert {
	u.Set(commissioncoinsetting.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsert) UpdateDeleteAt() *CommissionCoinSettingUpsert {
	u.SetExcluded(commissioncoinsetting.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *CommissionCoinSettingUpsert) AddDeleteAt(v uint32) *CommissionCoinSettingUpsert {
	u.Add(commissioncoinsetting.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CommissionCoinSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(commissioncoinsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CommissionCoinSettingUpsertOne) UpdateNewValues() *CommissionCoinSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(commissioncoinsetting.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CommissionCoinSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CommissionCoinSettingUpsertOne) Ignore() *CommissionCoinSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CommissionCoinSettingUpsertOne) DoNothing() *CommissionCoinSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CommissionCoinSettingCreate.OnConflict
// documentation for more info.
func (u *CommissionCoinSettingUpsertOne) Update(set func(*CommissionCoinSettingUpsert)) *CommissionCoinSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CommissionCoinSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CommissionCoinSettingUpsertOne) SetCoinTypeID(v uuid.UUID) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertOne) UpdateCoinTypeID() *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetUsing sets the "using" field.
func (u *CommissionCoinSettingUpsertOne) SetUsing(v bool) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetUsing(v)
	})
}

// UpdateUsing sets the "using" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertOne) UpdateUsing() *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateUsing()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CommissionCoinSettingUpsertOne) SetCreateAt(v uint32) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *CommissionCoinSettingUpsertOne) AddCreateAt(v uint32) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertOne) UpdateCreateAt() *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CommissionCoinSettingUpsertOne) SetUpdateAt(v uint32) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *CommissionCoinSettingUpsertOne) AddUpdateAt(v uint32) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertOne) UpdateUpdateAt() *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CommissionCoinSettingUpsertOne) SetDeleteAt(v uint32) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *CommissionCoinSettingUpsertOne) AddDeleteAt(v uint32) *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertOne) UpdateDeleteAt() *CommissionCoinSettingUpsertOne {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CommissionCoinSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CommissionCoinSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommissionCoinSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CommissionCoinSettingUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CommissionCoinSettingUpsertOne.ID is not supported by MySQL driver. Use CommissionCoinSettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CommissionCoinSettingUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CommissionCoinSettingCreateBulk is the builder for creating many CommissionCoinSetting entities in bulk.
type CommissionCoinSettingCreateBulk struct {
	config
	builders []*CommissionCoinSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the CommissionCoinSetting entities in the database.
func (ccscb *CommissionCoinSettingCreateBulk) Save(ctx context.Context) ([]*CommissionCoinSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccscb.builders))
	nodes := make([]*CommissionCoinSetting, len(ccscb.builders))
	mutators := make([]Mutator, len(ccscb.builders))
	for i := range ccscb.builders {
		func(i int, root context.Context) {
			builder := ccscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommissionCoinSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, ccscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccscb *CommissionCoinSettingCreateBulk) SaveX(ctx context.Context) []*CommissionCoinSetting {
	v, err := ccscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccscb *CommissionCoinSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := ccscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccscb *CommissionCoinSettingCreateBulk) ExecX(ctx context.Context) {
	if err := ccscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CommissionCoinSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CommissionCoinSettingUpsert) {
//			SetCoinTypeID(v+v).
//		}).
//		Exec(ctx)
//
func (ccscb *CommissionCoinSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *CommissionCoinSettingUpsertBulk {
	ccscb.conflict = opts
	return &CommissionCoinSettingUpsertBulk{
		create: ccscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CommissionCoinSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccscb *CommissionCoinSettingCreateBulk) OnConflictColumns(columns ...string) *CommissionCoinSettingUpsertBulk {
	ccscb.conflict = append(ccscb.conflict, sql.ConflictColumns(columns...))
	return &CommissionCoinSettingUpsertBulk{
		create: ccscb,
	}
}

// CommissionCoinSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of CommissionCoinSetting nodes.
type CommissionCoinSettingUpsertBulk struct {
	create *CommissionCoinSettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CommissionCoinSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(commissioncoinsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CommissionCoinSettingUpsertBulk) UpdateNewValues() *CommissionCoinSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(commissioncoinsetting.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CommissionCoinSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CommissionCoinSettingUpsertBulk) Ignore() *CommissionCoinSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CommissionCoinSettingUpsertBulk) DoNothing() *CommissionCoinSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CommissionCoinSettingCreateBulk.OnConflict
// documentation for more info.
func (u *CommissionCoinSettingUpsertBulk) Update(set func(*CommissionCoinSettingUpsert)) *CommissionCoinSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CommissionCoinSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CommissionCoinSettingUpsertBulk) SetCoinTypeID(v uuid.UUID) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertBulk) UpdateCoinTypeID() *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetUsing sets the "using" field.
func (u *CommissionCoinSettingUpsertBulk) SetUsing(v bool) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetUsing(v)
	})
}

// UpdateUsing sets the "using" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertBulk) UpdateUsing() *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateUsing()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CommissionCoinSettingUpsertBulk) SetCreateAt(v uint32) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *CommissionCoinSettingUpsertBulk) AddCreateAt(v uint32) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertBulk) UpdateCreateAt() *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CommissionCoinSettingUpsertBulk) SetUpdateAt(v uint32) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *CommissionCoinSettingUpsertBulk) AddUpdateAt(v uint32) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertBulk) UpdateUpdateAt() *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CommissionCoinSettingUpsertBulk) SetDeleteAt(v uint32) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *CommissionCoinSettingUpsertBulk) AddDeleteAt(v uint32) *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CommissionCoinSettingUpsertBulk) UpdateDeleteAt() *CommissionCoinSettingUpsertBulk {
	return u.Update(func(s *CommissionCoinSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CommissionCoinSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CommissionCoinSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CommissionCoinSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommissionCoinSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
