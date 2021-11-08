// Code generated by entc, DO NOT EDIT.

package newuserrewardsetting

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// RegistrationCouponID applies equality check predicate on the "registration_coupon_id" field. It's identical to RegistrationCouponIDEQ.
func RegistrationCouponID(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegistrationCouponID), v))
	})
}

// KycCouponID applies equality check predicate on the "kyc_coupon_id" field. It's identical to KycCouponIDEQ.
func KycCouponID(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKycCouponID), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// RegistrationCouponIDEQ applies the EQ predicate on the "registration_coupon_id" field.
func RegistrationCouponIDEQ(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegistrationCouponID), v))
	})
}

// RegistrationCouponIDNEQ applies the NEQ predicate on the "registration_coupon_id" field.
func RegistrationCouponIDNEQ(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegistrationCouponID), v))
	})
}

// RegistrationCouponIDIn applies the In predicate on the "registration_coupon_id" field.
func RegistrationCouponIDIn(vs ...uuid.UUID) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegistrationCouponID), v...))
	})
}

// RegistrationCouponIDNotIn applies the NotIn predicate on the "registration_coupon_id" field.
func RegistrationCouponIDNotIn(vs ...uuid.UUID) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegistrationCouponID), v...))
	})
}

// RegistrationCouponIDGT applies the GT predicate on the "registration_coupon_id" field.
func RegistrationCouponIDGT(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegistrationCouponID), v))
	})
}

// RegistrationCouponIDGTE applies the GTE predicate on the "registration_coupon_id" field.
func RegistrationCouponIDGTE(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegistrationCouponID), v))
	})
}

// RegistrationCouponIDLT applies the LT predicate on the "registration_coupon_id" field.
func RegistrationCouponIDLT(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegistrationCouponID), v))
	})
}

// RegistrationCouponIDLTE applies the LTE predicate on the "registration_coupon_id" field.
func RegistrationCouponIDLTE(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegistrationCouponID), v))
	})
}

// KycCouponIDEQ applies the EQ predicate on the "kyc_coupon_id" field.
func KycCouponIDEQ(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKycCouponID), v))
	})
}

// KycCouponIDNEQ applies the NEQ predicate on the "kyc_coupon_id" field.
func KycCouponIDNEQ(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKycCouponID), v))
	})
}

// KycCouponIDIn applies the In predicate on the "kyc_coupon_id" field.
func KycCouponIDIn(vs ...uuid.UUID) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKycCouponID), v...))
	})
}

// KycCouponIDNotIn applies the NotIn predicate on the "kyc_coupon_id" field.
func KycCouponIDNotIn(vs ...uuid.UUID) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKycCouponID), v...))
	})
}

// KycCouponIDGT applies the GT predicate on the "kyc_coupon_id" field.
func KycCouponIDGT(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKycCouponID), v))
	})
}

// KycCouponIDGTE applies the GTE predicate on the "kyc_coupon_id" field.
func KycCouponIDGTE(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKycCouponID), v))
	})
}

// KycCouponIDLT applies the LT predicate on the "kyc_coupon_id" field.
func KycCouponIDLT(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKycCouponID), v))
	})
}

// KycCouponIDLTE applies the LTE predicate on the "kyc_coupon_id" field.
func KycCouponIDLTE(v uuid.UUID) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKycCouponID), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.NewUserRewardSetting {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NewUserRewardSetting) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NewUserRewardSetting) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NewUserRewardSetting) predicate.NewUserRewardSetting {
	return predicate.NewUserRewardSetting(func(s *sql.Selector) {
		p(s.Not())
	})
}
