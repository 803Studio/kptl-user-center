package handler

import (
	"context"
	"errors"
	"github.com/803Studio/kptl-grpc-go/pkg/usercenter"
	"github.com/803Studio/kptl-user-center/internal/auth"
	"github.com/803Studio/kptl-user-center/internal/data"
	"github.com/803Studio/kptl-user-center/internal/model"
)

var (
	ErrEmptyIncoming      = errors.New("empty incoming message")
	ErrUnknownLoginMethod = errors.New("unknown login method")
)

func (KptlUserServer) Login(_ context.Context, in *usercenter.LoginRequest) (*usercenter.LoginResponse, error) {
	header := &usercenter.ResponseHeader{
		Status:  usercenter.ResponseStatus_InternalErr,
		Message: "",
	}

	res := &usercenter.LoginResponse{
		Header: header,
		Token:  "",
	}

	if in == nil {
		return nil, ErrEmptyIncoming
	}

	var err error = nil

	wxId := in.GetPayload()
	if wxId == "" {
		return res, nil
	}

	users, err := data.SelectUserAccountByWxId(1, wxId)
	if err != nil {
		return nil, err
	}

	var user *model.UserAccount
	if len(users) == 0 {
		user = new(model.UserAccount)
		user.WxId = wxId
		user.Role = model.RoleUser
		lastId, err := data.InsertIntoUserAccount(user)
		if err != nil {
			return nil, err
		}
		user.Id = uint32(lastId)
	} else {
		user = users[0]
	}

	token, err := auth.Sign(user, false)
	header.Status = usercenter.ResponseStatus_OK
	header.Message = "login ok"
	res.Token = token
	return res, nil
}
