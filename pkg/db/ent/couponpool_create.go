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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/google/uuid"
)

// CouponPoolCreate is the builder for creating a CouponPool entity.
type CouponPoolCreate struct {
	config
	mutation *CouponPoolMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDenomination sets the "denomination" field.
func (cpc *CouponPoolCreate) SetDenomination(u uint64) *CouponPoolCreate {
	cpc.mutation.SetDenomination(u)
	return cpc
}

// SetCirculation sets the "circulation" field.
func (cpc *CouponPoolCreate) SetCirculation(i int32) *CouponPoolCreate {
	cpc.mutation.SetCirculation(i)
	return cpc
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (cpc *CouponPoolCreate) SetReleaseByUserID(u uuid.UUID) *CouponPoolCreate {
	cpc.mutation.SetReleaseByUserID(u)
	return cpc
}

// SetStart sets the "start" field.
func (cpc *CouponPoolCreate) SetStart(u uint32) *CouponPoolCreate {
	cpc.mutation.SetStart(u)
	return cpc
}

// SetDurationDays sets the "duration_days" field.
func (cpc *CouponPoolCreate) SetDurationDays(i int32) *CouponPoolCreate {
	cpc.mutation.SetDurationDays(i)
	return cpc
}

// SetAppID sets the "app_id" field.
func (cpc *CouponPoolCreate) SetAppID(u uuid.UUID) *CouponPoolCreate {
	cpc.mutation.SetAppID(u)
	return cpc
}

// SetMessage sets the "message" field.
func (cpc *CouponPoolCreate) SetMessage(s string) *CouponPoolCreate {
	cpc.mutation.SetMessage(s)
	return cpc
}

// SetName sets the "name" field.
func (cpc *CouponPoolCreate) SetName(s string) *CouponPoolCreate {
	cpc.mutation.SetName(s)
	return cpc
}

// SetCreateAt sets the "create_at" field.
func (cpc *CouponPoolCreate) SetCreateAt(u uint32) *CouponPoolCreate {
	cpc.mutation.SetCreateAt(u)
	return cpc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cpc *CouponPoolCreate) SetNillableCreateAt(u *uint32) *CouponPoolCreate {
	if u != nil {
		cpc.SetCreateAt(*u)
	}
	return cpc
}

// SetUpdateAt sets the "update_at" field.
func (cpc *CouponPoolCreate) SetUpdateAt(u uint32) *CouponPoolCreate {
	cpc.mutation.SetUpdateAt(u)
	return cpc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (cpc *CouponPoolCreate) SetNillableUpdateAt(u *uint32) *CouponPoolCreate {
	if u != nil {
		cpc.SetUpdateAt(*u)
	}
	return cpc
}

// SetDeleteAt sets the "delete_at" field.
func (cpc *CouponPoolCreate) SetDeleteAt(u uint32) *CouponPoolCreate {
	cpc.mutation.SetDeleteAt(u)
	return cpc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cpc *CouponPoolCreate) SetNillableDeleteAt(u *uint32) *CouponPoolCreate {
	if u != nil {
		cpc.SetDeleteAt(*u)
	}
	return cpc
}

// SetID sets the "id" field.
func (cpc *CouponPoolCreate) SetID(u uuid.UUID) *CouponPoolCreate {
	cpc.mutation.SetID(u)
	return cpc
}

// Mutation returns the CouponPoolMutation object of the builder.
func (cpc *CouponPoolCreate) Mutation() *CouponPoolMutation {
	return cpc.mutation
}

// Save creates the CouponPool in the database.
func (cpc *CouponPoolCreate) Save(ctx context.Context) (*CouponPool, error) {
	var (
		err  error
		node *CouponPool
	)
	cpc.defaults()
	if len(cpc.hooks) == 0 {
		if err = cpc.check(); err != nil {
			return nil, err
		}
		node, err = cpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponPoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpc.check(); err != nil {
				return nil, err
			}
			cpc.mutation = mutation
			if node, err = cpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cpc.hooks) - 1; i >= 0; i-- {
			if cpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *CouponPoolCreate) SaveX(ctx context.Context) *CouponPool {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpc *CouponPoolCreate) Exec(ctx context.Context) error {
	_, err := cpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpc *CouponPoolCreate) ExecX(ctx context.Context) {
	if err := cpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpc *CouponPoolCreate) defaults() {
	if _, ok := cpc.mutation.CreateAt(); !ok {
		v := couponpool.DefaultCreateAt()
		cpc.mutation.SetCreateAt(v)
	}
	if _, ok := cpc.mutation.UpdateAt(); !ok {
		v := couponpool.DefaultUpdateAt()
		cpc.mutation.SetUpdateAt(v)
	}
	if _, ok := cpc.mutation.DeleteAt(); !ok {
		v := couponpool.DefaultDeleteAt()
		cpc.mutation.SetDeleteAt(v)
	}
	if _, ok := cpc.mutation.ID(); !ok {
		v := couponpool.DefaultID()
		cpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpc *CouponPoolCreate) check() error {
	if _, ok := cpc.mutation.Denomination(); !ok {
		return &ValidationError{Name: "denomination", err: errors.New(`ent: missing required field "denomination"`)}
	}
	if _, ok := cpc.mutation.Circulation(); !ok {
		return &ValidationError{Name: "circulation", err: errors.New(`ent: missing required field "circulation"`)}
	}
	if _, ok := cpc.mutation.ReleaseByUserID(); !ok {
		return &ValidationError{Name: "release_by_user_id", err: errors.New(`ent: missing required field "release_by_user_id"`)}
	}
	if _, ok := cpc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "start"`)}
	}
	if _, ok := cpc.mutation.DurationDays(); !ok {
		return &ValidationError{Name: "duration_days", err: errors.New(`ent: missing required field "duration_days"`)}
	}
	if _, ok := cpc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := cpc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if v, ok := cpc.mutation.Message(); ok {
		if err := couponpool.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "message": %w`, err)}
		}
	}
	if _, ok := cpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cpc.mutation.Name(); ok {
		if err := couponpool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cpc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := cpc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := cpc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (cpc *CouponPoolCreate) sqlSave(ctx context.Context) (*CouponPool, error) {
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
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

func (cpc *CouponPoolCreate) createSpec() (*CouponPool, *sqlgraph.CreateSpec) {
	var (
		_node = &CouponPool{config: cpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: couponpool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponpool.FieldID,
			},
		}
	)
	_spec.OnConflict = cpc.conflict
	if id, ok := cpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cpc.mutation.Denomination(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: couponpool.FieldDenomination,
		})
		_node.Denomination = value
	}
	if value, ok := cpc.mutation.Circulation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldCirculation,
		})
		_node.Circulation = value
	}
	if value, ok := cpc.mutation.ReleaseByUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponpool.FieldReleaseByUserID,
		})
		_node.ReleaseByUserID = value
	}
	if value, ok := cpc.mutation.Start(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldStart,
		})
		_node.Start = value
	}
	if value, ok := cpc.mutation.DurationDays(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: couponpool.FieldDurationDays,
		})
		_node.DurationDays = value
	}
	if value, ok := cpc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponpool.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := cpc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponpool.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := cpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponpool.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cpc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := cpc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := cpc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponpool.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CouponPool.Create().
