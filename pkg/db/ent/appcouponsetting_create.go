// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcouponsetting"
	"github.com/google/uuid"
)

// AppCouponSettingCreate is the builder for creating a AppCouponSetting entity.
type AppCouponSettingCreate struct {
	config
	mutation *AppCouponSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (acsc *AppCouponSettingCreate) SetAppID(u uuid.UUID) *AppCouponSettingCreate {
	acsc.mutation.SetAppID(u)
	return acsc
}

// SetDominationLimit sets the "domination_limit" field.
func (acsc *AppCouponSettingCreate) SetDominationLimit(u uint64) *AppCouponSettingCreate {
	acsc.mutation.SetDominationLimit(u)
	return acsc
}

// SetTotalLimit sets the "total_limit" field.
func (acsc *AppCouponSettingCreate) SetTotalLimit(i int32) *AppCouponSettingCreate {
	acsc.mutation.SetTotalLimit(i)
	return acsc
}

// SetCreateAt sets the "create_at" field.
func (acsc *AppCouponSettingCreate) SetCreateAt(u uint32) *AppCouponSettingCreate {
	acsc.mutation.SetCreateAt(u)
	return acsc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (acsc *AppCouponSettingCreate) SetNillableCreateAt(u *uint32) *AppCouponSettingCreate {
	if u != nil {
		acsc.SetCreateAt(*u)
	}
	return acsc
}

// SetUpdateAt sets the "update_at" field.
func (acsc *AppCouponSettingCreate) SetUpdateAt(u uint32) *AppCouponSettingCreate {
	acsc.mutation.SetUpdateAt(u)
	return acsc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (acsc *AppCouponSettingCreate) SetNillableUpdateAt(u *uint32) *AppCouponSettingCreate {
	if u != nil {
		acsc.SetUpdateAt(*u)
	}
	return acsc
}

// SetDeleteAt sets the "delete_at" field.
func (acsc *AppCouponSettingCreate) SetDeleteAt(u uint32) *AppCouponSettingCreate {
	acsc.mutation.SetDeleteAt(u)
	return acsc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (acsc *AppCouponSettingCreate) SetNillableDeleteAt(u *uint32) *AppCouponSettingCreate {
	if u != nil {
		acsc.SetDeleteAt(*u)
	}
	return acsc
}

// SetID sets the "id" field.
func (acsc *AppCouponSettingCreate) SetID(u uuid.UUID) *AppCouponSettingCreate {
	acsc.mutation.SetID(u)
	return acsc
}

// Mutation returns the AppCouponSettingMutation object of the builder.
func (acsc *AppCouponSettingCreate) Mutation() *AppCouponSettingMutation {
	return acsc.mutation
}

// Save creates the AppCouponSetting in the database.
func (acsc *AppCouponSettingCreate) Save(ctx context.Context) (*AppCouponSetting, error) {
	var (
		err  error
		node *AppCouponSetting
	)
	acsc.defaults()
	if len(acsc.hooks) == 0 {
		if err = acsc.check(); err != nil {
			return nil, err
		}
		node, err = acsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppCouponSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acsc.check(); err != nil {
				return nil, err
			}
			acsc.mutation = mutation
			if node, err = acsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(acsc.hooks) - 1; i >= 0; i-- {
			if acsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (acsc *AppCouponSettingCreate) SaveX(ctx context.Context) *AppCouponSetting {
	v, err := acsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acsc *AppCouponSettingCreate) Exec(ctx context.Context) error {
	_, err := acsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acsc *AppCouponSettingCreate) ExecX(ctx context.Context) {
	if err := acsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acsc *AppCouponSettingCreate) defaults() {
	if _, ok := acsc.mutation.CreateAt(); !ok {
		v := appcouponsetting.DefaultCreateAt()
		acsc.mutation.SetCreateAt(v)
	}
	if _, ok := acsc.mutation.UpdateAt(); !ok {
		v := appcouponsetting.DefaultUpdateAt()
		acsc.mutation.SetUpdateAt(v)
	}
	if _, ok := acsc.mutation.DeleteAt(); !ok {
		v := appcouponsetting.DefaultDeleteAt()
		acsc.mutation.SetDeleteAt(v)
	}
	if _, ok := acsc.mutation.ID(); !ok {
		v := appcouponsetting.DefaultID()
		acsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acsc *AppCouponSettingCreate) check() error {
	if _, ok := acsc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := acsc.mutation.DominationLimit(); !ok {
		return &ValidationError{Name: "domination_limit", err: errors.New(`ent: missing required field "domination_limit"`)}
	}
	if _, ok := acsc.mutation.TotalLimit(); !ok {
		return &ValidationError{Name: "total_limit", err: errors.New(`ent: missing required field "total_limit"`)}
	}
	if _, ok := acsc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := acsc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := acsc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (acsc *AppCouponSettingCreate) sqlSave(ctx context.Context) (*AppCouponSetting, error) {
	_node, _spec := acsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (acsc *AppCouponSettingCreate) createSpec() (*AppCouponSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &AppCouponSetting{config: acsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appcouponsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcouponsetting.FieldID,
			},
		}
	)
	_spec.OnConflict = acsc.conflict
	if id, ok := acsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := acsc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcouponsetting.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := acsc.mutation.DominationLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: appcouponsetting.FieldDominationLimit,
		})
		_node.DominationLimit = value
	}
	if value, ok := acsc.mutation.TotalLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: appcouponsetting.FieldTotalLimit,
		})
		_node.TotalLimit = value
	}
	if value, ok := acsc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcouponsetting.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := acsc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcouponsetting.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := acsc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcouponsetting.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppCouponSetting.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppCouponSettingUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (acsc *AppCouponSettingCreate) OnConflict(opts ...sql.ConflictOption) *AppCouponSettingUpsertOne {
	acsc.conflict = opts
	return &AppCouponSettingUpsertOne{
		create: acsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppCouponSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acsc *AppCouponSettingCreate) OnConflictColumns(columns ...string) *AppCouponSettingUpsertOne {
	acsc.conflict = append(acsc.conflict, sql.ConflictColumns(columns...))
	return &AppCouponSettingUpsertOne{
		create: acsc,
	}
}

type (
	// AppCouponSettingUpsertOne is the builder for "upsert"-ing
	//  one AppCouponSetting node.
	AppCouponSettingUpsertOne struct {
		create *AppCouponSettingCreate
	}

	// AppCouponSettingUpsert is the "OnConflict" setter.
	AppCouponSettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *AppCouponSettingUpsert) SetAppID(v uuid.UUID) *AppCouponSettingUpsert {
	u.Set(appcouponsetting.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppCouponSettingUpsert) UpdateAppID() *AppCouponSettingUpsert {
	u.SetExcluded(appcouponsetting.FieldAppID)
	return u
}

// SetDominationLimit sets the "domination_limit" field.
func (u *AppCouponSettingUpsert) SetDominationLimit(v uint64) *AppCouponSettingUpsert {
	u.Set(appcouponsetting.FieldDominationLimit, v)
	return u
}

// UpdateDominationLimit sets the "domination_limit" field to the value that was provided on create.
func (u *AppCouponSettingUpsert) UpdateDominationLimit() *AppCouponSettingUpsert {
	u.SetExcluded(appcouponsetting.FieldDominationLimit)
	return u
}

// SetTotalLimit sets the "total_limit" field.
func (u *AppCouponSettingUpsert) SetTotalLimit(v int32) *AppCouponSettingUpsert {
	u.Set(appcouponsetting.FieldTotalLimit, v)
	return u
}

// UpdateTotalLimit sets the "total_limit" field to the value that was provided on create.
func (u *AppCouponSettingUpsert) UpdateTotalLimit() *AppCouponSettingUpsert {
	u.SetExcluded(appcouponsetting.FieldTotalLimit)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppCouponSettingUpsert) SetCreateAt(v uint32) *AppCouponSettingUpsert {
	u.Set(appcouponsetting.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsert) UpdateCreateAt() *AppCouponSettingUpsert {
	u.SetExcluded(appcouponsetting.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *AppCouponSettingUpsert) SetUpdateAt(v uint32) *AppCouponSettingUpsert {
	u.Set(appcouponsetting.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsert) UpdateUpdateAt() *AppCouponSettingUpsert {
	u.SetExcluded(appcouponsetting.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppCouponSettingUpsert) SetDeleteAt(v uint32) *AppCouponSettingUpsert {
	u.Set(appcouponsetting.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsert) UpdateDeleteAt() *AppCouponSettingUpsert {
	u.SetExcluded(appcouponsetting.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppCouponSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appcouponsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppCouponSettingUpsertOne) UpdateNewValues() *AppCouponSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appcouponsetting.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppCouponSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppCouponSettingUpsertOne) Ignore() *AppCouponSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppCouponSettingUpsertOne) DoNothing() *AppCouponSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCouponSettingCreate.OnConflict
// documentation for more info.
func (u *AppCouponSettingUpsertOne) Update(set func(*AppCouponSettingUpsert)) *AppCouponSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppCouponSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppCouponSettingUpsertOne) SetAppID(v uuid.UUID) *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppCouponSettingUpsertOne) UpdateAppID() *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetDominationLimit sets the "domination_limit" field.
func (u *AppCouponSettingUpsertOne) SetDominationLimit(v uint64) *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetDominationLimit(v)
	})
}

// UpdateDominationLimit sets the "domination_limit" field to the value that was provided on create.
func (u *AppCouponSettingUpsertOne) UpdateDominationLimit() *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateDominationLimit()
	})
}

// SetTotalLimit sets the "total_limit" field.
func (u *AppCouponSettingUpsertOne) SetTotalLimit(v int32) *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetTotalLimit(v)
	})
}

