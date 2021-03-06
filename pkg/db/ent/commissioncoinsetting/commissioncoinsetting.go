// Code generated by ent, DO NOT EDIT.

package commissioncoinsetting

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the commissioncoinsetting type in the database.
	Label = "commission_coin_setting"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldUsing holds the string denoting the using field in the database.
	FieldUsing = "using"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the commissioncoinsetting in the database.
	Table = "commission_coin_settings"
)

// Columns holds all SQL columns for commissioncoinsetting fields.
var Columns = []string{
	FieldID,
	FieldCoinTypeID,
	FieldUsing,
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
	// DefaultUsing holds the default value on creation for the "using" field.
	DefaultUsing bool
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