//		SetDenomination(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CouponPoolUpsert) {
//			SetDenomination(v+v).
//		}).
//		Exec(ctx)
//
func (cpc *CouponPoolCreate) OnConflict(opts ...sql.ConflictOption) *CouponPoolUpsertOne {
	cpc.conflict = opts
	return &CouponPoolUpsertOne{
		create: cpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CouponPool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cpc *CouponPoolCreate) OnConflictColumns(columns ...string) *CouponPoolUpsertOne {
	cpc.conflict = append(cpc.conflict, sql.ConflictColumns(columns...))
	return &CouponPoolUpsertOne{
		create: cpc,
	}
}

type (
	// CouponPoolUpsertOne is the builder for "upsert"-ing
	//  one CouponPool node.
	CouponPoolUpsertOne struct {
		create *CouponPoolCreate
	}

	// CouponPoolUpsert is the "OnConflict" setter.
	CouponPoolUpsert struct {
		*sql.UpdateSet
	}
)

// SetDenomination sets the "denomination" field.
func (u *CouponPoolUpsert) SetDenomination(v uint64) *CouponPoolUpsert {
	u.Set(couponpool.FieldDenomination, v)
	return u
}

// UpdateDenomination sets the "denomination" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateDenomination() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldDenomination)
	return u
}

// SetCirculation sets the "circulation" field.
func (u *CouponPoolUpsert) SetCirculation(v int32) *CouponPoolUpsert {
	u.Set(couponpool.FieldCirculation, v)
	return u
}

