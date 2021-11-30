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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userspecialreduction"
	"github.com/google/uuid"
)

// UserSpecialReductionCreate is the builder for creating a UserSpecialReduction entity.
type UserSpecialReductionCreate struct {
	config
	mutation *UserSpecialReductionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (usrc *UserSpecialReductionCreate) SetAppID(u uuid.UUID) *UserSpecialReductionCreate {
	usrc.mutation.SetAppID(u)
	return usrc
}

// SetUserID sets the "user_id" field.
func (usrc *UserSpecialReductionCreate) SetUserID(u uuid.UUID) *UserSpecialReductionCreate {
	usrc.mutation.SetUserID(u)
	return usrc
}

// SetAmount sets the "amount" field.
func (usrc *UserSpecialReductionCreate) SetAmount(u uint64) *UserSpecialReductionCreate {
	usrc.mutation.SetAmount(u)
	return usrc
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (usrc *UserSpecialReductionCreate) SetReleaseByUserID(u uuid.UUID) *UserSpecialReductionCreate {
	usrc.mutation.SetReleaseByUserID(u)
	return usrc
}

// SetStart sets the "start" field.
func (usrc *UserSpecialReductionCreate) SetStart(u uint32) *UserSpecialReductionCreate {
	usrc.mutation.SetStart(u)
	return usrc
}

// SetDurationDays sets the "duration_days" field.
func (usrc *UserSpecialReductionCreate) SetDurationDays(i int32) *UserSpecialReductionCreate {
	usrc.mutation.SetDurationDays(i)
	return usrc
}

// SetMessage sets the "message" field.
func (usrc *UserSpecialReductionCreate) SetMessage(s string) *UserSpecialReductionCreate {
	usrc.mutation.SetMessage(s)
	return usrc
}

// SetCreateAt sets the "create_at" field.
func (usrc *UserSpecialReductionCreate) SetCreateAt(u uint32) *UserSpecialReductionCreate {
	usrc.mutation.SetCreateAt(u)
	return usrc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (usrc *UserSpecialReductionCreate) SetNillableCreateAt(u *uint32) *UserSpecialReductionCreate {
	if u != nil {
		usrc.SetCreateAt(*u)
	}
	return usrc
}

// SetUpdateAt sets the "update_at" field.
func (usrc *UserSpecialReductionCreate) SetUpdateAt(u uint32) *UserSpecialReductionCreate {
	usrc.mutation.SetUpdateAt(u)
	return usrc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (usrc *UserSpecialReductionCreate) SetNillableUpdateAt(u *uint32) *UserSpecialReductionCreate {
	if u != nil {
		usrc.SetUpdateAt(*u)
	}
	return usrc
}

// SetDeleteAt sets the "delete_at" field.
func (usrc *UserSpecialReductionCreate) SetDeleteAt(u uint32) *UserSpecialReductionCreate {
	usrc.mutation.SetDeleteAt(u)
	return usrc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (usrc *UserSpecialReductionCreate) SetNillableDeleteAt(u *uint32) *UserSpecialReductionCreate {
	if u != nil {
		usrc.SetDeleteAt(*u)
	}
	return usrc
}

// SetID sets the "id" field.
func (usrc *UserSpecialReductionCreate) SetID(u uuid.UUID) *UserSpecialReductionCreate {
	usrc.mutation.SetID(u)
	return usrc
}

// Mutation returns the UserSpecialReductionMutation object of the builder.
func (usrc *UserSpecialReductionCreate) Mutation() *UserSpecialReductionMutation {
	return usrc.mutation
}

// Save creates the UserSpecialReduction in the database.
func (usrc *UserSpecialReductionCreate) Save(ctx context.Context) (*UserSpecialReduction, error) {
	var (
		err  error
		node *UserSpecialReduction
	)
	usrc.defaults()
	if len(usrc.hooks) == 0 {
		if err = usrc.check(); err != nil {
			return nil, err
		}
		node, err = usrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSpecialReductionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usrc.check(); err != nil {
				return nil, err
			}
			usrc.mutation = mutation
			if node, err = usrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(usrc.hooks) - 1; i >= 0; i-- {
			if usrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = usrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (usrc *UserSpecialReductionCreate) SaveX(ctx context.Context) *UserSpecialReduction {
	v, err := usrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usrc *UserSpecialReductionCreate) Exec(ctx context.Context) error {
	_, err := usrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usrc *UserSpecialReductionCreate) ExecX(ctx context.Context) {
	if err := usrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usrc *UserSpecialReductionCreate) defaults() {
	if _, ok := usrc.mutation.CreateAt(); !ok {
		v := userspecialreduction.DefaultCreateAt()
		usrc.mutation.SetCreateAt(v)
	}
	if _, ok := usrc.mutation.UpdateAt(); !ok {
		v := userspecialreduction.DefaultUpdateAt()
		usrc.mutation.SetUpdateAt(v)
	}
	if _, ok := usrc.mutation.DeleteAt(); !ok {
		v := userspecialreduction.DefaultDeleteAt()
		usrc.mutation.SetDeleteAt(v)
	}
	if _, ok := usrc.mutation.ID(); !ok {
		v := userspecialreduction.DefaultID()
		usrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usrc *UserSpecialReductionCreate) check() error {
	if _, ok := usrc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := usrc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := usrc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if _, ok := usrc.mutation.ReleaseByUserID(); !ok {
		return &ValidationError{Name: "release_by_user_id", err: errors.New(`ent: missing required field "release_by_user_id"`)}
	}
	if _, ok := usrc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "start"`)}
	}
	if _, ok := usrc.mutation.DurationDays(); !ok {
		return &ValidationError{Name: "duration_days", err: errors.New(`ent: missing required field "duration_days"`)}
	}
	if _, ok := usrc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if v, ok := usrc.mutation.Message(); ok {
		if err := userspecialreduction.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "message": %w`, err)}
		}
	}
	if _, ok := usrc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := usrc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := usrc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (usrc *UserSpecialReductionCreate) sqlSave(ctx context.Context) (*UserSpecialReduction, error) {
	_node, _spec := usrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usrc.driver, _spec); err != nil {
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

func (usrc *UserSpecialReductionCreate) createSpec() (*UserSpecialReduction, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSpecialReduction{config: usrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userspecialreduction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userspecialreduction.FieldID,
			},
		}
	)
	_spec.OnConflict = usrc.conflict
	if id, ok := usrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := usrc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := usrc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := usrc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userspecialreduction.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := usrc.mutation.ReleaseByUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userspecialreduction.FieldReleaseByUserID,
		})
		_node.ReleaseByUserID = value
	}
	if value, ok := usrc.mutation.Start(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldStart,
		})
		_node.Start = value
	}
	if value, ok := usrc.mutation.DurationDays(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userspecialreduction.FieldDurationDays,
		})
		_node.DurationDays = value
	}
	if value, ok := usrc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userspecialreduction.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := usrc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := usrc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := usrc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userspecialreduction.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserSpecialReduction.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserSpecialReductionUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (usrc *UserSpecialReductionCreate) OnConflict(opts ...sql.ConflictOption) *UserSpecialReductionUpsertOne {
	usrc.conflict = opts
	return &UserSpecialReductionUpsertOne{
		create: usrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserSpecialReduction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (usrc *UserSpecialReductionCreate) OnConflictColumns(columns ...string) *UserSpecialReductionUpsertOne {
	usrc.conflict = append(usrc.conflict, sql.ConflictColumns(columns...))
	return &UserSpecialReductionUpsertOne{
		create: usrc,
	}
}

type (
	// UserSpecialReductionUpsertOne is the builder for "upsert"-ing
	//  one UserSpecialReduction node.
	UserSpecialReductionUpsertOne struct {
		create *UserSpecialReductionCreate
	}

	// UserSpecialReductionUpsert is the "OnConflict" setter.
	UserSpecialReductionUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *UserSpecialReductionUpsert) SetAppID(v uuid.UUID) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateAppID() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserSpecialReductionUpsert) SetUserID(v uuid.UUID) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateUserID() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldUserID)
	return u
}

