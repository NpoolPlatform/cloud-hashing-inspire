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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appinvitationsetting"
	"github.com/google/uuid"
)

// AppInvitationSettingCreate is the builder for creating a AppInvitationSetting entity.
type AppInvitationSettingCreate struct {
	config
	mutation *AppInvitationSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (aisc *AppInvitationSettingCreate) SetAppID(u uuid.UUID) *AppInvitationSettingCreate {
	aisc.mutation.SetAppID(u)
	return aisc
}

// SetCount sets the "count" field.
func (aisc *AppInvitationSettingCreate) SetCount(u uint32) *AppInvitationSettingCreate {
	aisc.mutation.SetCount(u)
	return aisc
}

// SetDiscount sets the "discount" field.
func (aisc *AppInvitationSettingCreate) SetDiscount(u uint32) *AppInvitationSettingCreate {
	aisc.mutation.SetDiscount(u)
	return aisc
}

// SetCreateAt sets the "create_at" field.
func (aisc *AppInvitationSettingCreate) SetCreateAt(u uint32) *AppInvitationSettingCreate {
	aisc.mutation.SetCreateAt(u)
	return aisc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (aisc *AppInvitationSettingCreate) SetNillableCreateAt(u *uint32) *AppInvitationSettingCreate {
	if u != nil {
		aisc.SetCreateAt(*u)
	}
	return aisc
}

// SetUpdateAt sets the "update_at" field.
func (aisc *AppInvitationSettingCreate) SetUpdateAt(u uint32) *AppInvitationSettingCreate {
	aisc.mutation.SetUpdateAt(u)
	return aisc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (aisc *AppInvitationSettingCreate) SetNillableUpdateAt(u *uint32) *AppInvitationSettingCreate {
	if u != nil {
		aisc.SetUpdateAt(*u)
	}
	return aisc
}

// SetDeleteAt sets the "delete_at" field.
func (aisc *AppInvitationSettingCreate) SetDeleteAt(u uint32) *AppInvitationSettingCreate {
	aisc.mutation.SetDeleteAt(u)
	return aisc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aisc *AppInvitationSettingCreate) SetNillableDeleteAt(u *uint32) *AppInvitationSettingCreate {
	if u != nil {
		aisc.SetDeleteAt(*u)
	}
	return aisc
}

// SetTitle sets the "title" field.
func (aisc *AppInvitationSettingCreate) SetTitle(s string) *AppInvitationSettingCreate {
	aisc.mutation.SetTitle(s)
	return aisc
}

// SetBadgeLarge sets the "badge_large" field.
func (aisc *AppInvitationSettingCreate) SetBadgeLarge(s string) *AppInvitationSettingCreate {
	aisc.mutation.SetBadgeLarge(s)
	return aisc
}

// SetBadgeSmall sets the "badge_small" field.
func (aisc *AppInvitationSettingCreate) SetBadgeSmall(s string) *AppInvitationSettingCreate {
	aisc.mutation.SetBadgeSmall(s)
	return aisc
}

// SetID sets the "id" field.
func (aisc *AppInvitationSettingCreate) SetID(u uuid.UUID) *AppInvitationSettingCreate {
	aisc.mutation.SetID(u)
	return aisc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (aisc *AppInvitationSettingCreate) SetNillableID(u *uuid.UUID) *AppInvitationSettingCreate {
	if u != nil {
		aisc.SetID(*u)
	}
	return aisc
}

// Mutation returns the AppInvitationSettingMutation object of the builder.
func (aisc *AppInvitationSettingCreate) Mutation() *AppInvitationSettingMutation {
	return aisc.mutation
}

// Save creates the AppInvitationSetting in the database.
func (aisc *AppInvitationSettingCreate) Save(ctx context.Context) (*AppInvitationSetting, error) {
	var (
		err  error
		node *AppInvitationSetting
	)
	aisc.defaults()
	if len(aisc.hooks) == 0 {
		if err = aisc.check(); err != nil {
			return nil, err
		}
		node, err = aisc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppInvitationSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aisc.check(); err != nil {
				return nil, err
			}
			aisc.mutation = mutation
			if node, err = aisc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aisc.hooks) - 1; i >= 0; i-- {
			if aisc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aisc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aisc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aisc *AppInvitationSettingCreate) SaveX(ctx context.Context) *AppInvitationSetting {
	v, err := aisc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aisc *AppInvitationSettingCreate) Exec(ctx context.Context) error {
	_, err := aisc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aisc *AppInvitationSettingCreate) ExecX(ctx context.Context) {
	if err := aisc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aisc *AppInvitationSettingCreate) defaults() {
	if _, ok := aisc.mutation.CreateAt(); !ok {
		v := appinvitationsetting.DefaultCreateAt()
		aisc.mutation.SetCreateAt(v)
	}
	if _, ok := aisc.mutation.UpdateAt(); !ok {
		v := appinvitationsetting.DefaultUpdateAt()
		aisc.mutation.SetUpdateAt(v)
	}
	if _, ok := aisc.mutation.DeleteAt(); !ok {
		v := appinvitationsetting.DefaultDeleteAt()
		aisc.mutation.SetDeleteAt(v)
	}
	if _, ok := aisc.mutation.ID(); !ok {
		v := appinvitationsetting.DefaultID()
		aisc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aisc *AppInvitationSettingCreate) check() error {
	if _, ok := aisc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppInvitationSetting.app_id"`)}
	}
	if _, ok := aisc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`ent: missing required field "AppInvitationSetting.count"`)}
	}
	if _, ok := aisc.mutation.Discount(); !ok {
		return &ValidationError{Name: "discount", err: errors.New(`ent: missing required field "AppInvitationSetting.discount"`)}
	}
	if _, ok := aisc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "AppInvitationSetting.create_at"`)}
	}
	if _, ok := aisc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "AppInvitationSetting.update_at"`)}
	}
	if _, ok := aisc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "AppInvitationSetting.delete_at"`)}
	}
	if _, ok := aisc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "AppInvitationSetting.title"`)}
	}
	if _, ok := aisc.mutation.BadgeLarge(); !ok {
		return &ValidationError{Name: "badge_large", err: errors.New(`ent: missing required field "AppInvitationSetting.badge_large"`)}
	}
	if _, ok := aisc.mutation.BadgeSmall(); !ok {
		return &ValidationError{Name: "badge_small", err: errors.New(`ent: missing required field "AppInvitationSetting.badge_small"`)}
	}
	return nil
}