// UpdateCirculation sets the "circulation" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateCirculation() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldCirculation)
	return u
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (u *CouponPoolUpsert) SetReleaseByUserID(v uuid.UUID) *CouponPoolUpsert {
	u.Set(couponpool.FieldReleaseByUserID, v)
	return u
}

// UpdateReleaseByUserID sets the "release_by_user_id" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateReleaseByUserID() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldReleaseByUserID)
	return u
}

// SetStart sets the "start" field.
func (u *CouponPoolUpsert) SetStart(v uint32) *CouponPoolUpsert {
	u.Set(couponpool.FieldStart, v)
	return u
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateStart() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldStart)
	return u
}

// SetDurationDays sets the "duration_days" field.
func (u *CouponPoolUpsert) SetDurationDays(v int32) *CouponPoolUpsert {
	u.Set(couponpool.FieldDurationDays, v)
	return u
}

// UpdateDurationDays sets the "duration_days" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateDurationDays() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldDurationDays)
	return u
}

// SetAppID sets the "app_id" field.
func (u *CouponPoolUpsert) SetAppID(v uuid.UUID) *CouponPoolUpsert {
	u.Set(couponpool.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateAppID() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldAppID)
	return u
}

// SetMessage sets the "message" field.
func (u *CouponPoolUpsert) SetMessage(v string) *CouponPoolUpsert {
	u.Set(couponpool.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateMessage() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldMessage)
	return u
}

// SetName sets the "name" field.
func (u *CouponPoolUpsert) SetName(v string) *CouponPoolUpsert {
	u.Set(couponpool.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateName() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldName)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *CouponPoolUpsert) SetCreateAt(v uint32) *CouponPoolUpsert {
	u.Set(couponpool.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateCreateAt() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *CouponPoolUpsert) SetUpdateAt(v uint32) *CouponPoolUpsert {
	u.Set(couponpool.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateUpdateAt() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *CouponPoolUpsert) SetDeleteAt(v uint32) *CouponPoolUpsert {
	u.Set(couponpool.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CouponPoolUpsert) UpdateDeleteAt() *CouponPoolUpsert {
	u.SetExcluded(couponpool.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CouponPool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(couponpool.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CouponPoolUpsertOne) UpdateNewValues() *CouponPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(couponpool.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CouponPool.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CouponPoolUpsertOne) Ignore() *CouponPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CouponPoolUpsertOne) DoNothing() *CouponPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CouponPoolCreate.OnConflict
// documentation for more info.
func (u *CouponPoolUpsertOne) Update(set func(*CouponPoolUpsert)) *CouponPoolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CouponPoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetDenomination sets the "denomination" field.
func (u *CouponPoolUpsertOne) SetDenomination(v uint64) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetDenomination(v)
	})
}

// UpdateDenomination sets the "denomination" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateDenomination() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateDenomination()
	})
}

// SetCirculation sets the "circulation" field.
func (u *CouponPoolUpsertOne) SetCirculation(v int32) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetCirculation(v)
	})
}

// UpdateCirculation sets the "circulation" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateCirculation() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateCirculation()
	})
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (u *CouponPoolUpsertOne) SetReleaseByUserID(v uuid.UUID) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetReleaseByUserID(v)
	})
}

// UpdateReleaseByUserID sets the "release_by_user_id" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateReleaseByUserID() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateReleaseByUserID()
	})
}

// SetStart sets the "start" field.
func (u *CouponPoolUpsertOne) SetStart(v uint32) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateStart() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateStart()
	})
}

// SetDurationDays sets the "duration_days" field.
func (u *CouponPoolUpsertOne) SetDurationDays(v int32) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetDurationDays(v)
	})
}

// UpdateDurationDays sets the "duration_days" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateDurationDays() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateDurationDays()
	})
}

// SetAppID sets the "app_id" field.
func (u *CouponPoolUpsertOne) SetAppID(v uuid.UUID) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateAppID() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateAppID()
	})
}

// SetMessage sets the "message" field.
func (u *CouponPoolUpsertOne) SetMessage(v string) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateMessage() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateMessage()
	})
}

