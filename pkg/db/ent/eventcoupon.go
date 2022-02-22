// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/eventcoupon"
	"github.com/google/uuid"
)

// EventCoupon is the model entity for the EventCoupon schema.
type EventCoupon struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// ActivityID holds the value of the "activity_id" field.
	ActivityID uuid.UUID `json:"activity_id,omitempty"`
	// CouponID holds the value of the "coupon_id" field.
	CouponID uuid.UUID `json:"coupon_id,omitempty"`
	// Event holds the value of the "event" field.
	Event string `json:"event,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EventCoupon) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case eventcoupon.FieldCreateAt, eventcoupon.FieldUpdateAt, eventcoupon.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case eventcoupon.FieldEvent:
			values[i] = new(sql.NullString)
		case eventcoupon.FieldID, eventcoupon.FieldAppID, eventcoupon.FieldActivityID, eventcoupon.FieldCouponID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EventCoupon", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EventCoupon fields.
func (ec *EventCoupon) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case eventcoupon.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ec.ID = *value
			}
		case eventcoupon.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ec.AppID = *value
			}
		case eventcoupon.FieldActivityID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field activity_id", values[i])
			} else if value != nil {
				ec.ActivityID = *value
			}
		case eventcoupon.FieldCouponID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coupon_id", values[i])
			} else if value != nil {
				ec.CouponID = *value
			}
		case eventcoupon.FieldEvent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event", values[i])
			} else if value.Valid {
				ec.Event = value.String
			}
		case eventcoupon.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ec.CreateAt = uint32(value.Int64)
			}
		case eventcoupon.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ec.UpdateAt = uint32(value.Int64)
			}
		case eventcoupon.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ec.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this EventCoupon.
// Note that you need to call EventCoupon.Unwrap() before calling this method if this EventCoupon
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EventCoupon) Update() *EventCouponUpdateOne {
	return (&EventCouponClient{config: ec.config}).UpdateOne(ec)
}

// Unwrap unwraps the EventCoupon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EventCoupon) Unwrap() *EventCoupon {
	tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: EventCoupon is not a transactional entity")
	}
	ec.config.driver = tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EventCoupon) String() string {
	var builder strings.Builder
	builder.WriteString("EventCoupon(")
	builder.WriteString(fmt.Sprintf("id=%v", ec.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.AppID))
	builder.WriteString(", activity_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.ActivityID))
	builder.WriteString(", coupon_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.CouponID))
	builder.WriteString(", event=")
	builder.WriteString(ec.Event)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ec.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ec.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ec.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// EventCoupons is a parsable slice of EventCoupon.
type EventCoupons []*EventCoupon

func (ec EventCoupons) config(cfg config) {
	for _i := range ec {
		ec[_i].config = cfg
	}
}
