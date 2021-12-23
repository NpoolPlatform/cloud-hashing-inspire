package newuserrewardsetting

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-inspire/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateNewUserRewardSetting(info *npool.NewUserRewardSetting) error {
	if _, err := uuid.Parse(info.GetRegistrationCouponID()); err != nil {
		return xerrors.Errorf("invalid registration coupon id: %v", err)
	}
	if _, err := uuid.Parse(info.GetKycCouponID()); err != nil {
		return xerrors.Errorf("invalid kyc coupon id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToNewUserRewardSetting(row *ent.NewUserRewardSetting) *npool.NewUserRewardSetting {
	return &npool.NewUserRewardSetting{
		ID:                         row.ID.String(),
		AppID:                      row.AppID.String(),
		RegistrationCouponID:       row.RegistrationCouponID.String(),
		KycCouponID:                row.KycCouponID.String(),
		AutoGenerateInvitationCode: row.AutoGenerateInvitationCode,
	}
}

func Create(ctx context.Context, in *npool.CreateNewUserRewardSettingRequest) (*npool.CreateNewUserRewardSettingResponse, error) {
	if err := validateNewUserRewardSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		NewUserRewardSetting.
		Create().
		SetRegistrationCouponID(uuid.MustParse(in.GetInfo().GetRegistrationCouponID())).
		SetKycCouponID(uuid.MustParse(in.GetInfo().GetKycCouponID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create new user reward setting: %v", err)
	}

	return &npool.CreateNewUserRewardSettingResponse{
		Info: dbRowToNewUserRewardSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetNewUserRewardSettingRequest) (*npool.GetNewUserRewardSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		NewUserRewardSetting.
		Query().
		Where(
			newuserrewardsetting.And(
				newuserrewardsetting.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query new user reward setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty new user reward setting")
	}

	return &npool.GetNewUserRewardSettingResponse{
		Info: dbRowToNewUserRewardSetting(infos[0]),
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetNewUserRewardSettingByAppRequest) (*npool.GetNewUserRewardSettingByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		NewUserRewardSetting.
		Query().
		Where(
			newuserrewardsetting.And(
				newuserrewardsetting.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query new user reward setting: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty new user reward setting")
	}

	return &npool.GetNewUserRewardSettingByAppResponse{
		Info: dbRowToNewUserRewardSetting(infos[0]),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateNewUserRewardSettingRequest) (*npool.UpdateNewUserRewardSettingResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invlaid id: %v", err)
	}

	if err := validateNewUserRewardSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		NewUserRewardSetting.
		UpdateOneID(id).
		SetRegistrationCouponID(uuid.MustParse(in.GetInfo().GetRegistrationCouponID())).
		SetKycCouponID(uuid.MustParse(in.GetInfo().GetKycCouponID())).
		SetAutoGenerateInvitationCode(in.GetInfo().GetAutoGenerateInvitationCode()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update new user reward setting: %v", err)
	}

	return &npool.UpdateNewUserRewardSettingResponse{
		Info: dbRowToNewUserRewardSetting(info),
	}, nil
}
