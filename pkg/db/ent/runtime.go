// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userinvitationcodeFields := schema.UserInvitationCode{}.Fields()
	_ = userinvitationcodeFields
	// userinvitationcodeDescCreateAt is the schema descriptor for create_at field.
	userinvitationcodeDescCreateAt := userinvitationcodeFields[4].Descriptor()
	// userinvitationcode.DefaultCreateAt holds the default value on creation for the create_at field.
	userinvitationcode.DefaultCreateAt = userinvitationcodeDescCreateAt.Default.(func() uint32)
	// userinvitationcodeDescUpdateAt is the schema descriptor for update_at field.
	userinvitationcodeDescUpdateAt := userinvitationcodeFields[5].Descriptor()
	// userinvitationcode.DefaultUpdateAt holds the default value on creation for the update_at field.
	userinvitationcode.DefaultUpdateAt = userinvitationcodeDescUpdateAt.Default.(func() uint32)
	// userinvitationcode.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userinvitationcode.UpdateDefaultUpdateAt = userinvitationcodeDescUpdateAt.UpdateDefault.(func() uint32)
	// userinvitationcodeDescDeleteAt is the schema descriptor for delete_at field.
	userinvitationcodeDescDeleteAt := userinvitationcodeFields[6].Descriptor()
	// userinvitationcode.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userinvitationcode.DefaultDeleteAt = userinvitationcodeDescDeleteAt.Default.(func() uint32)
	// userinvitationcodeDescID is the schema descriptor for id field.
	userinvitationcodeDescID := userinvitationcodeFields[0].Descriptor()
	// userinvitationcode.DefaultID holds the default value on creation for the id field.
	userinvitationcode.DefaultID = userinvitationcodeDescID.Default.(func() uuid.UUID)
}
