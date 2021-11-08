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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"
	"github.com/google/uuid"
)

// NewUserRewardSettingCreate is the builder for creating a NewUserRewardSetting entity.
type NewUserRewardSettingCreate struct {
	config
	mutation *NewUserRewardSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (nursc *NewUserRewardSettingCreate) SetAppID(u uuid.UUID) *NewUserRewardSettingCreate {
	nursc.mutation.SetAppID(u)
	return nursc
}

// SetRegistrationCouponID sets the "registration_coupon_id" field.
func (nursc *NewUserRewardSettingCreate) SetRegistrationCouponID(u uuid.UUID) *NewUserRewardSettingCreate {
	nursc.mutation.SetRegistrationCouponID(u)
	return nursc
}

// SetKycCouponID sets the "kyc_coupon_id" field.
func (nursc *NewUserRewardSettingCreate) SetKycCouponID(u uuid.UUID) *NewUserRewardSettingCreate {
	nursc.mutation.SetKycCouponID(u)
	return nursc
}

// SetCreateAt sets the "create_at" field.
func (nursc *NewUserRewardSettingCreate) SetCreateAt(u uint32) *NewUserRewardSettingCreate {
	nursc.mutation.SetCreateAt(u)
	return nursc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (nursc *NewUserRewardSettingCreate) SetNillableCreateAt(u *uint32) *NewUserRewardSettingCreate {
	if u != nil {
		nursc.SetCreateAt(*u)
	}
	return nursc
}

// SetUpdateAt sets the "update_at" field.
func (nursc *NewUserRewardSettingCreate) SetUpdateAt(u uint32) *NewUserRewardSettingCreate {
	nursc.mutation.SetUpdateAt(u)
	return nursc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (nursc *NewUserRewardSettingCreate) SetNillableUpdateAt(u *uint32) *NewUserRewardSettingCreate {
	if u != nil {
		nursc.SetUpdateAt(*u)
	}
	return nursc
}

// SetDeleteAt sets the "delete_at" field.
func (nursc *NewUserRewardSettingCreate) SetDeleteAt(u uint32) *NewUserRewardSettingCreate {
	nursc.mutation.SetDeleteAt(u)
	return nursc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (nursc *NewUserRewardSettingCreate) SetNillableDeleteAt(u *uint32) *NewUserRewardSettingCreate {
	if u != nil {
		nursc.SetDeleteAt(*u)
	}
	return nursc
}

// SetID sets the "id" field.
func (nursc *NewUserRewardSettingCreate) SetID(u uuid.UUID) *NewUserRewardSettingCreate {
	nursc.mutation.SetID(u)
	return nursc
}

// Mutation returns the NewUserRewardSettingMutation object of the builder.
func (nursc *NewUserRewardSettingCreate) Mutation() *NewUserRewardSettingMutation {
	return nursc.mutation
}

// Save creates the NewUserRewardSetting in the database.
func (nursc *NewUserRewardSettingCreate) Save(ctx context.Context) (*NewUserRewardSetting, error) {
	var (
		err  error
		node *NewUserRewardSetting
	)
	nursc.defaults()
	if len(nursc.hooks) == 0 {
		if err = nursc.check(); err != nil {
			return nil, err
		}
		node, err = nursc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewUserRewardSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nursc.check(); err != nil {
				return nil, err
			}
			nursc.mutation = mutation
			if node, err = nursc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nursc.hooks) - 1; i >= 0; i-- {
			if nursc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nursc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nursc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nursc *NewUserRewardSettingCreate) SaveX(ctx context.Context) *NewUserRewardSetting {
	v, err := nursc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nursc *NewUserRewardSettingCreate) Exec(ctx context.Context) error {
	_, err := nursc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nursc *NewUserRewardSettingCreate) ExecX(ctx context.Context) {
	if err := nursc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nursc *NewUserRewardSettingCreate) defaults() {
	if _, ok := nursc.mutation.CreateAt(); !ok {
		v := newuserrewardsetting.DefaultCreateAt()
		nursc.mutation.SetCreateAt(v)
	}
	if _, ok := nursc.mutation.UpdateAt(); !ok {
		v := newuserrewardsetting.DefaultUpdateAt()
		nursc.mutation.SetUpdateAt(v)
	}
	if _, ok := nursc.mutation.DeleteAt(); !ok {
		v := newuserrewardsetting.DefaultDeleteAt()
		nursc.mutation.SetDeleteAt(v)
	}
	if _, ok := nursc.mutation.ID(); !ok {
		v := newuserrewardsetting.DefaultID()
		nursc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nursc *NewUserRewardSettingCreate) check() error {
	if _, ok := nursc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := nursc.mutation.RegistrationCouponID(); !ok {
		return &ValidationError{Name: "registration_coupon_id", err: errors.New(`ent: missing required field "registration_coupon_id"`)}
	}
	if _, ok := nursc.mutation.KycCouponID(); !ok {
		return &ValidationError{Name: "kyc_coupon_id", err: errors.New(`ent: missing required field "kyc_coupon_id"`)}
	}
	if _, ok := nursc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := nursc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := nursc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (nursc *NewUserRewardSettingCreate) sqlSave(ctx context.Context) (*NewUserRewardSetting, error) {
	_node, _spec := nursc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nursc.driver, _spec); err != nil {
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

func (nursc *NewUserRewardSettingCreate) createSpec() (*NewUserRewardSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &NewUserRewardSetting{config: nursc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: newuserrewardsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: newuserrewardsetting.FieldID,
			},
		}
	)
	_spec.OnConflict = nursc.conflict
	if id, ok := nursc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nursc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: newuserrewardsetting.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := nursc.mutation.RegistrationCouponID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: newuserrewardsetting.FieldRegistrationCouponID,
		})
		_node.RegistrationCouponID = value
	}
	if value, ok := nursc.mutation.KycCouponID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: newuserrewardsetting.FieldKycCouponID,
		})
		_node.KycCouponID = value
	}
	if value, ok := nursc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: newuserrewardsetting.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := nursc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: newuserrewardsetting.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := nursc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: newuserrewardsetting.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NewUserRewardSetting.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NewUserRewardSettingUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (nursc *NewUserRewardSettingCreate) OnConflict(opts ...sql.ConflictOption) *NewUserRewardSettingUpsertOne {
	nursc.conflict = opts
	return &NewUserRewardSettingUpsertOne{
		create: nursc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (nursc *NewUserRewardSettingCreate) OnConflictColumns(columns ...string) *NewUserRewardSettingUpsertOne {
	nursc.conflict = append(nursc.conflict, sql.ConflictColumns(columns...))
	return &NewUserRewardSettingUpsertOne{
		create: nursc,
	}
}

type (
	// NewUserRewardSettingUpsertOne is the builder for "upsert"-ing
	//  one NewUserRewardSetting node.
	NewUserRewardSettingUpsertOne struct {
		create *NewUserRewardSettingCreate
	}

	// NewUserRewardSettingUpsert is the "OnConflict" setter.
	NewUserRewardSettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *NewUserRewardSettingUpsert) SetAppID(v uuid.UUID) *NewUserRewardSettingUpsert {
	u.Set(newuserrewardsetting.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsert) UpdateAppID() *NewUserRewardSettingUpsert {
	u.SetExcluded(newuserrewardsetting.FieldAppID)
	return u
}

// SetRegistrationCouponID sets the "registration_coupon_id" field.
func (u *NewUserRewardSettingUpsert) SetRegistrationCouponID(v uuid.UUID) *NewUserRewardSettingUpsert {
	u.Set(newuserrewardsetting.FieldRegistrationCouponID, v)
	return u
}

// UpdateRegistrationCouponID sets the "registration_coupon_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsert) UpdateRegistrationCouponID() *NewUserRewardSettingUpsert {
	u.SetExcluded(newuserrewardsetting.FieldRegistrationCouponID)
	return u
}

// SetKycCouponID sets the "kyc_coupon_id" field.
func (u *NewUserRewardSettingUpsert) SetKycCouponID(v uuid.UUID) *NewUserRewardSettingUpsert {
	u.Set(newuserrewardsetting.FieldKycCouponID, v)
	return u
}

// UpdateKycCouponID sets the "kyc_coupon_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsert) UpdateKycCouponID() *NewUserRewardSettingUpsert {
	u.SetExcluded(newuserrewardsetting.FieldKycCouponID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *NewUserRewardSettingUpsert) SetCreateAt(v uint32) *NewUserRewardSettingUpsert {
	u.Set(newuserrewardsetting.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsert) UpdateCreateAt() *NewUserRewardSettingUpsert {
	u.SetExcluded(newuserrewardsetting.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *NewUserRewardSettingUpsert) SetUpdateAt(v uint32) *NewUserRewardSettingUpsert {
	u.Set(newuserrewardsetting.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsert) UpdateUpdateAt() *NewUserRewardSettingUpsert {
	u.SetExcluded(newuserrewardsetting.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *NewUserRewardSettingUpsert) SetDeleteAt(v uint32) *NewUserRewardSettingUpsert {
	u.Set(newuserrewardsetting.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsert) UpdateDeleteAt() *NewUserRewardSettingUpsert {
	u.SetExcluded(newuserrewardsetting.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(newuserrewardsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NewUserRewardSettingUpsertOne) UpdateNewValues() *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(newuserrewardsetting.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.NewUserRewardSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *NewUserRewardSettingUpsertOne) Ignore() *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NewUserRewardSettingUpsertOne) DoNothing() *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NewUserRewardSettingCreate.OnConflict
// documentation for more info.
func (u *NewUserRewardSettingUpsertOne) Update(set func(*NewUserRewardSettingUpsert)) *NewUserRewardSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NewUserRewardSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *NewUserRewardSettingUpsertOne) SetAppID(v uuid.UUID) *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertOne) UpdateAppID() *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetRegistrationCouponID sets the "registration_coupon_id" field.
func (u *NewUserRewardSettingUpsertOne) SetRegistrationCouponID(v uuid.UUID) *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetRegistrationCouponID(v)
	})
}

// UpdateRegistrationCouponID sets the "registration_coupon_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertOne) UpdateRegistrationCouponID() *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateRegistrationCouponID()
	})
}

// SetKycCouponID sets the "kyc_coupon_id" field.
func (u *NewUserRewardSettingUpsertOne) SetKycCouponID(v uuid.UUID) *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetKycCouponID(v)
	})
}

