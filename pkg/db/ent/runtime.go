// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/agencysetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appcouponsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/couponpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/purchaseinvitation"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userinvitationcode"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	agencysettingFields := schema.AgencySetting{}.Fields()
	_ = agencysettingFields
	// agencysettingDescCreateAt is the schema descriptor for create_at field.
	agencysettingDescCreateAt := agencysettingFields[8].Descriptor()
	// agencysetting.DefaultCreateAt holds the default value on creation for the create_at field.
	agencysetting.DefaultCreateAt = agencysettingDescCreateAt.Default.(func() uint32)
	// agencysettingDescUpdateAt is the schema descriptor for update_at field.
	agencysettingDescUpdateAt := agencysettingFields[9].Descriptor()
	// agencysetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	agencysetting.DefaultUpdateAt = agencysettingDescUpdateAt.Default.(func() uint32)
	// agencysetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	agencysetting.UpdateDefaultUpdateAt = agencysettingDescUpdateAt.UpdateDefault.(func() uint32)
	// agencysettingDescDeleteAt is the schema descriptor for delete_at field.
	agencysettingDescDeleteAt := agencysettingFields[10].Descriptor()
	// agencysetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	agencysetting.DefaultDeleteAt = agencysettingDescDeleteAt.Default.(func() uint32)
	// agencysettingDescID is the schema descriptor for id field.
	agencysettingDescID := agencysettingFields[0].Descriptor()
	// agencysetting.DefaultID holds the default value on creation for the id field.
	agencysetting.DefaultID = agencysettingDescID.Default.(func() uuid.UUID)
	appcouponsettingFields := schema.AppCouponSetting{}.Fields()
	_ = appcouponsettingFields
	// appcouponsettingDescCreateAt is the schema descriptor for create_at field.
	appcouponsettingDescCreateAt := appcouponsettingFields[4].Descriptor()
	// appcouponsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	appcouponsetting.DefaultCreateAt = appcouponsettingDescCreateAt.Default.(func() uint32)
	// appcouponsettingDescUpdateAt is the schema descriptor for update_at field.
	appcouponsettingDescUpdateAt := appcouponsettingFields[5].Descriptor()
	// appcouponsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	appcouponsetting.DefaultUpdateAt = appcouponsettingDescUpdateAt.Default.(func() uint32)
	// appcouponsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appcouponsetting.UpdateDefaultUpdateAt = appcouponsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// appcouponsettingDescDeleteAt is the schema descriptor for delete_at field.
	appcouponsettingDescDeleteAt := appcouponsettingFields[6].Descriptor()
	// appcouponsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appcouponsetting.DefaultDeleteAt = appcouponsettingDescDeleteAt.Default.(func() uint32)
	// appcouponsettingDescID is the schema descriptor for id field.
	appcouponsettingDescID := appcouponsettingFields[0].Descriptor()
	// appcouponsetting.DefaultID holds the default value on creation for the id field.
	appcouponsetting.DefaultID = appcouponsettingDescID.Default.(func() uuid.UUID)
	couponallocatedFields := schema.CouponAllocated{}.Fields()
	_ = couponallocatedFields
	// couponallocatedDescUsed is the schema descriptor for used field.
	couponallocatedDescUsed := couponallocatedFields[3].Descriptor()
	// couponallocated.DefaultUsed holds the default value on creation for the used field.
	couponallocated.DefaultUsed = couponallocatedDescUsed.Default.(bool)
	// couponallocatedDescCreateAt is the schema descriptor for create_at field.
	couponallocatedDescCreateAt := couponallocatedFields[5].Descriptor()
	// couponallocated.DefaultCreateAt holds the default value on creation for the create_at field.
	couponallocated.DefaultCreateAt = couponallocatedDescCreateAt.Default.(func() uint32)
	// couponallocatedDescUpdateAt is the schema descriptor for update_at field.
	couponallocatedDescUpdateAt := couponallocatedFields[6].Descriptor()
	// couponallocated.DefaultUpdateAt holds the default value on creation for the update_at field.
	couponallocated.DefaultUpdateAt = couponallocatedDescUpdateAt.Default.(func() uint32)
	// couponallocated.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	couponallocated.UpdateDefaultUpdateAt = couponallocatedDescUpdateAt.UpdateDefault.(func() uint32)
	// couponallocatedDescDeleteAt is the schema descriptor for delete_at field.
	couponallocatedDescDeleteAt := couponallocatedFields[7].Descriptor()
	// couponallocated.DefaultDeleteAt holds the default value on creation for the delete_at field.
	couponallocated.DefaultDeleteAt = couponallocatedDescDeleteAt.Default.(func() uint32)
	// couponallocatedDescID is the schema descriptor for id field.
	couponallocatedDescID := couponallocatedFields[0].Descriptor()
	// couponallocated.DefaultID holds the default value on creation for the id field.
	couponallocated.DefaultID = couponallocatedDescID.Default.(func() uuid.UUID)
	couponpoolFields := schema.CouponPool{}.Fields()
	_ = couponpoolFields
	// couponpoolDescUsed is the schema descriptor for used field.
	couponpoolDescUsed := couponpoolFields[3].Descriptor()
	// couponpool.DefaultUsed holds the default value on creation for the used field.
	couponpool.DefaultUsed = couponpoolDescUsed.Default.(int)
	// couponpoolDescMessage is the schema descriptor for message field.
	couponpoolDescMessage := couponpoolFields[8].Descriptor()
	// couponpool.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	couponpool.MessageValidator = couponpoolDescMessage.Validators[0].(func(string) error)
	// couponpoolDescName is the schema descriptor for name field.
	couponpoolDescName := couponpoolFields[9].Descriptor()
	// couponpool.NameValidator is a validator for the "name" field. It is called by the builders before save.
	couponpool.NameValidator = couponpoolDescName.Validators[0].(func(string) error)
	// couponpoolDescCreateAt is the schema descriptor for create_at field.
	couponpoolDescCreateAt := couponpoolFields[10].Descriptor()
	// couponpool.DefaultCreateAt holds the default value on creation for the create_at field.
	couponpool.DefaultCreateAt = couponpoolDescCreateAt.Default.(func() uint32)
	// couponpoolDescUpdateAt is the schema descriptor for update_at field.
	couponpoolDescUpdateAt := couponpoolFields[11].Descriptor()
	// couponpool.DefaultUpdateAt holds the default value on creation for the update_at field.
	couponpool.DefaultUpdateAt = couponpoolDescUpdateAt.Default.(func() uint32)
	// couponpool.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	couponpool.UpdateDefaultUpdateAt = couponpoolDescUpdateAt.UpdateDefault.(func() uint32)
	// couponpoolDescDeleteAt is the schema descriptor for delete_at field.
	couponpoolDescDeleteAt := couponpoolFields[12].Descriptor()
	// couponpool.DefaultDeleteAt holds the default value on creation for the delete_at field.
	couponpool.DefaultDeleteAt = couponpoolDescDeleteAt.Default.(func() uint32)
	// couponpoolDescID is the schema descriptor for id field.
	couponpoolDescID := couponpoolFields[0].Descriptor()
	// couponpool.DefaultID holds the default value on creation for the id field.
	couponpool.DefaultID = couponpoolDescID.Default.(func() uuid.UUID)
	newuserrewardsettingFields := schema.NewUserRewardSetting{}.Fields()
	_ = newuserrewardsettingFields
	// newuserrewardsettingDescCreateAt is the schema descriptor for create_at field.
	newuserrewardsettingDescCreateAt := newuserrewardsettingFields[4].Descriptor()
	// newuserrewardsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	newuserrewardsetting.DefaultCreateAt = newuserrewardsettingDescCreateAt.Default.(func() uint32)
	// newuserrewardsettingDescUpdateAt is the schema descriptor for update_at field.
	newuserrewardsettingDescUpdateAt := newuserrewardsettingFields[5].Descriptor()
	// newuserrewardsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	newuserrewardsetting.DefaultUpdateAt = newuserrewardsettingDescUpdateAt.Default.(func() uint32)
	// newuserrewardsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	newuserrewardsetting.UpdateDefaultUpdateAt = newuserrewardsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// newuserrewardsettingDescDeleteAt is the schema descriptor for delete_at field.
	newuserrewardsettingDescDeleteAt := newuserrewardsettingFields[6].Descriptor()
	// newuserrewardsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	newuserrewardsetting.DefaultDeleteAt = newuserrewardsettingDescDeleteAt.Default.(func() uint32)
	// newuserrewardsettingDescID is the schema descriptor for id field.
	newuserrewardsettingDescID := newuserrewardsettingFields[0].Descriptor()
	// newuserrewardsetting.DefaultID holds the default value on creation for the id field.
	newuserrewardsetting.DefaultID = newuserrewardsettingDescID.Default.(func() uuid.UUID)
	purchaseinvitationFields := schema.PurchaseInvitation{}.Fields()
	_ = purchaseinvitationFields
	// purchaseinvitationDescCreateAt is the schema descriptor for create_at field.
	purchaseinvitationDescCreateAt := purchaseinvitationFields[4].Descriptor()
	// purchaseinvitation.DefaultCreateAt holds the default value on creation for the create_at field.
	purchaseinvitation.DefaultCreateAt = purchaseinvitationDescCreateAt.Default.(func() uint32)
	// purchaseinvitationDescUpdateAt is the schema descriptor for update_at field.
	purchaseinvitationDescUpdateAt := purchaseinvitationFields[5].Descriptor()
	// purchaseinvitation.DefaultUpdateAt holds the default value on creation for the update_at field.
	purchaseinvitation.DefaultUpdateAt = purchaseinvitationDescUpdateAt.Default.(func() uint32)
	// purchaseinvitation.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	purchaseinvitation.UpdateDefaultUpdateAt = purchaseinvitationDescUpdateAt.UpdateDefault.(func() uint32)
	// purchaseinvitationDescDeleteAt is the schema descriptor for delete_at field.
	purchaseinvitationDescDeleteAt := purchaseinvitationFields[6].Descriptor()
	// purchaseinvitation.DefaultDeleteAt holds the default value on creation for the delete_at field.
	purchaseinvitation.DefaultDeleteAt = purchaseinvitationDescDeleteAt.Default.(func() uint32)
	// purchaseinvitationDescID is the schema descriptor for id field.
	purchaseinvitationDescID := purchaseinvitationFields[0].Descriptor()
	// purchaseinvitation.DefaultID holds the default value on creation for the id field.
	purchaseinvitation.DefaultID = purchaseinvitationDescID.Default.(func() uuid.UUID)
	registrationinvitationFields := schema.RegistrationInvitation{}.Fields()
	_ = registrationinvitationFields
	// registrationinvitationDescCreateAt is the schema descriptor for create_at field.
	registrationinvitationDescCreateAt := registrationinvitationFields[1].Descriptor()
	// registrationinvitation.DefaultCreateAt holds the default value on creation for the create_at field.
	registrationinvitation.DefaultCreateAt = registrationinvitationDescCreateAt.Default.(func() uint32)
	// registrationinvitationDescUpdateAt is the schema descriptor for update_at field.
	registrationinvitationDescUpdateAt := registrationinvitationFields[2].Descriptor()
	// registrationinvitation.DefaultUpdateAt holds the default value on creation for the update_at field.
	registrationinvitation.DefaultUpdateAt = registrationinvitationDescUpdateAt.Default.(func() uint32)
	// registrationinvitation.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	registrationinvitation.UpdateDefaultUpdateAt = registrationinvitationDescUpdateAt.UpdateDefault.(func() uint32)
	// registrationinvitationDescDeleteAt is the schema descriptor for delete_at field.
	registrationinvitationDescDeleteAt := registrationinvitationFields[3].Descriptor()
	// registrationinvitation.DefaultDeleteAt holds the default value on creation for the delete_at field.
	registrationinvitation.DefaultDeleteAt = registrationinvitationDescDeleteAt.Default.(func() uint32)
	// registrationinvitationDescID is the schema descriptor for id field.
	registrationinvitationDescID := registrationinvitationFields[0].Descriptor()
	// registrationinvitation.DefaultID holds the default value on creation for the id field.
	registrationinvitation.DefaultID = registrationinvitationDescID.Default.(func() uuid.UUID)
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
