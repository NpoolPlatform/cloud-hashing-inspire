package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// UserSpecialReduction holds the schema definition for the UserSpecialReduction entity.
type UserSpecialReduction struct {
	ent.Schema
}

// Fields of the UserSpecialReduction.
func (UserSpecialReduction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Uint64("amount"),
		field.UUID("release_by_user_id", uuid.UUID{}),
		field.Uint32("start"),
		field.Int32("duration_days"),
		field.String("message").
			MaxLen(512),
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

// Edges of the UserSpecialReduction.
func (UserSpecialReduction) Edges() []ent.Edge {
	return nil
}
