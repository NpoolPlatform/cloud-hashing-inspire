package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// AppPurchaseAmountSetting holds the schema definition for the AppPurchaseAmountSetting entity.
type AppPurchaseAmountSetting struct {
	ent.Schema
}

// Fields of the AppPurchaseAmountSetting.
func (AppPurchaseAmountSetting) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.UUID("good_id", uuid.UUID{}),
		field.String("title"),
		field.Uint64("amount"),
		field.Uint32("percent"),
		field.Uint32("start"),
		field.Uint32("end"),
		field.String("badge_large"),
		field.String("badge_small"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the AppPurchaseAmountSetting.
func (AppPurchaseAmountSetting) Edges() []ent.Edge {
	return nil
}

// Indexes of the AppPurchaseAmountSetting.
func (AppPurchaseAmountSetting) Indexes() []ent.Index {
	return nil
}