// SetName sets the "name" field.
func (u *CouponPoolUpsertOne) SetName(v string) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateName() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateName()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CouponPoolUpsertOne) SetCreateAt(v uint32) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateCreateAt() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CouponPoolUpsertOne) SetUpdateAt(v uint32) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateUpdateAt() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CouponPoolUpsertOne) SetDeleteAt(v uint32) *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CouponPoolUpsertOne) UpdateDeleteAt() *CouponPoolUpsertOne {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CouponPoolUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CouponPoolCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CouponPoolUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CouponPoolUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CouponPoolUpsertOne.ID is not supported by MySQL driver. Use CouponPoolUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CouponPoolUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CouponPoolCreateBulk is the builder for creating many CouponPool entities in bulk.
type CouponPoolCreateBulk struct {
	config
	builders []*CouponPoolCreate
	conflict []sql.ConflictOption
}

// Save creates the CouponPool entities in the database.
func (cpcb *CouponPoolCreateBulk) Save(ctx context.Context) ([]*CouponPool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*CouponPool, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CouponPoolMutation)
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
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *CouponPoolCreateBulk) SaveX(ctx context.Context) []*CouponPool {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpcb *CouponPoolCreateBulk) Exec(ctx context.Context) error {
	_, err := cpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpcb *CouponPoolCreateBulk) ExecX(ctx context.Context) {
	if err := cpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CouponPool.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CouponPoolUpsert) {
//			SetDenomination(v+v).
//		}).
//		Exec(ctx)
//
func (cpcb *CouponPoolCreateBulk) OnConflict(opts ...sql.ConflictOption) *CouponPoolUpsertBulk {
	cpcb.conflict = opts
	return &CouponPoolUpsertBulk{
		create: cpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CouponPool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cpcb *CouponPoolCreateBulk) OnConflictColumns(columns ...string) *CouponPoolUpsertBulk {
	cpcb.conflict = append(cpcb.conflict, sql.ConflictColumns(columns...))
	return &CouponPoolUpsertBulk{
		create: cpcb,
	}
}

// CouponPoolUpsertBulk is the builder for "upsert"-ing
// a bulk of CouponPool nodes.
type CouponPoolUpsertBulk struct {
	create *CouponPoolCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CouponPool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(couponpool.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CouponPoolUpsertBulk) UpdateNewValues() *CouponPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(couponpool.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CouponPool.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CouponPoolUpsertBulk) Ignore() *CouponPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CouponPoolUpsertBulk) DoNothing() *CouponPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CouponPoolCreateBulk.OnConflict
// documentation for more info.
func (u *CouponPoolUpsertBulk) Update(set func(*CouponPoolUpsert)) *CouponPoolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CouponPoolUpsert{UpdateSet: update})
	}))
	return u
}

// SetDenomination sets the "denomination" field.
func (u *CouponPoolUpsertBulk) SetDenomination(v uint64) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetDenomination(v)
	})
}

// UpdateDenomination sets the "denomination" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateDenomination() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateDenomination()
	})
}

// SetCirculation sets the "circulation" field.
func (u *CouponPoolUpsertBulk) SetCirculation(v int32) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetCirculation(v)
	})
}

// UpdateCirculation sets the "circulation" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateCirculation() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateCirculation()
	})
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (u *CouponPoolUpsertBulk) SetReleaseByUserID(v uuid.UUID) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetReleaseByUserID(v)
	})
}

// UpdateReleaseByUserID sets the "release_by_user_id" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateReleaseByUserID() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateReleaseByUserID()
	})
}

// SetStart sets the "start" field.
func (u *CouponPoolUpsertBulk) SetStart(v uint32) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateStart() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateStart()
	})
}

// SetDurationDays sets the "duration_days" field.
func (u *CouponPoolUpsertBulk) SetDurationDays(v int32) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetDurationDays(v)
	})
}

// UpdateDurationDays sets the "duration_days" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateDurationDays() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateDurationDays()
	})
}

// SetAppID sets the "app_id" field.
func (u *CouponPoolUpsertBulk) SetAppID(v uuid.UUID) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateAppID() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateAppID()
	})
}

// SetMessage sets the "message" field.
func (u *CouponPoolUpsertBulk) SetMessage(v string) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateMessage() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateMessage()
	})
}

// SetName sets the "name" field.
func (u *CouponPoolUpsertBulk) SetName(v string) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateName() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateName()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CouponPoolUpsertBulk) SetCreateAt(v uint32) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateCreateAt() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CouponPoolUpsertBulk) SetUpdateAt(v uint32) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateUpdateAt() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CouponPoolUpsertBulk) SetDeleteAt(v uint32) *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CouponPoolUpsertBulk) UpdateDeleteAt() *CouponPoolUpsertBulk {
	return u.Update(func(s *CouponPoolUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CouponPoolUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CouponPoolCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CouponPoolCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CouponPoolUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
