// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/purchaseinvitation"
	"github.com/google/uuid"
)

// PurchaseInvitation is the model entity for the PurchaseInvitation schema.
type PurchaseInvitation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// InvitationCodeID holds the value of the "invitation_code_id" field.
	InvitationCodeID uuid.UUID `json:"invitation_code_id,omitempty"`
	// Fullfilled holds the value of the "fullfilled" field.
	Fullfilled bool `json:"fullfilled,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PurchaseInvitation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case purchaseinvitation.FieldFullfilled:
			values[i] = new(sql.NullBool)
		case purchaseinvitation.FieldCreateAt, purchaseinvitation.FieldUpdateAt, purchaseinvitation.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case purchaseinvitation.FieldID, purchaseinvitation.FieldAppID, purchaseinvitation.FieldOrderID, purchaseinvitation.FieldInvitationCodeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PurchaseInvitation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PurchaseInvitation fields.
func (pi *PurchaseInvitation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case purchaseinvitation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pi.ID = *value
			}
		case purchaseinvitation.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				pi.AppID = *value
			}
		case purchaseinvitation.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				pi.OrderID = *value
			}
		case purchaseinvitation.FieldInvitationCodeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field invitation_code_id", values[i])
			} else if value != nil {
				pi.InvitationCodeID = *value
			}
		case purchaseinvitation.FieldFullfilled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field fullfilled", values[i])
			} else if value.Valid {
				pi.Fullfilled = value.Bool
			}
		case purchaseinvitation.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				pi.CreateAt = uint32(value.Int64)
			}
		case purchaseinvitation.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				pi.UpdateAt = uint32(value.Int64)
			}
		case purchaseinvitation.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				pi.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PurchaseInvitation.
// Note that you need to call PurchaseInvitation.Unwrap() before calling this method if this PurchaseInvitation
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *PurchaseInvitation) Update() *PurchaseInvitationUpdateOne {
	return (&PurchaseInvitationClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the PurchaseInvitation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *PurchaseInvitation) Unwrap() *PurchaseInvitation {
	tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: PurchaseInvitation is not a transactional entity")
	}
	pi.config.driver = tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *PurchaseInvitation) String() string {
	var builder strings.Builder
	builder.WriteString("PurchaseInvitation(")
	builder.WriteString(fmt.Sprintf("id=%v", pi.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.AppID))
	builder.WriteString(", order_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.OrderID))
	builder.WriteString(", invitation_code_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.InvitationCodeID))
	builder.WriteString(", fullfilled=")
	builder.WriteString(fmt.Sprintf("%v", pi.Fullfilled))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", pi.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", pi.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", pi.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// PurchaseInvitations is a parsable slice of PurchaseInvitation.
type PurchaseInvitations []*PurchaseInvitation

func (pi PurchaseInvitations) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}
