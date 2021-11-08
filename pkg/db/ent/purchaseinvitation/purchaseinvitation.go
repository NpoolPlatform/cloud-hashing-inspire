// Code generated by entc, DO NOT EDIT.

package purchaseinvitation

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the purchaseinvitation type in the database.
	Label = "purchase_invitation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldInvitationCodeID holds the string denoting the invitation_code_id field in the database.
	FieldInvitationCodeID = "invitation_code_id"
	// FieldFullfilled holds the string denoting the fullfilled field in the database.
	FieldFullfilled = "fullfilled"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the purchaseinvitation in the database.
	Table = "purchase_invitations"
)

// Columns holds all SQL columns for purchaseinvitation fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldOrderID,
	FieldInvitationCodeID,
	FieldFullfilled,
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
	// DefaultFullfilled holds the default value on creation for the "fullfilled" field.
	DefaultFullfilled bool
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