// SetAmount sets the "amount" field.
func (u *UserSpecialReductionUpsert) SetAmount(v uint64) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateAmount() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldAmount)
	return u
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (u *UserSpecialReductionUpsert) SetReleaseByUserID(v uuid.UUID) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldReleaseByUserID, v)
	return u
}

// UpdateReleaseByUserID sets the "release_by_user_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateReleaseByUserID() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldReleaseByUserID)
	return u
}

// SetStart sets the "start" field.
func (u *UserSpecialReductionUpsert) SetStart(v uint32) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldStart, v)
	return u
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateStart() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldStart)
	return u
}

// SetDurationDays sets the "duration_days" field.
func (u *UserSpecialReductionUpsert) SetDurationDays(v int32) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldDurationDays, v)
	return u
}

// UpdateDurationDays sets the "duration_days" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateDurationDays() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldDurationDays)
	return u
}

// SetMessage sets the "message" field.
func (u *UserSpecialReductionUpsert) SetMessage(v string) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateMessage() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldMessage)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *UserSpecialReductionUpsert) SetCreateAt(v uint32) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateCreateAt() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *UserSpecialReductionUpsert) SetUpdateAt(v uint32) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateUpdateAt() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserSpecialReductionUpsert) SetDeleteAt(v uint32) *UserSpecialReductionUpsert {
	u.Set(userspecialreduction.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsert) UpdateDeleteAt() *UserSpecialReductionUpsert {
	u.SetExcluded(userspecialreduction.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserSpecialReduction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userspecialreduction.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserSpecialReductionUpsertOne) UpdateNewValues() *UserSpecialReductionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userspecialreduction.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserSpecialReduction.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserSpecialReductionUpsertOne) Ignore() *UserSpecialReductionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserSpecialReductionUpsertOne) DoNothing() *UserSpecialReductionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserSpecialReductionCreate.OnConflict
