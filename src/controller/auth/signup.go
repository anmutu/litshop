package auth

import (
	"litshop/src/lvm/context"
	"litshop/src/lvm/literr"
	"litshop/src/lvm/types"
	"litshop/src/model"
	"litshop/src/module/customer"
	"litshop/src/pkg/util"
	"litshop/src/request"
	"litshop/src/response"
)

// 常规注册
// 手机号+密码
// 邮箱+密码
func SignUp(ctx *context.Context, request *request.SignUpRequest) (*response.SignUpResponse, error) {
	user := &model.Customer{}
	login := &model.CustomerLogin{}

	var loginField string
	loginMethod := types.CustomerAuthTypeNull
	if util.IsValidatePhone(request.Phone) {
		loginMethod = types.CustomerAuthTypePhone
		user.Phone = request.Phone
		loginField = request.Phone
	} else {
		if util.IsValidateEmail(request.Email) {
			loginMethod = types.CustomerAuthTypeEmail
			loginField = request.Email
			user.Email = request.Email
		}
	}

	if !loginMethod.IsValid() {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	if !isValidatePassword(request.Password) {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	if request.Password != request.RePassword {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	// 检查是否已注册
	exists := false
	if loginMethod == types.CustomerAuthTypeEmail {
		exists = customer.IsEmailExists(loginField)
	}
	if loginMethod == types.CustomerAuthTypePhone {
		exists = customer.IsPhoneExists(loginField)
	}

	if exists {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	// 注册
	login.AuthType = loginMethod
	login.AuthId = loginField
	user.Password = request.Password

	err := customer.Create(user, login)
	if err != nil {
		return nil, literr.NewWithCode(literr.ErrCodeInvalidRequestParams)
	}

	return nil, nil
}