// UpdateTotalLimit sets the "total_limit" field to the value that was provided on create.
func (u *AppCouponSettingUpsertOne) UpdateTotalLimit() *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateTotalLimit()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppCouponSettingUpsertOne) SetCreateAt(v uint32) *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsertOne) UpdateCreateAt() *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppCouponSettingUpsertOne) SetUpdateAt(v uint32) *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsertOne) UpdateUpdateAt() *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppCouponSettingUpsertOne) SetDeleteAt(v uint32) *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsertOne) UpdateDeleteAt() *AppCouponSettingUpsertOne {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *AppCouponSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCouponSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppCouponSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppCouponSettingUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppCouponSettingUpsertOne.ID is not supported by MySQL driver. Use AppCouponSettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppCouponSettingUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppCouponSettingCreateBulk is the builder for creating many AppCouponSetting entities in bulk.
type AppCouponSettingCreateBulk struct {
	config
	builders []*AppCouponSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the AppCouponSetting entities in the database.
func (acscb *AppCouponSettingCreateBulk) Save(ctx context.Context) ([]*AppCouponSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acscb.builders))
	nodes := make([]*AppCouponSetting, len(acscb.builders))
	mutators := make([]Mutator, len(acscb.builders))
	for i := range acscb.builders {
		func(i int, root context.Context) {
			builder := acscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppCouponSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, acscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, acscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acscb *AppCouponSettingCreateBulk) SaveX(ctx context.Context) []*AppCouponSetting {
	v, err := acscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acscb *AppCouponSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := acscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acscb *AppCouponSettingCreateBulk) ExecX(ctx context.Context) {
	if err := acscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppCouponSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppCouponSettingUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (acscb *AppCouponSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppCouponSettingUpsertBulk {
	acscb.conflict = opts
	return &AppCouponSettingUpsertBulk{
		create: acscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppCouponSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acscb *AppCouponSettingCreateBulk) OnConflictColumns(columns ...string) *AppCouponSettingUpsertBulk {
	acscb.conflict = append(acscb.conflict, sql.ConflictColumns(columns...))
	return &AppCouponSettingUpsertBulk{
		create: acscb,
	}
}

// AppCouponSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of AppCouponSetting nodes.
type AppCouponSettingUpsertBulk struct {
	create *AppCouponSettingCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppCouponSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appcouponsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppCouponSettingUpsertBulk) UpdateNewValues() *AppCouponSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appcouponsetting.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppCouponSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppCouponSettingUpsertBulk) Ignore() *AppCouponSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppCouponSettingUpsertBulk) DoNothing() *AppCouponSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCouponSettingCreateBulk.OnConflict
// documentation for more info.
func (u *AppCouponSettingUpsertBulk) Update(set func(*AppCouponSettingUpsert)) *AppCouponSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppCouponSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppCouponSettingUpsertBulk) SetAppID(v uuid.UUID) *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppCouponSettingUpsertBulk) UpdateAppID() *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetDominationLimit sets the "domination_limit" field.
func (u *AppCouponSettingUpsertBulk) SetDominationLimit(v uint64) *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetDominationLimit(v)
	})
}

// UpdateDominationLimit sets the "domination_limit" field to the value that was provided on create.
func (u *AppCouponSettingUpsertBulk) UpdateDominationLimit() *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateDominationLimit()
	})
}

// SetTotalLimit sets the "total_limit" field.
func (u *AppCouponSettingUpsertBulk) SetTotalLimit(v int32) *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetTotalLimit(v)
	})
}

// UpdateTotalLimit sets the "total_limit" field to the value that was provided on create.
func (u *AppCouponSettingUpsertBulk) UpdateTotalLimit() *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateTotalLimit()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppCouponSettingUpsertBulk) SetCreateAt(v uint32) *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsertBulk) UpdateCreateAt() *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppCouponSettingUpsertBulk) SetUpdateAt(v uint32) *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsertBulk) UpdateUpdateAt() *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppCouponSettingUpsertBulk) SetDeleteAt(v uint32) *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppCouponSettingUpsertBulk) UpdateDeleteAt() *AppCouponSettingUpsertBulk {
	return u.Update(func(s *AppCouponSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *AppCouponSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppCouponSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCouponSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppCouponSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}