// documentation for more info.
func (u *UserSpecialReductionUpsertOne) Update(set func(*UserSpecialReductionUpsert)) *UserSpecialReductionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserSpecialReductionUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserSpecialReductionUpsertOne) SetAppID(v uuid.UUID) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateAppID() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserSpecialReductionUpsertOne) SetUserID(v uuid.UUID) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateUserID() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateUserID()
	})
}

// SetAmount sets the "amount" field.
func (u *UserSpecialReductionUpsertOne) SetAmount(v uint64) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateAmount() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateAmount()
	})
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (u *UserSpecialReductionUpsertOne) SetReleaseByUserID(v uuid.UUID) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetReleaseByUserID(v)
	})
}

// UpdateReleaseByUserID sets the "release_by_user_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateReleaseByUserID() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateReleaseByUserID()
	})
}

// SetStart sets the "start" field.
func (u *UserSpecialReductionUpsertOne) SetStart(v uint32) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateStart() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateStart()
	})
}

// SetDurationDays sets the "duration_days" field.
func (u *UserSpecialReductionUpsertOne) SetDurationDays(v int32) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetDurationDays(v)
	})
}

// UpdateDurationDays sets the "duration_days" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateDurationDays() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateDurationDays()
	})
}

// SetMessage sets the "message" field.
func (u *UserSpecialReductionUpsertOne) SetMessage(v string) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateMessage() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserSpecialReductionUpsertOne) SetCreateAt(v uint32) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateCreateAt() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserSpecialReductionUpsertOne) SetUpdateAt(v uint32) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateUpdateAt() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserSpecialReductionUpsertOne) SetDeleteAt(v uint32) *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertOne) UpdateDeleteAt() *UserSpecialReductionUpsertOne {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserSpecialReductionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserSpecialReductionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserSpecialReductionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserSpecialReductionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserSpecialReductionUpsertOne.ID is not supported by MySQL driver. Use UserSpecialReductionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserSpecialReductionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserSpecialReductionCreateBulk is the builder for creating many UserSpecialReduction entities in bulk.
type UserSpecialReductionCreateBulk struct {
	config
	builders []*UserSpecialReductionCreate
	conflict []sql.ConflictOption
}

// Save creates the UserSpecialReduction entities in the database.
func (usrcb *UserSpecialReductionCreateBulk) Save(ctx context.Context) ([]*UserSpecialReduction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(usrcb.builders))
	nodes := make([]*UserSpecialReduction, len(usrcb.builders))
	mutators := make([]Mutator, len(usrcb.builders))
	for i := range usrcb.builders {
		func(i int, root context.Context) {
			builder := usrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSpecialReductionMutation)
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
					_, err = mutators[i+1].Mutate(root, usrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = usrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, usrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, usrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (usrcb *UserSpecialReductionCreateBulk) SaveX(ctx context.Context) []*UserSpecialReduction {
	v, err := usrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usrcb *UserSpecialReductionCreateBulk) Exec(ctx context.Context) error {
	_, err := usrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usrcb *UserSpecialReductionCreateBulk) ExecX(ctx context.Context) {
	if err := usrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserSpecialReduction.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserSpecialReductionUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (usrcb *UserSpecialReductionCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserSpecialReductionUpsertBulk {
	usrcb.conflict = opts
	return &UserSpecialReductionUpsertBulk{
		create: usrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserSpecialReduction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (usrcb *UserSpecialReductionCreateBulk) OnConflictColumns(columns ...string) *UserSpecialReductionUpsertBulk {
	usrcb.conflict = append(usrcb.conflict, sql.ConflictColumns(columns...))
	return &UserSpecialReductionUpsertBulk{
		create: usrcb,
	}
}

// UserSpecialReductionUpsertBulk is the builder for "upsert"-ing
// a bulk of UserSpecialReduction nodes.
type UserSpecialReductionUpsertBulk struct {
	create *UserSpecialReductionCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserSpecialReduction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userspecialreduction.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserSpecialReductionUpsertBulk) UpdateNewValues() *UserSpecialReductionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userspecialreduction.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserSpecialReduction.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserSpecialReductionUpsertBulk) Ignore() *UserSpecialReductionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserSpecialReductionUpsertBulk) DoNothing() *UserSpecialReductionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserSpecialReductionCreateBulk.OnConflict
// documentation for more info.
func (u *UserSpecialReductionUpsertBulk) Update(set func(*UserSpecialReductionUpsert)) *UserSpecialReductionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserSpecialReductionUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserSpecialReductionUpsertBulk) SetAppID(v uuid.UUID) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateAppID() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserSpecialReductionUpsertBulk) SetUserID(v uuid.UUID) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateUserID() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateUserID()
	})
}

