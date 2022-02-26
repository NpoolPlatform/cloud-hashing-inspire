// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/apppurchaseamountsetting"
	"github.com/google/uuid"
)

// AppPurchaseAmountSetting is the model entity for the AppPurchaseAmountSetting schema.
type AppPurchaseAmountSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// BadgeLarge holds the value of the "badge_large" field.
	BadgeLarge string `json:"badge_large,omitempty"`
	// BadgeSmall holds the value of the "badge_small" field.
	BadgeSmall string `json:"badge_small,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// Percent holds the value of the "percent" field.
	Percent uint32 `json:"percent,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppPurchaseAmountSetting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case apppurchaseamountsetting.FieldAmount, apppurchaseamountsetting.FieldPercent, apppurchaseamountsetting.FieldCreateAt, apppurchaseamountsetting.FieldUpdateAt, apppurchaseamountsetting.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case apppurchaseamountsetting.FieldTitle, apppurchaseamountsetting.FieldBadgeLarge, apppurchaseamountsetting.FieldBadgeSmall:
			values[i] = new(sql.NullString)
		case apppurchaseamountsetting.FieldID, apppurchaseamountsetting.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppPurchaseAmountSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppPurchaseAmountSetting fields.
func (apas *AppPurchaseAmountSetting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apppurchaseamountsetting.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				apas.ID = *value
			}
		case apppurchaseamountsetting.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				apas.AppID = *value
			}
		case apppurchaseamountsetting.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				apas.Title = value.String
			}
		case apppurchaseamountsetting.FieldBadgeLarge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field badge_large", values[i])
			} else if value.Valid {
				apas.BadgeLarge = value.String
			}
		case apppurchaseamountsetting.FieldBadgeSmall:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field badge_small", values[i])
			} else if value.Valid {
				apas.BadgeSmall = value.String
			}
		case apppurchaseamountsetting.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				apas.Amount = uint64(value.Int64)
			}
		case apppurchaseamountsetting.FieldPercent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field percent", values[i])
			} else if value.Valid {
				apas.Percent = uint32(value.Int64)
			}
		case apppurchaseamountsetting.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				apas.CreateAt = uint32(value.Int64)
			}
		case apppurchaseamountsetting.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				apas.UpdateAt = uint32(value.Int64)
			}
		case apppurchaseamountsetting.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				apas.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppPurchaseAmountSetting.
// Note that you need to call AppPurchaseAmountSetting.Unwrap() before calling this method if this AppPurchaseAmountSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (apas *AppPurchaseAmountSetting) Update() *AppPurchaseAmountSettingUpdateOne {
	return (&AppPurchaseAmountSettingClient{config: apas.config}).UpdateOne(apas)
}

// Unwrap unwraps the AppPurchaseAmountSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (apas *AppPurchaseAmountSetting) Unwrap() *AppPurchaseAmountSetting {
	tx, ok := apas.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppPurchaseAmountSetting is not a transactional entity")
	}
	apas.config.driver = tx.drv
	return apas
}

// String implements the fmt.Stringer.
func (apas *AppPurchaseAmountSetting) String() string {
	var builder strings.Builder
	builder.WriteString("AppPurchaseAmountSetting(")
	builder.WriteString(fmt.Sprintf("id=%v", apas.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", apas.AppID))
	builder.WriteString(", title=")
	builder.WriteString(apas.Title)
	builder.WriteString(", badge_large=")
	builder.WriteString(apas.BadgeLarge)
	builder.WriteString(", badge_small=")
	builder.WriteString(apas.BadgeSmall)
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", apas.Amount))
	builder.WriteString(", percent=")
	builder.WriteString(fmt.Sprintf("%v", apas.Percent))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", apas.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", apas.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", apas.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// AppPurchaseAmountSettings is a parsable slice of AppPurchaseAmountSetting.
type AppPurchaseAmountSettings []*AppPurchaseAmountSetting

func (apas AppPurchaseAmountSettings) config(cfg config) {
	for _i := range apas {
		apas[_i].config = cfg
	}
}