func (aisc *AppInvitationSettingCreate) sqlSave(ctx context.Context) (*AppInvitationSetting, error) {
	_node, _spec := aisc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aisc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (aisc *AppInvitationSettingCreate) createSpec() (*AppInvitationSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &AppInvitationSetting{config: aisc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appinvitationsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appinvitationsetting.FieldID,
			},
		}
	)
	_spec.OnConflict = aisc.conflict
	if id, ok := aisc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := aisc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appinvitationsetting.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := aisc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appinvitationsetting.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := aisc.mutation.Discount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appinvitationsetting.FieldDiscount,
		})
		_node.Discount = value
	}
	if value, ok := aisc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appinvitationsetting.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := aisc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appinvitationsetting.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := aisc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appinvitationsetting.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	if value, ok := aisc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appinvitationsetting.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := aisc.mutation.BadgeLarge(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appinvitationsetting.FieldBadgeLarge,
		})
		_node.BadgeLarge = value
	}
	if value, ok := aisc.mutation.BadgeSmall(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appinvitationsetting.FieldBadgeSmall,
		})
		_node.BadgeSmall = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppInvitationSetting.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppInvitationSettingUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (aisc *AppInvitationSettingCreate) OnConflict(opts ...sql.ConflictOption) *AppInvitationSettingUpsertOne {
	aisc.conflict = opts
	return &AppInvitationSettingUpsertOne{
		create: aisc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppInvitationSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (aisc *AppInvitationSettingCreate) OnConflictColumns(columns ...string) *AppInvitationSettingUpsertOne {
	aisc.conflict = append(aisc.conflict, sql.ConflictColumns(columns...))
	return &AppInvitationSettingUpsertOne{
		create: aisc,
	}
}

type (
	// AppInvitationSettingUpsertOne is the builder for "upsert"-ing
	//  one AppInvitationSetting node.
	AppInvitationSettingUpsertOne struct {
		create *AppInvitationSettingCreate
	}

	// AppInvitationSettingUpsert is the "OnConflict" setter.
	AppInvitationSettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *AppInvitationSettingUpsert) SetAppID(v uuid.UUID) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateAppID() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldAppID)
	return u
}

// SetCount sets the "count" field.
func (u *AppInvitationSettingUpsert) SetCount(v uint32) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldCount, v)
	return u
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateCount() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldCount)
	return u
}

// AddCount adds v to the "count" field.
func (u *AppInvitationSettingUpsert) AddCount(v uint32) *AppInvitationSettingUpsert {
	u.Add(appinvitationsetting.FieldCount, v)
	return u
}

// SetDiscount sets the "discount" field.
func (u *AppInvitationSettingUpsert) SetDiscount(v uint32) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldDiscount, v)
	return u
}

// UpdateDiscount sets the "discount" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateDiscount() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldDiscount)
	return u
}

