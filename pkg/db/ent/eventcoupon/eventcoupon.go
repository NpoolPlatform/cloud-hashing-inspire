// Code generated by entc, DO NOT EDIT.

package eventcoupon

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the eventcoupon type in the database.
	Label = "event_coupon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldActivityID holds the string denoting the activity_id field in the database.
	FieldActivityID = "activity_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCouponID holds the string denoting the coupon_id field in the database.
	FieldCouponID = "coupon_id"
	// FieldEvent holds the string denoting the event field in the database.
	FieldEvent = "event"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the eventcoupon in the database.
	Table = "event_coupons"
)

// Columns holds all SQL columns for eventcoupon fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldActivityID,
	FieldType,
	FieldCouponID,
	FieldEvent,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EventValidator is a validator for the "event" field. It is called by the builders before save.
	EventValidator func(string) error
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
