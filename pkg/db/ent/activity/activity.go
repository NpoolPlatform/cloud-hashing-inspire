// Code generated by entc, DO NOT EDIT.

package activity

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the activity type in the database.
	Label = "activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// FieldSystemActivity holds the string denoting the system_activity field in the database.
	FieldSystemActivity = "system_activity"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the activity in the database.
	Table = "activities"
)

// Columns holds all SQL columns for activity fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldCreatedBy,
	FieldName,
	FieldStart,
	FieldEnd,
	FieldSystemActivity,
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