// AddDiscount adds v to the "discount" field.
func (u *AppInvitationSettingUpsert) AddDiscount(v uint32) *AppInvitationSettingUpsert {
	u.Add(appinvitationsetting.FieldDiscount, v)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppInvitationSettingUpsert) SetCreateAt(v uint32) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateCreateAt() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppInvitationSettingUpsert) AddCreateAt(v uint32) *AppInvitationSettingUpsert {
	u.Add(appinvitationsetting.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *AppInvitationSettingUpsert) SetUpdateAt(v uint32) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateUpdateAt() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppInvitationSettingUpsert) AddUpdateAt(v uint32) *AppInvitationSettingUpsert {
	u.Add(appinvitationsetting.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppInvitationSettingUpsert) SetDeleteAt(v uint32) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateDeleteAt() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppInvitationSettingUpsert) AddDeleteAt(v uint32) *AppInvitationSettingUpsert {
	u.Add(appinvitationsetting.FieldDeleteAt, v)
	return u
}

// SetTitle sets the "title" field.
func (u *AppInvitationSettingUpsert) SetTitle(v string) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateTitle() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldTitle)
	return u
}

// SetBadgeLarge sets the "badge_large" field.
func (u *AppInvitationSettingUpsert) SetBadgeLarge(v string) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldBadgeLarge, v)
	return u
}

// UpdateBadgeLarge sets the "badge_large" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateBadgeLarge() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldBadgeLarge)
	return u
}

// SetBadgeSmall sets the "badge_small" field.
func (u *AppInvitationSettingUpsert) SetBadgeSmall(v string) *AppInvitationSettingUpsert {
	u.Set(appinvitationsetting.FieldBadgeSmall, v)
	return u
}

// UpdateBadgeSmall sets the "badge_small" field to the value that was provided on create.
func (u *AppInvitationSettingUpsert) UpdateBadgeSmall() *AppInvitationSettingUpsert {
	u.SetExcluded(appinvitationsetting.FieldBadgeSmall)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppInvitationSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appinvitationsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppInvitationSettingUpsertOne) UpdateNewValues() *AppInvitationSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appinvitationsetting.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppInvitationSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppInvitationSettingUpsertOne) Ignore() *AppInvitationSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppInvitationSettingUpsertOne) DoNothing() *AppInvitationSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppInvitationSettingCreate.OnConflict
// documentation for more info.
func (u *AppInvitationSettingUpsertOne) Update(set func(*AppInvitationSettingUpsert)) *AppInvitationSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppInvitationSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppInvitationSettingUpsertOne) SetAppID(v uuid.UUID) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateAppID() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetCount sets the "count" field.
func (u *AppInvitationSettingUpsertOne) SetCount(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *AppInvitationSettingUpsertOne) AddCount(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateCount() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateCount()
	})
}

// SetDiscount sets the "discount" field.
func (u *AppInvitationSettingUpsertOne) SetDiscount(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetDiscount(v)
	})
}

// AddDiscount adds v to the "discount" field.
func (u *AppInvitationSettingUpsertOne) AddDiscount(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddDiscount(v)
	})
}

// UpdateDiscount sets the "discount" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateDiscount() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateDiscount()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppInvitationSettingUpsertOne) SetCreateAt(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppInvitationSettingUpsertOne) AddCreateAt(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateCreateAt() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppInvitationSettingUpsertOne) SetUpdateAt(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppInvitationSettingUpsertOne) AddUpdateAt(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateUpdateAt() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppInvitationSettingUpsertOne) SetDeleteAt(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppInvitationSettingUpsertOne) AddDeleteAt(v uint32) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateDeleteAt() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetTitle sets the "title" field.
func (u *AppInvitationSettingUpsertOne) SetTitle(v string) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateTitle() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateTitle()
	})
}

// SetBadgeLarge sets the "badge_large" field.
func (u *AppInvitationSettingUpsertOne) SetBadgeLarge(v string) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetBadgeLarge(v)
	})
}

// UpdateBadgeLarge sets the "badge_large" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateBadgeLarge() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateBadgeLarge()
	})
}

// SetBadgeSmall sets the "badge_small" field.
func (u *AppInvitationSettingUpsertOne) SetBadgeSmall(v string) *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetBadgeSmall(v)
	})
}

// UpdateBadgeSmall sets the "badge_small" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertOne) UpdateBadgeSmall() *AppInvitationSettingUpsertOne {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateBadgeSmall()
	})
}

