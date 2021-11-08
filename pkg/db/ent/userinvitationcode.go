// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"
	"github.com/google/uuid"
)

// UserInvitationCode is the model entity for the UserInvitationCode schema.
type UserInvitationCode struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// InvitationCode holds the value of the "invitation_code" field.
	InvitationCode string `json:"invitation_code,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserInvitationCode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userinvitationcode.FieldCreateAt, userinvitationcode.FieldUpdateAt, userinvitationcode.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case userinvitationcode.FieldInvitationCode:
			values[i] = new(sql.NullString)
		case userinvitationcode.FieldID, userinvitationcode.FieldUserID, userinvitationcode.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserInvitationCode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserInvitationCode fields.
func (uic *UserInvitationCode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userinvitationcode.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				uic.ID = *value
			}
		case userinvitationcode.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				uic.UserID = *value
			}
		case userinvitationcode.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				uic.AppID = *value
			}
		case userinvitationcode.FieldInvitationCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invitation_code", values[i])
			} else if value.Valid {
				uic.InvitationCode = value.String
			}
		case userinvitationcode.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				uic.CreateAt = uint32(value.Int64)
			}
		case userinvitationcode.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				uic.UpdateAt = uint32(value.Int64)
			}
		case userinvitationcode.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				uic.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserInvitationCode.
// Note that you need to call UserInvitationCode.Unwrap() before calling this method if this UserInvitationCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (uic *UserInvitationCode) Update() *UserInvitationCodeUpdateOne {
	return (&UserInvitationCodeClient{config: uic.config}).UpdateOne(uic)
}

// Unwrap unwraps the UserInvitationCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uic *UserInvitationCode) Unwrap() *UserInvitationCode {
	tx, ok := uic.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserInvitationCode is not a transactional entity")
	}
	uic.config.driver = tx.drv
	return uic
}

// String implements the fmt.Stringer.
func (uic *UserInvitationCode) String() string {
	var builder strings.Builder
	builder.WriteString("UserInvitationCode(")
	builder.WriteString(fmt.Sprintf("id=%v", uic.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", uic.UserID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", uic.AppID))
	builder.WriteString(", invitation_code=")
	builder.WriteString(uic.InvitationCode)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", uic.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", uic.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", uic.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// UserInvitationCodes is a parsable slice of UserInvitationCode.
type UserInvitationCodes []*UserInvitationCode

func (uic UserInvitationCodes) config(cfg config) {
	for _i := range uic {
		uic[_i].config = cfg
	}
}