// SetAmount sets the "amount" field.
func (u *UserSpecialReductionUpsertBulk) SetAmount(v uint64) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateAmount() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateAmount()
	})
}

// SetReleaseByUserID sets the "release_by_user_id" field.
func (u *UserSpecialReductionUpsertBulk) SetReleaseByUserID(v uuid.UUID) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetReleaseByUserID(v)
	})
}

// UpdateReleaseByUserID sets the "release_by_user_id" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateReleaseByUserID() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateReleaseByUserID()
	})
}

// SetStart sets the "start" field.
func (u *UserSpecialReductionUpsertBulk) SetStart(v uint32) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateStart() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateStart()
	})
}

// SetDurationDays sets the "duration_days" field.
func (u *UserSpecialReductionUpsertBulk) SetDurationDays(v int32) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetDurationDays(v)
	})
}

// UpdateDurationDays sets the "duration_days" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateDurationDays() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateDurationDays()
	})
}

// SetMessage sets the "message" field.
func (u *UserSpecialReductionUpsertBulk) SetMessage(v string) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateMessage() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserSpecialReductionUpsertBulk) SetCreateAt(v uint32) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateCreateAt() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserSpecialReductionUpsertBulk) SetUpdateAt(v uint32) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateUpdateAt() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserSpecialReductionUpsertBulk) SetDeleteAt(v uint32) *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserSpecialReductionUpsertBulk) UpdateDeleteAt() *UserSpecialReductionUpsertBulk {
	return u.Update(func(s *UserSpecialReductionUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserSpecialReductionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserSpecialReductionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserSpecialReductionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserSpecialReductionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
