package common

import (
	"time"

	"etnasys.com/clientService/protogen/go/user"
	"etnasys.com/clientService/src/core/model"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateUserRequestToUsermodel(req *user.CreateUserRequest) *model.UserModel {
	internalUserId, _ := uuid.NewV6()
	return &model.UserModel{
		Idp:            req.GetIdp().String(),
		IdpUserId:      req.GetIdpUserId(),
		Name:           req.GetName(),
		Email:          req.GetEmail(),
		State:          req.GetState(),
		Pin:            req.GetPin(),
		Country:        req.GetCountry(),
		UserType:       req.GetUserType().String(),
		OrgId:          "DEFAULT",
		InternalUserId: internalUserId.String(),
		Status:         user.UserStatus_active.String(),
		CreatedAt:      ConvertGRPCTimestampToDate(timestamppb.Now()),
		UpdatedAt:      ConvertGRPCTimestampToDate(timestamppb.New(time.Time{})),
		DeletedAt:      ConvertGRPCTimestampToDate(timestamppb.New(time.Time{})),
	}
}

func CreateUserToUserModel(req *user.User) *model.UserModel {
	return &model.UserModel{
		Idp:            req.GetIdp().String(),
		IdpUserId:      req.GetIdpUserId(),
		Name:           req.GetName(),
		Email:          req.GetEmail(),
		State:          req.GetState(),
		Pin:            req.GetPin(),
		Country:        req.GetCountry(),
		UserType:       req.GetUserType().String(),
		OrgId:          req.GetOrgId(),
		InternalUserId: req.GetInternalUserId(),
		Status:         req.GetStatus().String(),
		CreatedAt:      ConvertGRPCTimestampToDate(req.GetCreatedAt()),
		UpdatedAt:      ConvertGRPCTimestampToDate(timestamppb.Now()),
		DeletedAt:      ConvertGRPCTimestampToDate(timestamppb.New(time.Time{})),
	}
}

func UserModelToUser(userModel *model.UserModel) *user.User {

	return &user.User{
		Id:             userModel.Id,
		Idp:            user.UserIdp(user.UserIdp(user.UserIdp_value[userModel.Idp]).Number()),
		IdpUserId:      userModel.IdpUserId,
		Name:           userModel.Name,
		Email:          userModel.Email,
		State:          userModel.State,
		Pin:            userModel.Pin,
		Country:        userModel.Country,
		UserType:       user.UserType(user.UserType(user.UserType_value[userModel.UserType]).Number()),
		OrgId:          userModel.OrgId,
		InternalUserId: userModel.InternalUserId,
		Status:         user.UserStatus(user.UserStatus_value[userModel.Status]),
		CreatedAt:      ConvertDateToGRPCTimestamp(userModel.CreatedAt),
		UpdatedAt:      ConvertDateToGRPCTimestamp(userModel.UpdatedAt),
		DeletedAt:      ConvertDateToGRPCTimestamp(userModel.DeletedAt),
	}
}