// UpdateKycCouponID sets the "kyc_coupon_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertOne) UpdateKycCouponID() *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateKycCouponID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *NewUserRewardSettingUpsertOne) SetCreateAt(v uint32) *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertOne) UpdateCreateAt() *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *NewUserRewardSettingUpsertOne) SetUpdateAt(v uint32) *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertOne) UpdateUpdateAt() *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *NewUserRewardSettingUpsertOne) SetDeleteAt(v uint32) *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertOne) UpdateDeleteAt() *NewUserRewardSettingUpsertOne {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *NewUserRewardSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NewUserRewardSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NewUserRewardSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NewUserRewardSettingUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NewUserRewardSettingUpsertOne.ID is not supported by MySQL driver. Use NewUserRewardSettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NewUserRewardSettingUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NewUserRewardSettingCreateBulk is the builder for creating many NewUserRewardSetting entities in bulk.
type NewUserRewardSettingCreateBulk struct {
	config
	builders []*NewUserRewardSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the NewUserRewardSetting entities in the database.
func (nurscb *NewUserRewardSettingCreateBulk) Save(ctx context.Context) ([]*NewUserRewardSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nurscb.builders))
	nodes := make([]*NewUserRewardSetting, len(nurscb.builders))
	mutators := make([]Mutator, len(nurscb.builders))
	for i := range nurscb.builders {
		func(i int, root context.Context) {
			builder := nurscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NewUserRewardSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, nurscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nurscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nurscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nurscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nurscb *NewUserRewardSettingCreateBulk) SaveX(ctx context.Context) []*NewUserRewardSetting {
	v, err := nurscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nurscb *NewUserRewardSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := nurscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nurscb *NewUserRewardSettingCreateBulk) ExecX(ctx context.Context) {
	if err := nurscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NewUserRewardSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NewUserRewardSettingUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (nurscb *NewUserRewardSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *NewUserRewardSettingUpsertBulk {
	nurscb.conflict = opts
	return &NewUserRewardSettingUpsertBulk{
		create: nurscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (nurscb *NewUserRewardSettingCreateBulk) OnConflictColumns(columns ...string) *NewUserRewardSettingUpsertBulk {
	nurscb.conflict = append(nurscb.conflict, sql.ConflictColumns(columns...))
	return &NewUserRewardSettingUpsertBulk{
		create: nurscb,
	}
}

// NewUserRewardSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of NewUserRewardSetting nodes.
type NewUserRewardSettingUpsertBulk struct {
	create *NewUserRewardSettingCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(newuserrewardsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *NewUserRewardSettingUpsertBulk) UpdateNewValues() *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(newuserrewardsetting.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NewUserRewardSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *NewUserRewardSettingUpsertBulk) Ignore() *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NewUserRewardSettingUpsertBulk) DoNothing() *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NewUserRewardSettingCreateBulk.OnConflict
// documentation for more info.
func (u *NewUserRewardSettingUpsertBulk) Update(set func(*NewUserRewardSettingUpsert)) *NewUserRewardSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NewUserRewardSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *NewUserRewardSettingUpsertBulk) SetAppID(v uuid.UUID) *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertBulk) UpdateAppID() *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetRegistrationCouponID sets the "registration_coupon_id" field.
func (u *NewUserRewardSettingUpsertBulk) SetRegistrationCouponID(v uuid.UUID) *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetRegistrationCouponID(v)
	})
}

// UpdateRegistrationCouponID sets the "registration_coupon_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertBulk) UpdateRegistrationCouponID() *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateRegistrationCouponID()
	})
}

// SetKycCouponID sets the "kyc_coupon_id" field.
func (u *NewUserRewardSettingUpsertBulk) SetKycCouponID(v uuid.UUID) *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetKycCouponID(v)
	})
}

// UpdateKycCouponID sets the "kyc_coupon_id" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertBulk) UpdateKycCouponID() *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateKycCouponID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *NewUserRewardSettingUpsertBulk) SetCreateAt(v uint32) *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertBulk) UpdateCreateAt() *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *NewUserRewardSettingUpsertBulk) SetUpdateAt(v uint32) *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertBulk) UpdateUpdateAt() *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *NewUserRewardSettingUpsertBulk) SetDeleteAt(v uint32) *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *NewUserRewardSettingUpsertBulk) UpdateDeleteAt() *NewUserRewardSettingUpsertBulk {
	return u.Update(func(s *NewUserRewardSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *NewUserRewardSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NewUserRewardSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NewUserRewardSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NewUserRewardSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