// Exec executes the query.
func (u *AppInvitationSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppInvitationSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppInvitationSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppInvitationSettingUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppInvitationSettingUpsertOne.ID is not supported by MySQL driver. Use AppInvitationSettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppInvitationSettingUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppInvitationSettingCreateBulk is the builder for creating many AppInvitationSetting entities in bulk.
type AppInvitationSettingCreateBulk struct {
	config
	builders []*AppInvitationSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the AppInvitationSetting entities in the database.
func (aiscb *AppInvitationSettingCreateBulk) Save(ctx context.Context) ([]*AppInvitationSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aiscb.builders))
	nodes := make([]*AppInvitationSetting, len(aiscb.builders))
	mutators := make([]Mutator, len(aiscb.builders))
	for i := range aiscb.builders {
		func(i int, root context.Context) {
			builder := aiscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppInvitationSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, aiscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = aiscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aiscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aiscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aiscb *AppInvitationSettingCreateBulk) SaveX(ctx context.Context) []*AppInvitationSetting {
	v, err := aiscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aiscb *AppInvitationSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := aiscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiscb *AppInvitationSettingCreateBulk) ExecX(ctx context.Context) {
	if err := aiscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppInvitationSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppInvitationSettingUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (aiscb *AppInvitationSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppInvitationSettingUpsertBulk {
	aiscb.conflict = opts
	return &AppInvitationSettingUpsertBulk{
		create: aiscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppInvitationSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (aiscb *AppInvitationSettingCreateBulk) OnConflictColumns(columns ...string) *AppInvitationSettingUpsertBulk {
	aiscb.conflict = append(aiscb.conflict, sql.ConflictColumns(columns...))
	return &AppInvitationSettingUpsertBulk{
		create: aiscb,
	}
}

// AppInvitationSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of AppInvitationSetting nodes.
type AppInvitationSettingUpsertBulk struct {
	create *AppInvitationSettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppInvitationSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appinvitationsetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppInvitationSettingUpsertBulk) UpdateNewValues() *AppInvitationSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appinvitationsetting.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppInvitationSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppInvitationSettingUpsertBulk) Ignore() *AppInvitationSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppInvitationSettingUpsertBulk) DoNothing() *AppInvitationSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppInvitationSettingCreateBulk.OnConflict
// documentation for more info.
func (u *AppInvitationSettingUpsertBulk) Update(set func(*AppInvitationSettingUpsert)) *AppInvitationSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppInvitationSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppInvitationSettingUpsertBulk) SetAppID(v uuid.UUID) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateAppID() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetCount sets the "count" field.
func (u *AppInvitationSettingUpsertBulk) SetCount(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *AppInvitationSettingUpsertBulk) AddCount(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateCount() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateCount()
	})
}

// SetDiscount sets the "discount" field.
func (u *AppInvitationSettingUpsertBulk) SetDiscount(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetDiscount(v)
	})
}

// AddDiscount adds v to the "discount" field.
func (u *AppInvitationSettingUpsertBulk) AddDiscount(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddDiscount(v)
	})
}

// UpdateDiscount sets the "discount" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateDiscount() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateDiscount()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppInvitationSettingUpsertBulk) SetCreateAt(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppInvitationSettingUpsertBulk) AddCreateAt(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateCreateAt() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppInvitationSettingUpsertBulk) SetUpdateAt(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppInvitationSettingUpsertBulk) AddUpdateAt(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateUpdateAt() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppInvitationSettingUpsertBulk) SetDeleteAt(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppInvitationSettingUpsertBulk) AddDeleteAt(v uint32) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateDeleteAt() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetTitle sets the "title" field.
func (u *AppInvitationSettingUpsertBulk) SetTitle(v string) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateTitle() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateTitle()
	})
}

// SetBadgeLarge sets the "badge_large" field.
func (u *AppInvitationSettingUpsertBulk) SetBadgeLarge(v string) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetBadgeLarge(v)
	})
}

// UpdateBadgeLarge sets the "badge_large" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateBadgeLarge() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateBadgeLarge()
	})
}

// SetBadgeSmall sets the "badge_small" field.
func (u *AppInvitationSettingUpsertBulk) SetBadgeSmall(v string) *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.SetBadgeSmall(v)
	})
}

// UpdateBadgeSmall sets the "badge_small" field to the value that was provided on create.
func (u *AppInvitationSettingUpsertBulk) UpdateBadgeSmall() *AppInvitationSettingUpsertBulk {
	return u.Update(func(s *AppInvitationSettingUpsert) {
		s.UpdateBadgeSmall()
	})
}

// Exec executes the query.
func (u *AppInvitationSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppInvitationSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppInvitationSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppInvitationSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}